package service

import (
	"cmTranscribe/internal/domain/model"
)

// CustomVocabularyService カスタムボキャブラリに関するビジネスロジックを定義するインターフェース
type CustomVocabularyService interface {
	CreateCustomVocabulary(vocabulary model.CustomVocabulary) error
	UpdateCustomVocabulary(vocabulary model.CustomVocabulary) error
}

// NewCustomVocabularyService ファクトリ関数
func NewCustomVocabularyService(impl CustomVocabularyService) CustomVocabularyService {
	return impl
}
