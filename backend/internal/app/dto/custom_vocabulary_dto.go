package dto

// CreateVocabularyDto カスタムボキャブラリ作成時に使用するリクエストデータ
type CreateVocabularyDto struct {
	VocabularyName string            `json:"name"`          // ボキャブラリーの名前
	LanguageCode   string            `json:"language_code"` // 言語コード
	Vocabularies   []VocabularyEntry `json:"vocabularies"`  // ボキャブラリーの語彙リスト
}

// UpdateVocabularyDto カスタムボキャブラリ更新時に使用するリクエストデータ
type UpdateVocabularyDto struct {
	VocabularyName string            `json:"name"`          // ボキャブラリーの名前
	LanguageCode   string            `json:"language_code"` // 言語コード
	Vocabularies   []VocabularyEntry `json:"vocabularies"`  // 追加するボキャブラリーの語彙のリスト
}

// VocabularyEntry ボキャブラリーの語彙を表す構造体
type VocabularyEntry struct {
	Phrase     string `json:"phrase"`
	SoundsLike string `json:"soundsLike"`
	IPA        string `json:"ipa"`
	DisplayAs  string `json:"displayAs"`
}
