package service

import (
	"cmTranscribe/internal/app/dto"
	"cmTranscribe/internal/domain/model"
	"cmTranscribe/internal/domain/repository"
	"cmTranscribe/internal/domain/service"
	"cmTranscribe/internal/shared/validator"
	"context"
	"fmt"
	"time"
)

// TranscriptionJobService 文字起こしに関連するサービスを提供します。
type TranscriptionJobService struct {
	Repo                    repository.TranscriptionJobRepository
	TranscriptionJobService service.TranscriptionJobService
}

// NewTranscriptionJobService 新しい TranscriptionJobService を作成します。
func NewTranscriptionJobService(repo repository.TranscriptionJobRepository, jobService service.TranscriptionJobService) *TranscriptionJobService {
	return &TranscriptionJobService{
		Repo:                    repo,
		TranscriptionJobService: jobService,
	}
}

// StartTranscriptionJob 新しい文字起こしジョブを開始します。
func (s *TranscriptionJobService) StartTranscriptionJob(ctx context.Context, req *dto.TranscriptionDto) (*dto.TranscriptionJobStatusResponseDto, error) {
	// UUIDを使用してユニークなジョブIDを生成
	//jobName := uuid.New().String()
	// ドメインモデルを作成
	job := model.NewTranscriptionJobDB(req.JobName, req.MediaURI, req.LanguageCode)

	// ジョブをリポジトリに保存
	err := s.Repo.Save(job)
	if err != nil {
		return nil, err
	}

	transcriptionJob := model.NewTranscriptionJob(req.JobName, req.MediaURI, req.LanguageCode, req.CustomVocabularyName)
	// バリデーションの実行
	if err := validator.Validate(transcriptionJob); err != nil {
		// エラーハンドリングのみ行う
		return nil, fmt.Errorf("error processing transcriptionJob: %v", err)
	}

	result, err := s.TranscriptionJobService.StartTranscriptionJob(ctx, transcriptionJob)
	if err != nil {
		return nil, fmt.Errorf("failed to start transcription job: %v", err)
	}
	response := &dto.TranscriptionJobStatusResponseDto{
		JobName:                result.JobName,
		TranscriptionJobStatus: result.TranscriptionJobStatus,
	}

	return response, nil
}

const timeFormat = "2006/01/02 15:04:05" // 'yyyy/MM/dd HH:mm:ss'形式

// GetTranscriptionJobList AWS Transcribeからジョブリストを取得します。
func (s *TranscriptionJobService) GetTranscriptionJobList(ctx context.Context) (*dto.TranscriptionJobsResponseDto, error) {
	// 日本時間のLocationを取得
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, fmt.Errorf("failed to load JST location: %v", err)
	}

	// AWS Transcribeからジョブリストを取得
	jobs, err := s.TranscriptionJobService.GetTranscriptionJobList(ctx) // ドメイン層のメソッドを呼び出す
	if err != nil {
		return nil, fmt.Errorf("failed to get transcription job list: %v", err)
	}

	// DTOに変換する
	var dtoJobs []dto.TranscriptionJobResponseDto
	for _, job := range jobs.Jobs { // jobs.Jobs はドメインモデル内のジョブリスト
		// CreationTimeとCompletionTimeをフォーマットしてDTOにセット
		dtoJob := dto.TranscriptionJobResponseDto{
			JobName:                job.JobName,
			CreationTime:           job.CreationTime.In(jst).Format(timeFormat),
			CompletionTime:         formatCompletionTime(job.CompletionTime, jst),
			LanguageCode:           job.LanguageCode,
			TranscriptionJobStatus: job.TranscriptionJobStatus,
			OutputLocationType:     job.OutputLocationType,
		}

		// DTOのValidateメソッドを呼び出す
		if err := dtoJob.Validate(); err != nil {
			return nil, fmt.Errorf("validation error: %v", err)
		}
		dtoJobs = append(dtoJobs, dtoJob)
	}

	// 最終的にDTOリストを返す
	response := dto.TranscriptionJobsResponseDto{
		Jobs: dtoJobs,
	}
	return &response, nil
}

// CompletionTimeのフォーマット（nilチェック付き、JSTに変換）
func formatCompletionTime(completionTime *time.Time, loc *time.Location) string {
	if completionTime != nil {
		return completionTime.In(loc).Format(timeFormat) // 日本時間に変換してフォーマット
	}
	return "" // nilの場合は空文字を返す
}
