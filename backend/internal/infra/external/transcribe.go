package external

import (
	"cmTranscribe/internal/infra/config"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
)

// TranscribeService Amazon Transcribeの操作を行うサービスです。
type TranscribeService struct {
	client *transcribeservice.TranscribeService
}

// NewTranscribeService 新しいTranscribeServiceを作成します。
func NewTranscribeService(region string) *TranscribeService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	return &TranscribeService{
		client: transcribeservice.New(sess),
	}
}

// StartTranscriptionJob 新しい文字起こしジョブを開始します。
func (t *TranscribeService) StartTranscriptionJob(jobName, mediaURI, bucketName string) error {
	input := &transcribeservice.StartTranscriptionJobInput{
		TranscriptionJobName: aws.String(jobName),
		LanguageCode:         aws.String(config.AppConfig.LanguageCode), // 言語コードは必要に応じて変更
		MediaFormat:          aws.String(config.AppConfig.MediaFormat),  // ファイル形式を指定
		Media: &transcribeservice.Media{
			MediaFileUri: aws.String(mediaURI),
		},
		OutputBucketName: aws.String(bucketName),
	}

	result, err := t.client.StartTranscriptionJob(input)
	if err != nil {
		fmt.Printf("Error starting transcription job: %v\n", err)
		return fmt.Errorf("failed to start transcription job: %v", err)
	}

	fmt.Printf("Started transcription job: %s\n", *result.TranscriptionJob.TranscriptionJobName)
	return nil
}
