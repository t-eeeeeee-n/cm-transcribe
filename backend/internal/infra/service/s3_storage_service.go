package service

import (
	"cmTranscribe/internal/domain/model"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
	"path/filepath"
)

// S3StorageService は AWS S3 とのやり取りを行う具体的な実装です
type S3StorageService struct {
	s3Client *s3.Client
}

// NewS3StorageService は S3StorageService のインスタンスを生成します
func NewS3StorageService(ctx context.Context, region string) (*S3StorageService, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %v", err)
	}

	return &S3StorageService{
		s3Client: s3.NewFromConfig(cfg),
	}, nil
}

// UploadToS3 はファイルを S3 にアップロードします
func (s *S3StorageService) UploadToS3(ctx context.Context, s3File model.S3File) (string, error) {
	// ファイルを開く
	file, err := os.Open(s3File.FilePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file for S3 upload: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}()

	// フォルダを指定してS3のキーを作成
	key := filepath.Join(s3File.KeyPrefix, filepath.Base(s3File.FilePath))

	// S3アップロード用のUploaderを作成 (aws-sdk-go-v2ではmanagerパッケージを使う)
	uploader := manager.NewUploader(s.s3Client)

	// ファイルをS3にアップロード
	_, err = uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s3File.BucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	return fmt.Sprintf("s3://%s/%s", s3File.BucketName, key), nil
}
