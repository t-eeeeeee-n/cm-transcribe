package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

// CustomVocabulary カスタムボキャブラリを表すドメインモデル
type CustomVocabulary struct {
	VocabularyName string // ボキャブラリーの名前
	LanguageCode   string // 言語コード
	FileUri        string // ボキャブラリーの語彙リストがあるURI
}

// NewCustomVocabulary 新しいCustomVocabularyを作成するファクトリ関数
func NewCustomVocabulary(name, language, fileUri string) *CustomVocabulary {
	return &CustomVocabulary{
		VocabularyName: name,
		LanguageCode:   language,
		FileUri:        fileUri,
	}
}

func (s *CustomVocabulary) Validate() error {
	if s.VocabularyName == "" || s.LanguageCode == "" {
		return fmt.Errorf("VocabularyName and LanguageCode are required")
	}
	return nil
}

// CustomVocabularyDB カスタムボキャブラリ（DB用）を表すドメインモデル
type CustomVocabularyDB struct {
	ID        string // ユニークな識別子
	Name      string // ボキャブラリーの名前
	Language  string // 言語コード
	FileUri   string // ボキャブラリーの語彙リストがあるURI
	State     string // ボキャブラリーのステータス
	CreatedAt time.Time
}

// NewCustomVocabularyDB 新しいCustomVocabularyを作成するファクトリ関数
func NewCustomVocabularyDB(name, language string, fileUri string) *CustomVocabularyDB {
	return &CustomVocabularyDB{
		ID:        uuid.New().String(),
		Name:      name,
		Language:  language,
		FileUri:   fileUri,
		State:     "Pending",
		CreatedAt: time.Now(),
	}
}

// CustomVocabularyResponse カスタムボキャブラリの返却値を表すドメインモデル
type CustomVocabularyResponse struct {
	VocabularyName   string    // ボキャブラリーの名前
	LanguageCode     string    // 言語コード
	FileUri          string    // ボキャブラリーの語彙リストがあるURI
	VocabularyState  string    // ステータス
	VocabularyLastModifiedTime time.Time // 最終更新日時
}

// NewCustomVocabularyResponse 新しいNewCustomVocabularyResponseを作成するファクトリ関数
func NewCustomVocabularyResponse(name, language, fileUri, state string, lastModifiedTime time.Time) *CustomVocabularyResponse {
	return &CustomVocabularyResponse{
		VocabularyName:   name,
		LanguageCode:     language,
		FileUri:          fileUri,
		VocabularyState:  state,
		VocabularyLastModifiedTime: lastModifiedTime,
	}
}
