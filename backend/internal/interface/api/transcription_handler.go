package api

import (
	"cmTranscribe/internal/app/service"
	"cmTranscribe/internal/infra/config"
	"encoding/json"
	"net/http"
)

// TranscriptionHandler APIリクエストを処理します。
type TranscriptionHandler struct {
	Service *service.TranscriptionService
}

// NewTranscriptionHandler 新しいTranscriptionHandlerを作成します。
func NewTranscriptionHandler(service *service.TranscriptionService) *TranscriptionHandler {
	return &TranscriptionHandler{
		Service: service,
	}
}

// Handle 文字起こしジョブのAPIリクエストを処理します。
func (h *TranscriptionHandler) Handle(w http.ResponseWriter, r *http.Request) {
	mediaURI := r.URL.Query().Get("media_uri")
	language := config.AppConfig.LanguageCode

	job, err := h.Service.StartTranscriptionJob(mediaURI, language)
	if err != nil {
		http.Error(w, "Failed to start transcription", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}
