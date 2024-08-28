package service

import (
	"cmTranscribe/internal/domain/model"
	"cmTranscribe/internal/domain/service"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
)

// TranscribeService Amazon Transcribeの操作を行うサービスです。
type TranscribeService struct {
	client *transcribeservice.TranscribeService
}

// NewTranscribeService ファクトリ関数
func NewTranscribeService(region string) service.TranscriptionJobService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	return &TranscribeService{
		client: transcribeservice.New(sess),
	}
}

// StartTranscriptionJob インターフェースの実装
func (t *TranscribeService) StartTranscriptionJob(input *model.TranscriptionJob) (*model.TranscriptionJob, error) {
	transcriptionInput := &transcribeservice.StartTranscriptionJobInput{
		TranscriptionJobName: aws.String(input.JobName),
		LanguageCode:         aws.String(input.LanguageCode),
		Media: &transcribeservice.Media{
			MediaFileUri: aws.String(input.MediaFileURI),
		},
	}

	if input.CustomVocabularyName != "" {
		transcriptionInput.Settings = &transcribeservice.Settings{
			VocabularyName: aws.String(input.CustomVocabularyName),
		}
	}

	result, err := t.client.StartTranscriptionJob(transcriptionInput)
	if err != nil {
		fmt.Printf("Error starting transcription job: %v\n", err)
		return nil, fmt.Errorf("failed to start transcription job: %v", err)
	}

	job := model.NewTranscriptionJob(*result.TranscriptionJob.TranscriptionJobName, input.MediaFileURI, input.LanguageCode, input.CustomVocabularyName)
	return job, nil
}
