package service

import (
	"cmTranscribe/internal/domain/model"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"path/filepath"
)

// S3StorageService は AWS S3 とのやり取りを行う具体的な実装です
type S3StorageService struct {
	s3Client *s3.S3
}

func NewS3StorageService(region string) *S3StorageService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	return &S3StorageService{
		s3Client: s3.New(sess),
	}
}

// UploadToS3 はファイルを S3 にアップロードします
func (s *S3StorageService) UploadToS3(s3File model.S3File) (string, error) {
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

	_, err = s.s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3File.BucketName), // バケット名を引数で受け取る
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	return fmt.Sprintf("s3://%s/%s", s3File.BucketName, key), nil
}
