package model

import (
	"fmt"
	"time"
)

// TranscriptionJob Amazon Transcribeに渡すデータ構造
type TranscriptionJob struct {
	JobName              string
	MediaFileURI         string
	LanguageCode         string
	CustomVocabularyName string
}

func NewTranscriptionJob(jobName, mediaFileUri, languageCode, customVocabularyName string) *TranscriptionJob {
	return &TranscriptionJob{
		JobName:              jobName,
		MediaFileURI:         mediaFileUri,
		LanguageCode:         languageCode,
		CustomVocabularyName: customVocabularyName,
	}
}

func (s *TranscriptionJob) Validate() error {
	if s.JobName == "" || s.MediaFileURI == "" || s.LanguageCode == "" {
		return fmt.Errorf("JobName, MediaFileURI, LanguageCode are required")
	}
	return nil
}

// TranscriptionJobDB 文字起こしジョブを表します。
type TranscriptionJobDB struct {
	ID           string
	MediaFileURI string
	Language     string
	Status       string
	CreatedAt    time.Time
}

// NewTranscriptionJobDB 新しいTranscriptionJobを作成します。
func NewTranscriptionJobDB(id, mediaFileUri, language string) *TranscriptionJobDB {
	return &TranscriptionJobDB{
		ID:           id,
		MediaFileURI: mediaFileUri,
		Language:     language,
		Status:       "Pending",
		CreatedAt:    time.Now(),
	}
}
