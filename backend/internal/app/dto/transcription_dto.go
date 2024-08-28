package dto

// TranscriptionDto トランスクリプションジョブ作成時に使用するリクエストデータ
type TranscriptionDto struct {
	MediaURI             string `json:"media_uri"`                        // メディアファイルのURI
	LanguageCode         string `json:"language_code"`                    // 言語コード
	CustomVocabularyName string `json:"custom_vocabulary_name,omitempty"` // カスタムボキャブラリ名 (オプション)
}
