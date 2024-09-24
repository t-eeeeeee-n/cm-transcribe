package persistence

import (
	"cmTranscribe/internal/domain/model"
	"fmt"
)

// TranscriptionJobRepository 文字起こしジョブを管理するためのリポジトリです。
type TranscriptionJobRepository struct {
	jobs map[string]*model.TranscriptionJobDB
}

// NewTranscriptionJobRepository 新しいTranscriptionJobRepositoryを作成します。
func NewTranscriptionJobRepository() (*TranscriptionJobRepository, error) {
	return &TranscriptionJobRepository{
		jobs: make(map[string]*model.TranscriptionJobDB),
	}, nil
}

// Save 文字起こしジョブを保存します。
func (r *TranscriptionJobRepository) Save(job *model.TranscriptionJobDB) error {
	if job == nil {
		return fmt.Errorf("failed to save transcription job: job is nil")
	}
	r.jobs[job.JobName] = job
	return nil
}

// FindByID IDで文字起こしジョブを検索します。
func (r *TranscriptionJobRepository) FindByID(id string) (*model.TranscriptionJobDB, error) {
	job, exists := r.jobs[id]
	if !exists {
		return nil, fmt.Errorf("transcription job with JobName %s not found", id)
	}
	return job, nil
}
