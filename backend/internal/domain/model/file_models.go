package model

import (
	"cmTranscribe/internal/app/dto"
	"cmTranscribe/internal/shared/constant"
	"fmt"
	"time"
)

type CSVFile struct {
	Name    string
	Path    string
	Content [][]string
}

func NewCSVFile(name, path string, content [][]string) *CSVFile {
	return &CSVFile{
		Name:    name,
		Path:    path,
		Content: content,
	}
}

func (s *CSVFile) Validate() error {
	if s.Name == "" || s.Path == "" || len(s.Content) == 0 {
		return fmt.Errorf("FilePath, BucketName, Content are required")
	}
	return nil
}

func GenerateFilePath(name string) string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("/tmp/%s_%s.csv", name, timestamp)
}

// ConvertEntriesToContent は[]dto.VocabularyEntryを[][]stringに変換する
func ConvertEntriesToContent(entries []dto.VocabularyEntry) [][]string {
	content := [][]string{constant.VocabularyCsvHeader} // ヘッダー
	for _, entry := range entries {
		record := []string{entry.Phrase, entry.IPA, entry.SoundsLike, entry.DisplayAs}
		content = append(content, record)
	}
	return content
}
