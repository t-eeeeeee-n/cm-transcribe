package service

import (
	"cmTranscribe/internal/app/dto"
	"cmTranscribe/internal/domain/model"
	"cmTranscribe/internal/domain/repository"
	"cmTranscribe/internal/domain/service"
	"cmTranscribe/internal/shared/validator"
	"fmt"
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
func (s *TranscriptionJobService) StartTranscriptionJob(req *dto.TranscriptionDto) (*model.TranscriptionJobDB, error) {
	// ドメインモデルを作成
	job := model.NewTranscriptionJobDB("job-id", req.MediaURI, req.LanguageCode)

	// ジョブをリポジトリに保存
	err := s.Repo.Save(job)
	if err != nil {
		return nil, err
	}

	transcriptionJob := model.NewTranscriptionJob(job.ID, req.MediaURI, req.LanguageCode, req.CustomVocabularyName)
	// バリデーションの実行
	if err := validator.Validate(transcriptionJob); err != nil {
		// エラーハンドリングのみ行う
		return nil, fmt.Errorf("error processing transcriptionJob: %v", err)
	}

	_, err = s.TranscriptionJobService.StartTranscriptionJob(transcriptionJob)
	if err != nil {
		return nil, fmt.Errorf("failed to start transcription job: %v", err)
	}

	return job, nil
}
