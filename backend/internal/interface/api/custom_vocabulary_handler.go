package api

import (
	"cmTranscribe/internal/app/dto"
	"cmTranscribe/internal/app/service"
	"cmTranscribe/internal/shared/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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

func (h *CustomVocabularyHandler) HandleVocabulary(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.HandleCreateVocabulary(w, r)
	case http.MethodPut:
		h.HandleUpdateVocabulary(w, r)
	case http.MethodGet:
		h.HandleGetVocabularyByName(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CustomVocabularyHandler) HandleCreateVocabulary(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateVocabularyDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to parse JSON")
		return
	}

	// DTOをサービスに渡して処理
	err := h.Service.CreateCustomVocabulary(r.Context(), req)
	if err != nil {
		log.Println("err.Error():", err.Error())
		log.Println("err:", err)
		if strings.Contains(err.Error(), "conflict: custom vocabulary name already exists") {
			utils.RespondWithError(w, http.StatusConflict, err.Error())
			return
		}
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
	err := h.Service.UpdateCustomVocabulary(r.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), "conflict: custom vocabulary name already exists") {
			utils.RespondWithError(w, http.StatusConflict, err.Error())
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update custom vocabulary")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Custom vocabulary updated successfully"}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// HandleGetVocabularyByName カスタムボキャブラリの内容を取得します。
func (h *CustomVocabularyHandler) HandleGetVocabularyByName(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータからカスタムボキャブラリーの名前を取得
	vocabularyName := r.URL.Query().Get("name")
	if vocabularyName == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Missing vocabulary name")
		return
	}

	// サービスを使ってカスタムボキャブラリーの内容を取得
	vocabulary, err := h.Service.GetCustomVocabularyByName(r.Context(), vocabularyName)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving vocabulary: %v", err))
		return
	}

	// JSONレスポンスを返す
	response, err := json.Marshal(vocabulary)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling response: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error write response: %v", err), http.StatusInternalServerError)
		return
	}
}
