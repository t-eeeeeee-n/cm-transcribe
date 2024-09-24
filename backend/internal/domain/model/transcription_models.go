package model

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
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
	JobName      string
	MediaFileURI string
	Language     string
	Status       string
	CreatedAt    time.Time
}

// NewTranscriptionJobDB 新しいTranscriptionJobを作成します。
func NewTranscriptionJobDB(jobName, mediaFileUri, language string) *TranscriptionJobDB {
	return &TranscriptionJobDB{
		JobName:      jobName,
		MediaFileURI: mediaFileUri,
		Language:     language,
		Status:       "Pending",
		CreatedAt:    time.Now(),
	}
}

// TranscriptionJobStatusResponse は、ジョブ名とステータスを表す構造体
type TranscriptionJobStatusResponse struct {
	JobName                string
	TranscriptionJobStatus string
}

// NewTranscriptionJobStatusResponse は AWS Transcribe のジョブステータスをドメインモデルに変換します
func NewTranscriptionJobStatusResponse(jobName, transcriptionJobStatus string) *TranscriptionJobStatusResponse {
	return &TranscriptionJobStatusResponse{
		JobName:                jobName,
		TranscriptionJobStatus: transcriptionJobStatus,
	}
}

// TranscriptionJobResponse カスタムボキャブラリの返却値を表すドメインモデル
type TranscriptionJobResponse struct {
	JobName                string
	CreationTime           time.Time
	CompletionTime         *time.Time // nil を許容するためポインタに変更
	LanguageCode           string
	TranscriptionJobStatus string
	OutputLocationType     string
}

// NewTranscriptionJobResponse は AWS Transcribe のジョブ情報をドメインモデルに変換します
func NewTranscriptionJobResponse(job types.TranscriptionJobSummary) *TranscriptionJobResponse {
	return &TranscriptionJobResponse{
		JobName:                aws.ToString(job.TranscriptionJobName),
		CreationTime:           aws.ToTime(job.CreationTime),
		CompletionTime:         job.CompletionTime, // 完了していない場合は nil
		LanguageCode:           string(job.LanguageCode),
		TranscriptionJobStatus: string(job.TranscriptionJobStatus),
		OutputLocationType:     string(job.OutputLocationType),
	}
}

// TranscriptionJobsResponse は複数のジョブの返却値を表すドメインモデル
type TranscriptionJobsResponse struct {
	Jobs []*TranscriptionJobResponse // TranscriptionJobResponse を配列で保持
}

// NewTranscriptionJobsResponse は AWS Transcribe のジョブリストをドメインモデルに変換します
func NewTranscriptionJobsResponse(jobs []types.TranscriptionJobSummary) *TranscriptionJobsResponse {
	jobResponses := make([]*TranscriptionJobResponse, len(jobs))
	for i, job := range jobs {
		jobResponses[i] = NewTranscriptionJobResponse(job)
	}

	return &TranscriptionJobsResponse{
		Jobs: jobResponses,
	}
}
