package routes

import (
	"cmTranscribe/internal/interface/api"
	"cmTranscribe/internal/shared/middleware"
	"net/http"
)

type Router struct {
	TranscriptionHandler    *api.TranscriptionJobHandler
	CustomVocabularyHandler *api.CustomVocabularyHandler
	S3UploadHandler         *api.S3UploadHandler
}

func NewRouter(
	transcriptionHandler *api.TranscriptionJobHandler,
	customVocabularyHandler *api.CustomVocabularyHandler,
	s3UploadHandler *api.S3UploadHandler,
) *Router {
	return &Router{
		TranscriptionHandler:    transcriptionHandler,
		CustomVocabularyHandler: customVocabularyHandler,
		S3UploadHandler:         s3UploadHandler,
	}
}

func (r *Router) RegisterRoutes() {
	http.Handle("/api/transcriptions/start", middleware.HttpMethodMiddleware(http.HandlerFunc(r.TranscriptionHandler.HandleStartJob), http.MethodPost))
	http.Handle("/api/custom/vocabulary/create", middleware.HttpMethodMiddleware(http.HandlerFunc(r.CustomVocabularyHandler.HandleCreateVocabulary), http.MethodPost))
	http.Handle("/api/custom/vocabulary/update", middleware.HttpMethodMiddleware(http.HandlerFunc(r.CustomVocabularyHandler.HandleUpdateVocabulary), http.MethodPost))
	http.Handle("/api/custom/vocabulary/get", middleware.HttpMethodMiddleware(http.HandlerFunc(r.CustomVocabularyHandler.HandleGetVocabularyByName), http.MethodGet))
	http.Handle("/api/s3/upload", middleware.HttpMethodMiddleware(http.HandlerFunc(r.S3UploadHandler.HandleUploadToS3), http.MethodPost))
}
