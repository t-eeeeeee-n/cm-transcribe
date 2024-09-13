package service

import (
	"cmTranscribe/internal/domain/model"
	"cmTranscribe/internal/domain/service"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"log"
)

// CustomVocabularyService AWS Transcribeを使用したCustomVocabularyServiceの実装
type CustomVocabularyService struct {
	//client   *transcribeservice.TranscribeService
	client   TranscribeServiceClient
	s3Client *s3.S3
}

// NewCustomVocabularyService 新しいAWSCustomVocabularyServiceを作成します
func NewCustomVocabularyService(region string) service.CustomVocabularyService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	return &CustomVocabularyService{
		client:   transcribeservice.New(sess),
		s3Client: s3.New(sess),
	}
}

func (s *CustomVocabularyService) CreateCustomVocabulary(vocabulary model.CustomVocabulary) error {
	input := &transcribeservice.CreateVocabularyInput{
		LanguageCode:      aws.String(vocabulary.LanguageCode),
		VocabularyName:    aws.String(vocabulary.VocabularyName),
		VocabularyFileUri: aws.String(vocabulary.FileUri),
	}

	_, err := s.client.CreateVocabulary(input)
	if err != nil {
		// AWS SDKのエラーの詳細を取得してバックエンドでログを出力
		var awsErr awserr.Error
		if errors.As(err, &awsErr) {
			log.Printf("Failed to create custom vocabulary: %s (code: %s, original error: %v)", awsErr.Message(), awsErr.Code(), awsErr.OrigErr())
		} else {
			log.Printf("Failed to create custom vocabulary: %v", err)
		}
		return fmt.Errorf("failed to create custom vocabulary: %v", err)
	}

	return nil
}

func (s *CustomVocabularyService) UpdateCustomVocabulary(vocabulary model.CustomVocabulary) error {
	input := &transcribeservice.UpdateVocabularyInput{
		LanguageCode:      aws.String(vocabulary.LanguageCode),
		VocabularyName:    aws.String(vocabulary.VocabularyName),
		VocabularyFileUri: aws.String(vocabulary.FileUri),
	}

	_, err := s.client.UpdateVocabulary(input)
	if err != nil {
		return fmt.Errorf("failed to update custom vocabulary: %v", err)
	}

	return nil
}

// GetCustomVocabularyByName 名前でカスタムボキャブラリを取得します
func (s *CustomVocabularyService) GetCustomVocabularyByName(name string) (*model.CustomVocabularyResponse, error) {
	// AWS Transcribeからカスタムボキャブラリの内容を取得
	input := &transcribeservice.GetVocabularyInput{
		VocabularyName: aws.String(name),
	}

	result, err := s.client.GetVocabulary(input)
	if err != nil {
		// AWS SDKのエラーの詳細を取得してログ出力
		var awsErr awserr.Error
		if errors.As(err, &awsErr) {
			log.Printf("Failed to get custom vocabulary: %s (code: %s, original error: %v)", awsErr.Message(), awsErr.Code(), awsErr.OrigErr())
		} else {
			log.Printf("Failed to get custom vocabulary: %v", err)
		}
		return nil, fmt.Errorf("failed to get custom vocabulary: %v", err)
	}

	// 取得した結果をドメインモデルに変換
	vocabulary := &model.CustomVocabularyResponse{
		VocabularyName:   *result.VocabularyName,
		LanguageCode:     *result.LanguageCode,
		FileUri:          *result.DownloadUri,
		VocabularyState:  *result.VocabularyState,
		VocabularyLastModifiedTime: *result.VocabularyLastModifiedTime,
	}

	return vocabulary, nil
}
