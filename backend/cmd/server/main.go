package main

import (
	"cmTranscribe/internal/infra/config"
	"cmTranscribe/internal/infra/container"
	"cmTranscribe/internal/interface/api"
	"cmTranscribe/internal/routes"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// コンテキストの作成 (キャンセルが可能)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// システム終了のシグナルをキャッチ
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// DIコンテナの初期化
	appContainer, err := container.NewAppContainer(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize app container: %v", err)
	}

	// APIハンドラの設定（プレゼンテーション層）
	transcriptionHandler := api.NewTranscriptionJobHandler(appContainer.TranscriptionJobService)
	customVocabularyHandler := api.NewCustomVocabularyHandler(appContainer.CustomVocabularyService)
	s3UploadHandler := api.NewS3UploadHandler(appContainer.S3UploadService)

	// ルーターの作成とルーティングの登録
	router := routes.NewRouter(
		transcriptionHandler,
		customVocabularyHandler,
		s3UploadHandler,
	)
	router.RegisterRoutes()

	// HTTPサーバーの作成
	srv := &http.Server{
		Addr:    ":" + config.AppConfig.Port,
		Handler: nil, // デフォルトで `http.DefaultServeMux` を使用
	}

	// サーバーの起動を別の goroutine で行う
	go func() {
		log.Printf("Starting server on :%s...\n", config.AppConfig.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// シグナルを受け取ったらシャットダウン処理を開始
	<-sigChan
	log.Println("Shutting down server...")

	// Graceful shutdown: 5秒のタイムアウトを設定
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
