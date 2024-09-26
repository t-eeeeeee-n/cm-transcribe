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

// TranscriptionJobSummaryDto GetTranscriptionJobList用のResponseDTO
type TranscriptionJobSummaryDto struct {
	JobName                string `json:"jobName"`
	CreationTime           string `json:"creationTime"`
	CompletionTime         string `json:"completionTime,omitempty"` // nilを許容し、nilの場合は省略される
	LanguageCode           string `json:"languageCode"`
	TranscriptionJobStatus string `json:"transcriptionJobStatus"`
	OutputLocationType     string `json:"outputLocationType"`
}

// Validate メソッドは、TranscriptionJobSummaryDto のバリデーションを行います
func (r *TranscriptionJobSummaryDto) Validate() error {
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

// TranscriptionJobsResponseDto GetTranscriptionJobList用のResponseDTO
type TranscriptionJobsResponseDto struct {
	Jobs []TranscriptionJobSummaryDto `json:"jobs"`
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

// TranscriptionJobResponseDto GetTranscriptionJob用のResponseDTO
type TranscriptionJobResponseDto struct {
	JobName                string `json:"jobName"`
	CreationTime           string `json:"creationTime"`
	CompletionTime         string `json:"completionTime,omitempty"`
	LanguageCode           string `json:"languageCode"`
	TranscriptionJobStatus string `json:"transcriptionJobStatus"`
	TranscriptFileUri      string `json:"transcriptFileUri"` // 実際の出力ファイルのURI
}

// Validate メソッドは、TranscriptionJobDetailResponseDto のバリデーションを行います
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
	return nil
}

// TranscriptionContentResponseDto is a struct that holds the simplified transcription data
type TranscriptionContentResponseDto struct {
	Transcript string                 `json:"transcript"` // 文字起こしのテキスト
	Confidence []WordConfidenceDto    `json:"confidence"` // 各単語の信頼度
	RawData    map[string]interface{} `json:"rawData"`    // 元のJSONデータ全体
}

// WordConfidenceDto holds each word with its confidence score
type WordConfidenceDto struct {
	Word       string `json:"word"`
	Confidence string `json:"confidence"`
}
