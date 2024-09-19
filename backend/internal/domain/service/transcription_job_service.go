package service

import (
	"cmTranscribe/internal/domain/model"
	"context"
)

type TranscriptionJobService interface {
	StartTranscriptionJob(ctx context.Context, input *model.TranscriptionJob) (*model.TranscriptionJob, error)
}

// NewTranscriptionJobService ファクトリ関数
func NewTranscriptionJobService(impl TranscriptionJobService) TranscriptionJobService {
	return impl
}
