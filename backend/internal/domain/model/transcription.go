package model

import "time"

// TranscriptionJob 文字起こしジョブを表します。
type TranscriptionJob struct {
	ID        string
	MediaURI  string
	Language  string
	Status    string
	CreatedAt time.Time
}

// NewTranscriptionJob 新しいTranscriptionJobを作成します。
func NewTranscriptionJob(id, mediaURI, language string) *TranscriptionJob {
	return &TranscriptionJob{
		ID:        id,
		MediaURI:  mediaURI,
		Language:  language,
		Status:    "Pending",
		CreatedAt: time.Now(),
	}
}

// Start 文字起こしジョブを開始します。
func (tj *TranscriptionJob) Start() {
	tj.Status = "InProgress"
}

// Complete 文字起こしジョブを完了します。
func (tj *TranscriptionJob) Complete() {
	tj.Status = "Completed"
}
