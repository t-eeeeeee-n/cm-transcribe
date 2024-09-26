package container

import (
	applicationService "cmTranscribe/internal/app/service"
	domainService "cmTranscribe/internal/domain/service"
	"cmTranscribe/internal/infra/config"
	"cmTranscribe/internal/infra/persistence"
	infraService "cmTranscribe/internal/infra/service"
	"context"
	"fmt"
)

// AppContainer は依存関係を保持する構造体です
type AppContainer struct {
	TranscriptionJobService *applicationService.TranscriptionJobService
	CustomVocabularyService *applicationService.CustomVocabularyService
	S3UploadService         *applicationService.S3UploadService
}

// NewAppContainer はアプリケーション全体の依存関係を初期化します
func NewAppContainer(ctx context.Context) (*AppContainer, error) {
	// 設定の読み込み
	if err := config.LoadConfig(); err != nil {
		return nil, err
	}

	// リポジトリの初期化
	transcriptionRepo, err := persistence.NewTranscriptionJobRepository()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize transcription repository: %w", err)
	}

	// 外部サービスの初期化
	transcribeInfraService, err := infraService.NewTranscribeService(ctx, config.AppConfig.AWSRegion)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize TranscribeInfraService: %w", err)
	}
	fileInfraService := infraService.NewFileService()
	s3StorageInfraService, err := infraService.NewS3StorageService(ctx, config.AppConfig.AWSRegion)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize S3StorageService: %w", err)
	}
	customVocabularyInfraService, err := infraService.NewCustomVocabularyService(ctx, config.AppConfig.AWSRegion)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize CustomVocabularyService: %w", err)
	}

	// ドメインサービスの初期化
	transcriptionJobService := domainService.NewTranscriptionJobService(transcribeInfraService)
	customVocabularyService := domainService.NewCustomVocabularyService(customVocabularyInfraService)
	fileService := domainService.NewFileService(fileInfraService)
	s3StorageService := domainService.NewS3StorageService(s3StorageInfraService)

	// アプリケーションサービスの初期化
	transcriptionJobAppService := applicationService.NewTranscriptionJobService(transcriptionRepo, transcriptionJobService, s3StorageService)
	customVocabularyAppService := applicationService.NewCustomVocabularyService(customVocabularyService, fileService, s3StorageService)
	s3UploadAppService := applicationService.NewS3UploadService(s3StorageService)

	return &AppContainer{
		TranscriptionJobService: transcriptionJobAppService,
		CustomVocabularyService: customVocabularyAppService,
		S3UploadService:         s3UploadAppService,
	}, nil
}
