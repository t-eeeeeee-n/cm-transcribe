package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config アプリケーションの設定を保持します。
type Config struct {
	Port         string
	AWSRegion    string
	S3BucketName string
	MediaFormat  string
}

// AppConfig アプリケーション全体で使用される設定を保持します。
var AppConfig *Config

// LoadConfig 設定を読み込みます。
func LoadConfig() error {
	// .env ファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		// .env ファイルが見つからない場合やロードに失敗した場合、アプリケーションを終了するかどうかを判断
		log.Printf("Error loading .env file: %v\n", err)
		return err
	}

	AppConfig = &Config{
		Port:         getEnv("PORT", "8080"),
		AWSRegion:    getEnv("AWS_REGION", "ap-northeast-1"),
		S3BucketName: getEnv("S3_BUCKET_NAME", ""),
		MediaFormat:  getEnv("MEDIA_FORMAT", "mp3"),
	}

	// 必須の設定項目が不足している場合、エラーを返す
	if AppConfig.S3BucketName == "" {
		return fmt.Errorf("S3_BUCKET_NAME is required but not set")
	}

	return nil
}

// getEnvは環境変数を取得し、存在しない場合はデフォルト値を返します。
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
