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

// TranscriptionJobResponse は単一の文字起こしジョブの返却値を表すドメインモデルです
type TranscriptionJobResponse struct {
	JobName                string
	CreationTime           time.Time
	CompletionTime         *time.Time
	LanguageCode           string
	TranscriptionJobStatus string
	OutputLocation         string
}

// NewTranscriptionJobResponse は AWS Transcribe のジョブ詳細情報をドメインモデルに変換します
func NewTranscriptionJobResponse(job *types.TranscriptionJob) *TranscriptionJobResponse {
	if job == nil {
		return nil // ジョブがnilの場合はnilを返す
	}
	return &TranscriptionJobResponse{
		JobName:                aws.ToString(job.TranscriptionJobName),
		CreationTime:           aws.ToTime(job.CreationTime),
		CompletionTime:         job.CompletionTime,
		LanguageCode:           string(job.LanguageCode),
		TranscriptionJobStatus: string(job.TranscriptionJobStatus),
		OutputLocation:         aws.ToString(job.Transcript.TranscriptFileUri),
	}
}

// TranscriptionJobSummaryResponse カスタムボキャブラリの返却値を表すドメインモデル
type TranscriptionJobSummaryResponse struct {
	JobName                string
	CreationTime           time.Time
	CompletionTime         *time.Time // nil を許容するためポインタに変更
	LanguageCode           string
	TranscriptionJobStatus string
	OutputLocationType     string
}

// NewTranscriptionJobSummaryResponse は AWS Transcribe のジョブ情報をドメインモデルに変換します
func NewTranscriptionJobSummaryResponse(job types.TranscriptionJobSummary) *TranscriptionJobSummaryResponse {
	return &TranscriptionJobSummaryResponse{
		JobName:                aws.ToString(job.TranscriptionJobName),
		CreationTime:           aws.ToTime(job.CreationTime),
		CompletionTime:         job.CompletionTime, // 完了していない場合は nil
		LanguageCode:           string(job.LanguageCode),
		TranscriptionJobStatus: string(job.TranscriptionJobStatus),
		OutputLocationType:     string(job.OutputLocationType),
	}
}

// TranscriptionJobSummariesResponse は複数のジョブの返却値を表すドメインモデル
type TranscriptionJobSummariesResponse struct {
	Jobs []*TranscriptionJobSummaryResponse // TranscriptionJobSummaryResponse を配列で保持
}

// NewTranscriptionJobSummariesResponse は AWS Transcribe のジョブリストをドメインモデルに変換します
func NewTranscriptionJobSummariesResponse(jobs []types.TranscriptionJobSummary) *TranscriptionJobSummariesResponse {
	jobResponses := make([]*TranscriptionJobSummaryResponse, len(jobs))
	for i, job := range jobs {
		jobResponses[i] = NewTranscriptionJobSummaryResponse(job)
	}

	return &TranscriptionJobSummariesResponse{
		Jobs: jobResponses,
	}
}
