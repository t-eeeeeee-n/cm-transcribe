package service

import (
	"cmTranscribe/internal/app/dto"
	"cmTranscribe/internal/domain/model"
	"cmTranscribe/internal/domain/repository"
	"cmTranscribe/internal/domain/service"
	"cmTranscribe/internal/shared/validator"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// TranscriptionJobService 文字起こしに関連するサービスを提供します。
type TranscriptionJobService struct {
	Repo                    repository.TranscriptionJobRepository
	TranscriptionJobService service.TranscriptionJobService
	S3StorageService        service.S3StorageService
}

// NewTranscriptionJobService 新しい TranscriptionJobService を作成します。
func NewTranscriptionJobService(
	repo repository.TranscriptionJobRepository,
	jobService service.TranscriptionJobService,
	s3StorageService service.S3StorageService,
) *TranscriptionJobService {
	return &TranscriptionJobService{
		Repo:                    repo,
		TranscriptionJobService: jobService,
		S3StorageService:        s3StorageService,
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
	var dtoJobs []dto.TranscriptionJobSummaryDto
	for _, job := range jobs.Jobs { // jobs.Jobs はドメインモデル内のジョブリスト
		// CreationTimeとCompletionTimeをフォーマットしてDTOにセット
		dtoJob := dto.TranscriptionJobSummaryDto{
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

// GetTranscriptionJob AWS Transcribeから特定のジョブを取得します。
func (s *TranscriptionJobService) GetTranscriptionJob(ctx context.Context, jobName string) (*dto.TranscriptionJobResponseDto, error) {
	// 日本時間のLocationを取得
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, fmt.Errorf("failed to load JST location: %v", err)
	}

	// AWS Transcribeから特定のジョブを取得
	job, err := s.TranscriptionJobService.GetTranscriptionJob(ctx, jobName) // ドメイン層のメソッドを呼び出し
	if err != nil {
		return nil, fmt.Errorf("failed to get transcription job: %v", err)
	}

	// DTOに変換する
	dtoJob := dto.TranscriptionJobResponseDto{
		JobName:                job.JobName,
		CreationTime:           job.CreationTime.In(jst).Format(timeFormat),   // 日本時間に変換
		CompletionTime:         formatCompletionTime(job.CompletionTime, jst), // 完了時間も日本時間に変換
		LanguageCode:           job.LanguageCode,
		TranscriptionJobStatus: job.TranscriptionJobStatus,
		TranscriptFileUri:      job.OutputLocation, // 出力ファイルのURLをセット
	}

	// DTOのバリデーションを実行
	if err := dtoJob.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	// DTOを返す
	return &dtoJob, nil
}

// CompletionTimeのフォーマット（nilチェック付き、JSTに変換）
func formatCompletionTime(completionTime *time.Time, loc *time.Location) string {
	if completionTime != nil {
		return completionTime.In(loc).Format(timeFormat) // 日本時間に変換してフォーマット
	}
	return "" // nilの場合は空文字を返す
}

// GetTranscriptionContent refactors transcription content for frontend
func (s *TranscriptionJobService) GetTranscriptionContent(ctx context.Context, transcriptFileUri string) (*dto.TranscriptionContentResponseDto, error) {
	// S3ストレージサービスを使って署名付きURLを生成
	signedURL, err := s.S3StorageService.GeneratePresignedURL(ctx, transcriptFileUri)
	if err != nil {
		return nil, fmt.Errorf("failed to generate signed URL: %v", err)
	}

	// 署名付きURLを使って文字起こしの内容を取得
	content, err := s.S3StorageService.GetTranscriptionContent(ctx, signedURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get transcription content: %v", err)
	}

	// JSONをパースして、全体のデータを取得
	var transcribeResult map[string]interface{}
	err = json.Unmarshal([]byte(content), &transcribeResult)
	if err != nil {
		return nil, fmt.Errorf("failed to parse transcription content: %v", err)
	}

	// transcriptsのテキストを取得
	transcripts := transcribeResult["results"].(map[string]interface{})["transcripts"].([]interface{})
	transcriptText := transcripts[0].(map[string]interface{})["transcript"].(string)

	// 各単語の信頼度を取得
	items := transcribeResult["results"].(map[string]interface{})["items"].([]interface{})
	var confidenceList []dto.WordConfidenceDto

	for _, item := range items {
		itemMap := item.(map[string]interface{})
		if itemMap["type"] == "pronunciation" {
			alternatives := itemMap["alternatives"].([]interface{})
			alt := alternatives[0].(map[string]interface{})
			confidenceList = append(confidenceList, dto.WordConfidenceDto{
				Word:       alt["content"].(string),
				Confidence: alt["confidence"].(string),
			})
		}
	}

	// 最終レスポンスを返す (パースしたデータと元のデータ全体の両方を含む)
	return &dto.TranscriptionContentResponseDto{
		Transcript: transcriptText,
		Confidence: confidenceList,
		RawData:    transcribeResult, // 元のデータをそのまま保持
	}, nil
}
