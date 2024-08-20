package main

import (
	"cmTranscribe/internal/app/service"
	"cmTranscribe/internal/infra/config"
	"cmTranscribe/internal/infra/external"
	"cmTranscribe/internal/infra/persistence"
	"cmTranscribe/internal/interface/api"
	"log"
	"net/http"
)

func main() {
	// 設定の読み込み
	config.LoadConfig()

	// リポジトリの初期化
	repo := persistence.NewTranscriptionJobRepository()

	// Amazon Transcribeサービスの初期化
	transcribeSvc := external.NewTranscribeService(config.AppConfig.AWSRegion)

	// サービスの初期化
	transcriptionService := service.NewTranscriptionService(repo, transcribeSvc)

	// APIハンドラの設定
	transcriptionHandler := api.NewTranscriptionHandler(transcriptionService)

	// ルーティングの設定
	http.HandleFunc("/api/transcribe", transcriptionHandler.Handle)

	// サーバーの起動
	log.Printf("Starting server on :%s...\n", config.AppConfig.Port)
	if err := http.ListenAndServe(":"+config.AppConfig.Port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
