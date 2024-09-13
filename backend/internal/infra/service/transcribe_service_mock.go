package service

import (
	"github.com/aws/aws-sdk-go/service/transcribeservice"
)

// TranscribeServiceClient インターフェースは、AWS Transcribeのクライアントを抽象化します。
type TranscribeServiceClient interface {
	GetVocabulary(input *transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error)
	CreateVocabulary(input *transcribeservice.CreateVocabularyInput) (*transcribeservice.CreateVocabularyOutput, error)
	UpdateVocabulary(input *transcribeservice.UpdateVocabularyInput) (*transcribeservice.UpdateVocabularyOutput, error)
}
