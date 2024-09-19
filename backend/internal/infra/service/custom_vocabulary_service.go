package service

import (
	"cmTranscribe/internal/domain/model"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go"
	"github.com/pkg/errors"
	"log"
)

type CustomVocabularyService struct {
	client   *transcribe.Client
	s3Client *s3.Client
}

// NewCustomVocabularyService 新しいAWSCustomVocabularyServiceを作成します
func NewCustomVocabularyService(ctx context.Context, region string) (*CustomVocabularyService, error) {
	// デフォルト設定とリージョンを使用してAWS設定を読み込み
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	return &CustomVocabularyService{
		client:   transcribe.NewFromConfig(cfg),
		s3Client: s3.NewFromConfig(cfg),
	}, nil
}

func (s *CustomVocabularyService) CreateCustomVocabulary(ctx context.Context, vocabulary model.CustomVocabulary) error {
	input := &transcribe.CreateVocabularyInput{
		LanguageCode:      types.LanguageCode(vocabulary.LanguageCode), // LanguageCode を直接設定
		VocabularyName:    aws.String(vocabulary.VocabularyName),
		VocabularyFileUri: aws.String(vocabulary.FileUri),
	}

	_, err := s.client.CreateVocabulary(ctx, input)
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) {
			log.Printf("Failed to create custom vocabulary: %s (code: %s)", apiErr.ErrorMessage(), apiErr.ErrorCode())
		} else {
			log.Printf("Failed to create custom vocabulary: %v", err)
		}
		return fmt.Errorf("failed to create custom vocabulary: %v", err)
	}

	return nil
}

// UpdateCustomVocabulary updates an existing custom vocabulary
func (s *CustomVocabularyService) UpdateCustomVocabulary(ctx context.Context, vocabulary model.CustomVocabulary) error {
	// マッピング関数で言語コードを変換
	languageCode, err := s.mapLanguageCode(vocabulary.LanguageCode)
	if err != nil {
		return fmt.Errorf("unsupported language code: %v", err)
	}

	// AWS SDK v2 用の UpdateVocabularyInput を作成
	input := &transcribe.UpdateVocabularyInput{
		LanguageCode:      languageCode,
		VocabularyName:    aws.String(vocabulary.VocabularyName),
		VocabularyFileUri: aws.String(vocabulary.FileUri),
	}

	// Transcribe サービスを呼び出してカスタムボキャブラリを更新
	_, err = s.client.UpdateVocabulary(ctx, input)
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) {
			log.Printf("Failed to update custom vocabulary: %s (code: %s)", apiErr.ErrorMessage(), apiErr.ErrorCode())
		} else {
			log.Printf("Failed to update custom vocabulary: %v", err)
		}
		return fmt.Errorf("failed to update custom vocabulary: %v", err)
	}

	return nil
}

// mapLanguageCode は、string 型の LanguageCode を types.LanguageCode に変換します
func (s *CustomVocabularyService) mapLanguageCode(languageCode string) (types.LanguageCode, error) {
	switch languageCode {
	case "ja-JP":
		return types.LanguageCodeJaJp, nil
	case "en-US":
		return types.LanguageCodeEnUs, nil
	case "fr-FR":
		return types.LanguageCodeFrFr, nil
	case "es-ES":
		return types.LanguageCodeEsEs, nil
	default:
		return "", fmt.Errorf("unsupported language code: %s", languageCode)
	}
}

// GetCustomVocabularyByName 名前でカスタムボキャブラリを取得します
func (s *CustomVocabularyService) GetCustomVocabularyByName(ctx context.Context, name string) (*model.CustomVocabularyResponse, error) {
	// AWS Transcribeからカスタムボキャブラリの内容を取得
	input := &transcribe.GetVocabularyInput{
		VocabularyName: aws.String(name),
	}

	// API リクエスト
	result, err := s.client.GetVocabulary(ctx, input)
	if err != nil {
		// AWS SDK v2のエラーハンドリング
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) {
			log.Printf("Failed to get custom vocabulary: %s (code: %s)", apiErr.ErrorMessage(), apiErr.ErrorCode())
		} else {
			log.Printf("Failed to get custom vocabulary: %v", err)
		}
		return nil, fmt.Errorf("failed to get custom vocabulary: %v", err)
	}

	// 取得した結果をドメインモデルに変換
	vocabulary := &model.CustomVocabularyResponse{
		VocabularyName:             aws.ToString(result.VocabularyName),
		LanguageCode:               string(result.LanguageCode),
		FileUri:                    aws.ToString(result.DownloadUri),
		VocabularyState:            string(result.VocabularyState),
		VocabularyLastModifiedTime: aws.ToTime(result.LastModifiedTime),
	}

	return vocabulary, nil
}
