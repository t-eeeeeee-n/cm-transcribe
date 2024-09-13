package dto

import "time"

// CreateVocabularyDto カスタムボキャブラリ作成時に使用するリクエストデータ
type CreateVocabularyDto struct {
	VocabularyName string       `json:"name"`          // ボキャブラリーの名前
	LanguageCode   string       `json:"language_code"` // 言語コード
	Vocabularies   []Vocabulary `json:"vocabularies"`  // ボキャブラリーの語彙リスト
}

// UpdateVocabularyDto カスタムボキャブラリ更新時に使用するリクエストデータ
type UpdateVocabularyDto struct {
	VocabularyName string       `json:"name"`          // ボキャブラリーの名前
	LanguageCode   string       `json:"language_code"` // 言語コード
	Vocabularies   []Vocabulary `json:"vocabularies"`  // 追加するボキャブラリーの語彙のリスト
}

// CustomVocabularyResponse クライアントに返すデータ構造
type CustomVocabularyResponse struct {
	VocabularyName   string       `json:"vocabularyName"`
	LanguageCode     string       `json:"languageCode"`
	Vocabularies     []Vocabulary `json:"vocabularies"`
	VocabularyState  string       `json:"vocabularyState"`
	VocabularyLastModifiedTime time.Time    `json:"lastModifiedTime"`
}

// Vocabulary ボキャブラリーの語彙を表す構造体
type Vocabulary struct {
	Phrase     string `json:"phrase"`
	SoundsLike string `json:"soundsLike"`
	IPA        string `json:"ipa"`
	DisplayAs  string `json:"displayAs"`
}
