package repository

import "cmTranscribe/internal/domain/model"

// TranscriptionJobRepository 文字起こしジョブのリポジトリインターフェースです。
type TranscriptionJobRepository interface {
	Save(job *model.TranscriptionJob) error
	FindByID(id string) (*model.TranscriptionJob, error)
}
