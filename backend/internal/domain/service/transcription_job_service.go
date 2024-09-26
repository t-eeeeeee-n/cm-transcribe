package service

import (
	"cmTranscribe/internal/domain/model"
	"context"
)

type TranscriptionJobService interface {
	StartTranscriptionJob(ctx context.Context, input *model.TranscriptionJob) (*model.TranscriptionJobStatusResponse, error)
	GetTranscriptionJobList(ctx context.Context) (*model.TranscriptionJobSummariesResponse, error)
	GetTranscriptionJob(ctx context.Context, jobName string) (*model.TranscriptionJobResponse, error)
}

// NewTranscriptionJobService ファクトリ関数
func NewTranscriptionJobService(impl TranscriptionJobService) TranscriptionJobService {
	return impl
}
