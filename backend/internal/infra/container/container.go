package container

import (
	applicationService "cmTranscribe/internal/app/service"
	domainService "cmTranscribe/internal/domain/service"
	"cmTranscribe/internal/infra/config"
	"cmTranscribe/internal/infra/persistence"
	infraService "cmTranscribe/internal/infra/service"
	"fmt"
)

// AppContainer は依存関係を保持する構造体です
type AppContainer struct {
	TranscriptionJobService *applicationService.TranscriptionJobService
	CustomVocabularyService *applicationService.CustomVocabularyService
}

// NewAppContainer はアプリケーション全体の依存関係を初期化します
func NewAppContainer() (*AppContainer, error) {
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
	transcribeInfraService := infraService.NewTranscribeService(config.AppConfig.AWSRegion)
	fileInfraService := infraService.NewFileService()
	s3StorageInfraService := infraService.NewS3StorageService(config.AppConfig.AWSRegion)
	customVocabularyInfraService := infraService.NewCustomVocabularyService(config.AppConfig.AWSRegion)

	// ドメインサービスの初期化
	transcriptionJobService := domainService.NewTranscriptionJobService(transcribeInfraService)
	customVocabularyService := domainService.NewCustomVocabularyService(customVocabularyInfraService)
	fileService := domainService.NewFileService(fileInfraService)
	s3StorageService := domainService.NewS3StorageService(s3StorageInfraService)

	// アプリケーションサービスの初期化
	transcriptionJobAppService := applicationService.NewTranscriptionJobService(transcriptionRepo, transcriptionJobService)
	customVocabularyAppService := applicationService.NewCustomVocabularyService(customVocabularyService, fileService, s3StorageService)

	return &AppContainer{
		TranscriptionJobService: transcriptionJobAppService,
		CustomVocabularyService: customVocabularyAppService,
	}, nil
}
