package main

import (
	"cmTranscribe/internal/infra/config"
	"cmTranscribe/internal/infra/container"
	"cmTranscribe/internal/interface/api"
	"cmTranscribe/internal/routes"
	"log"
	"net/http"
)

func main() {

	// DIコンテナの初期化
	appContainer, err := container.NewAppContainer()
	if err != nil {
		log.Fatalf("Failed to initialize app container: %v", err)
	}

	// APIハンドラの設定（プレゼンテーション層）
	transcriptionHandler := api.NewTranscriptionJobHandler(appContainer.TranscriptionJobService)
	customVocabularyHandler := api.NewCustomVocabularyHandler(appContainer.CustomVocabularyService)

	// ルーターの作成とルーティングの登録
	router := routes.NewRouter(
		transcriptionHandler,
		customVocabularyHandler,
	)
	router.RegisterRoutes()

	// サーバーの起動
	log.Printf("Starting server on :%s...\n", config.AppConfig.Port)
	if err := http.ListenAndServe(":"+config.AppConfig.Port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
