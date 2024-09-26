package service

import (
	"cmTranscribe/internal/domain/model"
	appConfig "cmTranscribe/internal/infra/config"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go"
	"log"
)

// TranscribeService Amazon Transcribeの操作を行うサービスです。
type TranscribeService struct {
	client *transcribe.Client
}

// NewTranscribeService ファクトリ関数
func NewTranscribeService(ctx context.Context, region string) (*TranscribeService, error) {
	// AWS Configを取得
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %v", err)
	}

	// Transcribe クライアントを作成
	client := transcribe.NewFromConfig(cfg)

	return &TranscribeService{
		client: client,
	}, nil
}

// StartTranscriptionJob インターフェースの実装
func (t *TranscribeService) StartTranscriptionJob(ctx context.Context, input *model.TranscriptionJob) (*model.TranscriptionJobStatusResponse, error) {
	bucketName := appConfig.AppConfig.S3BucketName
	// TranscriptionJobの入力を作成
	transcriptionInput := &transcribe.StartTranscriptionJobInput{
		TranscriptionJobName: aws.String(input.JobName),
		LanguageCode:         types.LanguageCode(input.LanguageCode),
		Media: &types.Media{
			MediaFileUri: aws.String(input.MediaFileURI),
		},
		OutputBucketName: aws.String(bucketName),
	}

	// カスタムボキャブラリの指定がある場合
	if input.CustomVocabularyName != "" {
		transcriptionInput.Settings = &types.Settings{
			VocabularyName: aws.String(input.CustomVocabularyName),
		}
	}

	// Transcriptionジョブを開始
	result, err := t.client.StartTranscriptionJob(ctx, transcriptionInput)
	if err != nil {
		var awsErr smithy.APIError
		if errors.As(err, &awsErr) && awsErr.ErrorCode() == "ConflictException" {
			// フロントエンドに重複エラーを知らせる
			return nil, fmt.Errorf("conflict: job name already exists")
		}
		fmt.Printf("Error starting transcription job: %v\n", err)
		return nil, fmt.Errorf("failed to start transcription job: %v", err)
	}

	// TranscriptionJobStatusResponse を作成して返す
	return model.NewTranscriptionJobStatusResponse(
		aws.ToString(result.TranscriptionJob.TranscriptionJobName),
		string(result.TranscriptionJob.TranscriptionJobStatus),
	), nil
}

// GetTranscriptionJob retrieves a single transcription job by its job name from AWS Transcribe.
func (t *TranscribeService) GetTranscriptionJob(ctx context.Context, jobName string) (*model.TranscriptionJobResponse, error) {
	// Create the input request with the provided job name.
	input := &transcribe.GetTranscriptionJobInput{
		TranscriptionJobName: aws.String(jobName),
	}

	// Call AWS Transcribes GetTranscriptionJob API.
	output, err := t.client.GetTranscriptionJob(ctx, input)
	if err != nil {
		// Handle API errors with logging and appropriate error messages.
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) {
			log.Printf("Failed to get transcription job: %s (code: %s)", apiErr.ErrorMessage(), apiErr.ErrorCode())
		} else {
			log.Printf("Failed to get transcription job: %v", err)
		}
		return nil, fmt.Errorf("failed to get transcription job: %v", err)
	}

	// Convert the AWS response to a domain model using a factory method.
	return model.NewTranscriptionJobResponse(output.TranscriptionJob), nil
}

// GetTranscriptionJobList retrieves the list of transcription jobs from AWS Transcribe.
func (t *TranscribeService) GetTranscriptionJobList(ctx context.Context) (*model.TranscriptionJobSummariesResponse, error) {
	// Create the request input for listing transcription jobs.
	input := &transcribe.ListTranscriptionJobsInput{}

	// Call AWS Transcribes ListTranscriptionJobs API.
	output, err := t.client.ListTranscriptionJobs(ctx, input)
	if err != nil {
		// Handle errors by logging and returning them.
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) {
			log.Printf("Failed to list transcription jobs: %s (code: %s)", apiErr.ErrorMessage(), apiErr.ErrorCode())
		} else {
			log.Printf("Failed to list transcription jobs: %v", err)
		}
		return nil, fmt.Errorf("failed to list transcription jobs: %v", err)
	}

	// Use the factory method to convert the AWS response to the domain model.
	return model.NewTranscriptionJobSummariesResponse(output.TranscriptionJobSummaries), nil
}
