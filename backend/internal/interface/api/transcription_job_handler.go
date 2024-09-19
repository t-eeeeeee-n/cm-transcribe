package api

import (
	"cmTranscribe/internal/app/dto"
	"cmTranscribe/internal/app/service"
	"cmTranscribe/internal/shared/utils"
	"encoding/json"
	"fmt"
	"net/http"
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

// HandleStartJob 文字起こしジョブのAPIリクエストを処理します。
func (h *TranscriptionJobHandler) HandleStartJob(w http.ResponseWriter, r *http.Request) {
	var req dto.TranscriptionDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to parse JSON")
		return
	}

	fmt.Printf("TranscriptionDto: %v\n", req)

	job, err := h.Service.StartTranscriptionJob(r.Context(), &req)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to start transcription")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, job)
}
