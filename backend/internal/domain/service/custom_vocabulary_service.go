package service

import (
	"cmTranscribe/internal/domain/model"
	"context"
)

// CustomVocabularyService カスタムボキャブラリに関するビジネスロジックを定義するインターフェース
type CustomVocabularyService interface {
	CreateCustomVocabulary(ctx context.Context, vocabulary model.CustomVocabulary) error
	UpdateCustomVocabulary(ctx context.Context, vocabulary model.CustomVocabulary) error
	GetCustomVocabularyByName(ctx context.Context, name string) (*model.CustomVocabularyResponse, error)
}

// NewCustomVocabularyService ファクトリ関数
func NewCustomVocabularyService(impl CustomVocabularyService) CustomVocabularyService {
	return impl
}
