package service

import (
	"cmTranscribe/internal/domain/model"
)

type TranscriptionJobService interface {
	StartTranscriptionJob(input *model.TranscriptionJob) (*model.TranscriptionJob, error)
}

// NewTranscriptionJobService ファクトリ関数
func NewTranscriptionJobService(impl TranscriptionJobService) TranscriptionJobService {
	return impl
}
