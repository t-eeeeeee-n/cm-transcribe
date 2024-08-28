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
func NewCustomVocabulary(name, language string, fileUri string) *CustomVocabulary {
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
	Status    string // ボキャブラリーのステータス
	CreatedAt time.Time
}

// NewCustomVocabularyDB 新しいCustomVocabularyを作成するファクトリ関数
func NewCustomVocabularyDB(name, language string, fileUri string) *CustomVocabularyDB {
	return &CustomVocabularyDB{
		ID:        uuid.New().String(),
		Name:      name,
		Language:  language,
		FileUri:   fileUri,
		Status:    "Pending",
		CreatedAt: time.Now(),
	}
}
