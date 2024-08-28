package service

import (
	"cmTranscribe/internal/domain/model"
	"cmTranscribe/internal/domain/service"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
)

// CustomVocabularyService AWS Transcribeを使用したCustomVocabularyServiceの実装
type CustomVocabularyService struct {
	client   *transcribeservice.TranscribeService
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
