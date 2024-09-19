package service

import (
	"cmTranscribe/internal/domain/model"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
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
func (t *TranscribeService) StartTranscriptionJob(ctx context.Context, input *model.TranscriptionJob) (*model.TranscriptionJob, error) {
	// TranscriptionJobの入力を作成
	transcriptionInput := &transcribe.StartTranscriptionJobInput{
		TranscriptionJobName: aws.String(input.JobName),
		LanguageCode:         types.LanguageCode(input.LanguageCode),
		Media: &types.Media{
			MediaFileUri: aws.String(input.MediaFileURI),
		},
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
		fmt.Printf("Error starting transcription job: %v\n", err)
		return nil, fmt.Errorf("failed to start transcription job: %v", err)
	}

	// 結果から新しいTranscriptionJobオブジェクトを作成
	job := model.NewTranscriptionJob(
		aws.ToString(result.TranscriptionJob.TranscriptionJobName),
		input.MediaFileURI,
		input.LanguageCode,
		input.CustomVocabularyName,
	)
	return job, nil
}
