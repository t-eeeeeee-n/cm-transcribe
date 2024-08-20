package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config アプリケーションの設定を保持します。
type Config struct {
	Port         string
	AWSRegion    string
	S3BucketName string
	LanguageCode string
	MediaFormat  string
}

// AppConfig アプリケーション全体で使用される設定を保持します。
var AppConfig *Config

// LoadConfig 設定を読み込みます。
func LoadConfig() {
	// .env ファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}

	AppConfig = &Config{
		Port:         getEnv("PORT", "8080"),
		AWSRegion:    getEnv("AWS_REGION", "ap-northeast-1"),
		S3BucketName: getEnv("S3_BUCKET_NAME", "bucket-name"),
		LanguageCode: getEnv("LANGUAGE_CODE", "ja-JP"),
		MediaFormat:  getEnv("MEDIA_FORMAT", "mp3"),
	}
	log.Printf("Config loaded: %+v\n", AppConfig)
}

// getEnvは環境変数を取得し、存在しない場合はデフォルト値を返します。
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
