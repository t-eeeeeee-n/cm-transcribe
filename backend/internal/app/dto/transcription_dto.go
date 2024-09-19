package dto

// TranscriptionDto トランスクリプションジョブ作成時に使用するリクエストデータ
type TranscriptionDto struct {
	MediaURI             string `json:"mediaUri"`                       // メディアファイルのURI
	LanguageCode         string `json:"languageCode"`                   // 言語コード
	CustomVocabularyName string `json:"customVocabularyName,omitempty"` // カスタムボキャブラリ名 (オプション)
}
