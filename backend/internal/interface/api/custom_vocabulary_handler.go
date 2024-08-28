package api

import (
	"cmTranscribe/internal/app/dto"
	"cmTranscribe/internal/app/service"
	"cmTranscribe/internal/shared/utils"
	"encoding/json"
	"net/http"
)

// CustomVocabularyHandler ハンドラ構造体
type CustomVocabularyHandler struct {
	Service *service.CustomVocabularyService
}

func NewCustomVocabularyHandler(service *service.CustomVocabularyService) *CustomVocabularyHandler {
	return &CustomVocabularyHandler{
		Service: service,
	}
}

func (h *CustomVocabularyHandler) HandleCreateVocabulary(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateVocabularyDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to parse JSON")
		return
	}

	// DTOをサービスに渡して処理
	err := h.Service.CreateCustomVocabulary(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create custom vocabulary")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Custom vocabulary created successfully"}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// HandleUpdateVocabulary カスタムボキャブラリの更新リクエストを処理します。
func (h *CustomVocabularyHandler) HandleUpdateVocabulary(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateVocabularyDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to parse JSON")
		return
	}

	// DTOをサービスに渡して処理
	err := h.Service.UpdateCustomVocabulary(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update custom vocabulary")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Custom vocabulary updated successfully"}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
