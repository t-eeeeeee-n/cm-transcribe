package repository

import "cmTranscribe/internal/domain/model"

// TranscriptionJobRepository 文字起こしジョブのリポジトリインターフェースです。
type TranscriptionJobRepository interface {
	Save(job *model.TranscriptionJobDB) error
	FindByID(id string) (*model.TranscriptionJobDB, error)
}
