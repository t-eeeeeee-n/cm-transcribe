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

// ConvertEntriesToContent は[]dto.Vocabulary[][]stringに変換する
func ConvertEntriesToContent(vocabularies []dto.Vocabulary) [][]string {
	content := [][]string{constant.VocabularyCsvHeader} // ヘッダー
	for _, vocabulary := range vocabularies {
		record := []string{vocabulary.Phrase, vocabulary.IPA, vocabulary.SoundsLike, vocabulary.DisplayAs}
		content = append(content, record)
	}
	return content
}
