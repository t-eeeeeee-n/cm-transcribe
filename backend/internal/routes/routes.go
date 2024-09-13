package routes

import (
	"cmTranscribe/internal/interface/api"
	"cmTranscribe/internal/shared/middleware"
	"net/http"
)

type Router struct {
	TranscriptionHandler    *api.TranscriptionJobHandler
	CustomVocabularyHandler *api.CustomVocabularyHandler
}

func NewRouter(
	transcriptionHandler *api.TranscriptionJobHandler,
	customVocabularyHandler *api.CustomVocabularyHandler,
) *Router {
	return &Router{
		TranscriptionHandler:    transcriptionHandler,
		CustomVocabularyHandler: customVocabularyHandler,
	}
}

func (r *Router) RegisterRoutes() {
	http.Handle("/api/transcriptions/start", middleware.HttpMethodMiddleware(http.HandlerFunc(r.TranscriptionHandler.HandleStartJob), http.MethodPost))
	http.Handle("/api/custom/vocabulary/create", middleware.HttpMethodMiddleware(http.HandlerFunc(r.CustomVocabularyHandler.HandleCreateVocabulary), http.MethodPost))
	http.Handle("/api/custom/vocabulary/update", middleware.HttpMethodMiddleware(http.HandlerFunc(r.CustomVocabularyHandler.HandleUpdateVocabulary), http.MethodPost))
	http.Handle("/api/custom/vocabulary/get", middleware.HttpMethodMiddleware(http.HandlerFunc(r.CustomVocabularyHandler.HandleGetVocabularyByName), http.MethodGet))
}
