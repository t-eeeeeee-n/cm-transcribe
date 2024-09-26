package api

import (
	"cmTranscribe/internal/app/dto"
	"cmTranscribe/internal/app/service"
	"cmTranscribe/internal/shared/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

// TranscriptionJobHandler APIリクエストを処理します。
type TranscriptionJobHandler struct {
	Service *service.TranscriptionJobService
}

// NewTranscriptionJobHandler 新しいTranscriptionHandlerを作成します。
func NewTranscriptionJobHandler(service *service.TranscriptionJobService) *TranscriptionJobHandler {
	return &TranscriptionJobHandler{
		Service: service,
	}
}

func (h *TranscriptionJobHandler) HandleStartJob(w http.ResponseWriter, r *http.Request) {
	var req dto.TranscriptionDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to parse JSON")
		return
	}

	fmt.Printf("TranscriptionDto: %v\n", req)

	job, err := h.Service.StartTranscriptionJob(r.Context(), &req)
	if err != nil {
		log.Println("err.Error():", err.Error())
		log.Println("err:", err)
		if strings.Contains(err.Error(), "conflict: job name already exists") {
			utils.RespondWithError(w, http.StatusConflict, err.Error()) // 409を返す
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to start transcription")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, job)
}

// HandleGetJobList 文字起こしジョブリストのAPIリクエストを処理します。
func (h *TranscriptionJobHandler) HandleGetJobList(w http.ResponseWriter, r *http.Request) {
	// サービスを使ってジョブリストを取得
	jobList, err := h.Service.GetTranscriptionJobList(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get transcription job list")
		return
	}

	// JSON形式でレスポンスを返す
	utils.RespondWithJSON(w, http.StatusOK, jobList)
}

// HandleGetJob 特定の文字起こしジョブのAPIリクエストを処理します。
func (h *TranscriptionJobHandler) HandleGetJob(w http.ResponseWriter, r *http.Request) {
	// パスパラメータからジョブ名を取得
	vars := mux.Vars(r)
	jobName := vars["jobName"]

	if jobName == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "jobName is required")
		return
	}

	// サービスを使って特定のジョブを取得
	job, err := h.Service.GetTranscriptionJob(r.Context(), jobName)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get transcription job")
		return
	}

	// JSON形式でレスポンスを返す
	utils.RespondWithJSON(w, http.StatusOK, job)
}

// HandleGetTranscriptionContent S3から文字起こし内容を取得し、DTOに詰めて返すエンドポイントハンドラ
func (h *TranscriptionJobHandler) HandleGetTranscriptionContent(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータからjobNameを取得
	jobName := r.URL.Query().Get("jobName")
	if jobName == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "jobName is required")
		return
	}

	// サービスを使って文字起こし内容を取得
	contentDto, err := h.Service.GetTranscriptionContent(r.Context(), jobName)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get transcription content")
		return
	}

	// 文字起こし内容をJSON形式で返す
	utils.RespondWithJSON(w, http.StatusOK, contentDto)
}
