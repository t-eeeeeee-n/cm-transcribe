package service

import (
	"cmTranscribe/internal/domain/model"
	"cmTranscribe/internal/domain/service"
	"encoding/csv"
	"fmt"
	"os"
)

// FileService ファイル操作に関連するサービスの実装
type FileService struct {
	fileCreator FileCreator
}

// FileCreator インターフェースでファイル操作を抽象化
type FileCreator interface {
	Create(name string) (*os.File, error)
	Remove(name string) error
}

// NewFileService デフォルトのOSFileCreatorを使用してFileServiceを初期化
func NewFileService() service.FileService {
	return &FileService{
		fileCreator: &OSFileCreator{},
	}
}

// OSFileCreator 実際のOSファイル操作の実装
type OSFileCreator struct{}

func (c *OSFileCreator) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func (c *OSFileCreator) Remove(name string) error {
	return os.Remove(name)
}

// CreateCSV CSVファイルを作成
func (f *FileService) CreateCSV(csvFile model.CSVFile) (string, *os.File, error) {
	file, err := f.fileCreator.Create(csvFile.Path)
	if err != nil {
		return "", nil, fmt.Errorf("failed to create CSV file: %v", err)
	}

	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	defer writer.Flush()

	// CSVの内容を書き込み
	for _, record := range csvFile.Content {
		if err := writer.Write(record); err != nil {
			return "", nil, fmt.Errorf("failed to write CSV record: %v", err)
		}
	}

	return csvFile.Path, file, nil
}

// Cleanup ファイルを削除
func (f *FileService) Cleanup(filePath string) error {
	return f.fileCreator.Remove(filePath)
}
