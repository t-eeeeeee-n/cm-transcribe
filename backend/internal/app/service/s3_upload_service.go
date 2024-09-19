package service

import (
	"cmTranscribe/internal/domain/model"
	"cmTranscribe/internal/domain/service"
	"cmTranscribe/internal/shared/validator" // validatorパッケージをインポート
	"context"
	"fmt"
)

type S3UploadService struct {
	s3StorageService service.S3StorageService
}

func NewS3UploadService(s3StorageService service.S3StorageService) *S3UploadService {
	return &S3UploadService{s3StorageService: s3StorageService}
}

// UploadToS3 は、指定されたファイルを S3 にアップロードします
func (s *S3UploadService) UploadToS3(ctc context.Context, filePath, bucketName, keyPrefix string) (string, error) {
	// S3Fileのインスタンスを作成
	s3File := model.S3File{
		FilePath:   filePath,
		BucketName: bucketName,
		KeyPrefix:  keyPrefix,
	}

	// バリデーションの共通ロジックを適用
	if err := validator.Validate(&s3File); err != nil {
		return "", fmt.Errorf("validation failed: %v", err)
	}

	// S3 にアップロード
	url, err := s.s3StorageService.UploadToS3(ctc, s3File)
	if err != nil {
		return "", fmt.Errorf("failed to upload to S3: %v", err)
	}
	return url, nil
}
