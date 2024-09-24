package dto

import (
	"errors"
)

// TranscriptionDto トランスクリプションジョブ作成時に使用するリクエストデータ
type TranscriptionDto struct {
	JobName              string `json:"jobName"`
	MediaURI             string `json:"mediaUri"`                       // メディアファイルのURI
	LanguageCode         string `json:"languageCode"`                   // 言語コード
	CustomVocabularyName string `json:"customVocabularyName,omitempty"` // カスタムボキャブラリ名 (オプション)
}

// TranscriptionJobStatusResponseDto 文字起こしジョブのレスポンスで使用されるDTO
type TranscriptionJobStatusResponseDto struct {
	JobName                string `json:"jobName"`
	TranscriptionJobStatus string `json:"jobStatus"`
}

// TranscriptionJobResponseDto は Transcription ジョブのレスポンスを表すDTOです
type TranscriptionJobResponseDto struct {
	JobName                string `json:"jobName"`
	CreationTime           string `json:"creationTime"`
	CompletionTime         string `json:"completionTime,omitempty"` // nilを許容し、nilの場合は省略される
	LanguageCode           string `json:"languageCode"`
	TranscriptionJobStatus string `json:"transcriptionJobStatus"`
	OutputLocationType     string `json:"outputLocationType"`
}

// Validate メソッドは、TranscriptionJobResponseDto のバリデーションを行います
func (r *TranscriptionJobResponseDto) Validate() error {
	if r.JobName == "" {
		return errors.New("JobName is required")
	}
	if r.LanguageCode == "" {
		return errors.New("LanguageCode is required")
	}
	if r.TranscriptionJobStatus == "" {
		return errors.New("TranscriptionJobStatus is required")
	}
	// 他のフィールドのバリデーションも必要なら追加する
	return nil
}

type TranscriptionJobsResponseDto struct {
	Jobs []TranscriptionJobResponseDto `json:"jobs"`
}

// Validate メソッドは、TranscriptionJobsResponseDto のバリデーションを行います
func (r *TranscriptionJobsResponseDto) Validate() error {
	if len(r.Jobs) == 0 {
		return errors.New("no jobs found")
	}
	for _, job := range r.Jobs {
		if err := job.Validate(); err != nil {
			return err
		}
	}
	return nil
}
