package service

import (
	"cmTranscribe/internal/domain/model"
	appConfig "cmTranscribe/internal/infra/config"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// S3StorageService は AWS S3 とのやり取りを行う具体的な実装です
type S3StorageService struct {
	s3Client      *s3.Client
	presignClient *s3.PresignClient
}

// NewS3StorageService は S3StorageService のインスタンスを生成します
func NewS3StorageService(ctx context.Context, region string) (*S3StorageService, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	presignClient := s3.NewPresignClient(s3Client)

	return &S3StorageService{
		s3Client:      s3Client,
		presignClient: presignClient,
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

// GeneratePresignedURL は、transcriptFileUri から署名付きURLを生成します
func (s *S3StorageService) GeneratePresignedURL(ctx context.Context, jobName string) (string, error) {
	// 環境変数からバケット名を取得
	bucketName := appConfig.AppConfig.S3BucketName
	key := fmt.Sprintf("%s.json", jobName)

	// presignClient を使って署名付きURLを生成
	req, err := s.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(15*time.Minute)) // 15分間有効

	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %v", err)
	}

	return req.URL, nil
}

// GetTranscriptionContent は、署名付きURLから文字起こしデータを取得する関数です
func (s *S3StorageService) GetTranscriptionContent(ctx context.Context, signedURL string) (string, error) {
	// HTTP GET リクエストで署名付きURLにアクセス
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, signedURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// HTTPクライアントでリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}()

	// レスポンスボディを読み込む
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(content), nil
}
