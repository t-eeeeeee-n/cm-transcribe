package service

import (
	"cmTranscribe/internal/domain/model"
	"os"
)

type FileService interface {
	CreateCSV(csvFile model.CSVFile) (string, *os.File, error)
	Cleanup(filePath string) error
}

// NewFileService ファクトリ関数
func NewFileService(impl FileService) FileService {
	return impl
}
