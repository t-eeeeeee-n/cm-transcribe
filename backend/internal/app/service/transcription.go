package service

import (
	"cmTranscribe/internal/domain/model"
	"cmTranscribe/internal/domain/repository"
	"cmTranscribe/internal/infra/config"
	"cmTranscribe/internal/infra/external"
	"fmt"
)

// TranscriptionService 文字起こしに関連するサービスを提供します。
type TranscriptionService struct {
	Repo          repository.TranscriptionJobRepository
	TranscribeSvc *external.TranscribeService
}

// NewTranscriptionService 新しいTranscriptionServiceを作成します。
func NewTranscriptionService(repo repository.TranscriptionJobRepository, transcribeSvc *external.TranscribeService) *TranscriptionService {
	return &TranscriptionService{
		Repo:          repo,
		TranscribeSvc: transcribeSvc,
	}
}

// StartTranscriptionJob 新しい文字起こしジョブを開始します。
func (s *TranscriptionService) StartTranscriptionJob(mediaURI, language string) (*model.TranscriptionJob, error) {
	job := model.NewTranscriptionJob("job-id", mediaURI, language)
	err := s.Repo.Save(job)
	if err != nil {
		return nil, err
	}

	// Amazon Transcribeで文字起こしジョブを開始
	err = s.TranscribeSvc.StartTranscriptionJob(job.ID, mediaURI, config.AppConfig.S3BucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to start transcription job: %v", err)
	}

	return job, nil
}
