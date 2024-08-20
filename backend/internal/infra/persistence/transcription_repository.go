package persistence

import (
	"cmTranscribe/internal/domain/model"
	"errors"
)

// TranscriptionJobRepository 文字起こしジョブを管理するためのリポジトリです。
type TranscriptionJobRepository struct {
	jobs map[string]*model.TranscriptionJob
}

// NewTranscriptionJobRepository 新しいTranscriptionJobRepositoryを作成します。
func NewTranscriptionJobRepository() *TranscriptionJobRepository {
	return &TranscriptionJobRepository{
		jobs: make(map[string]*model.TranscriptionJob),
	}
}

// Save 文字起こしジョブを保存します。
func (r *TranscriptionJobRepository) Save(job *model.TranscriptionJob) error {
	if job == nil {
		return errors.New("job is nil")
	}
	r.jobs[job.ID] = job
	return nil
}

// FindByID IDで文字起こしジョブを検索します。
func (r *TranscriptionJobRepository) FindByID(id string) (*model.TranscriptionJob, error) {
	job, exists := r.jobs[id]
	if !exists {
		return nil, errors.New("job not found")
	}
	return job, nil
}
