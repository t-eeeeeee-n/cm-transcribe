package service

import (
	"cmTranscribe/internal/domain/model"
	"context"
)

type S3StorageService interface {
	UploadToS3(ctx context.Context, s3File model.S3File) (string, error)
}

// NewS3StorageService ファクトリ関数
func NewS3StorageService(impl S3StorageService) S3StorageService {
	return impl
}
