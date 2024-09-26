package routes

import (
	"cmTranscribe/internal/interface/api"
	"cmTranscribe/internal/shared/middleware"
	"github.com/gorilla/mux"
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

func (r *Router) RegisterRoutes() *mux.Router {

	router := mux.NewRouter()

	router.Handle("/api/transcriptions", middleware.HttpMethodMiddleware(http.HandlerFunc(r.TranscriptionHandler.HandleGetJobList), http.MethodGet))
	router.Handle("/api/transcriptions/start", middleware.HttpMethodMiddleware(http.HandlerFunc(r.TranscriptionHandler.HandleStartJob), http.MethodPost))
	router.Handle("/api/transcriptions/content", middleware.HttpMethodMiddleware(http.HandlerFunc(r.TranscriptionHandler.HandleGetTranscriptionContent), http.MethodGet))
	router.Handle("/api/transcriptions/{jobName}", middleware.HttpMethodMiddleware(http.HandlerFunc(r.TranscriptionHandler.HandleGetJob), http.MethodGet))
	router.Methods(http.MethodPost).Path("/api/custom/vocabulary").Handler(middleware.HttpMethodMiddleware(http.HandlerFunc(r.CustomVocabularyHandler.HandleCreateVocabulary), http.MethodPost))
	router.Methods(http.MethodPut).Path("/api/custom/vocabulary").Handler(middleware.HttpMethodMiddleware(http.HandlerFunc(r.CustomVocabularyHandler.HandleUpdateVocabulary), http.MethodPut))
	router.Methods(http.MethodGet).Path("/api/custom/vocabulary").Handler(middleware.HttpMethodMiddleware(http.HandlerFunc(r.CustomVocabularyHandler.HandleGetVocabularyByName), http.MethodGet))
	//router.Handle("/api/custom/vocabulary", http.HandlerFunc(r.CustomVocabularyHandler.HandleVocabulary))
	router.Handle("/api/s3/upload", middleware.HttpMethodMiddleware(http.HandlerFunc(r.S3UploadHandler.HandleUploadToS3), http.MethodPost))
	return router
}
