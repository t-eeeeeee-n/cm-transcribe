package api

import (
	"cmTranscribe/internal/app/service"
	"cmTranscribe/internal/infra/config"
	"cmTranscribe/internal/shared/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type S3UploadHandler struct {
	uploadService *service.S3UploadService
}

func NewS3UploadHandler(uploadService *service.S3UploadService) *S3UploadHandler {
	return &S3UploadHandler{uploadService: uploadService}
}

func (h *S3UploadHandler) HandleUploadToS3(w http.ResponseWriter, r *http.Request) {
	// 1. FormDataからファイルを取得
	file, handler, err := r.FormFile("file")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to parse file from form data")
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Failed to close file: %v\n", err)
		}
	}()

	// 2. 一時ディレクトリにファイルを保存
	tempFilePath, err := h.saveFileToTempDirectory(file, handler.Filename)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer func() {
		if err := os.Remove(tempFilePath); err != nil {
			fmt.Printf("Failed to remove temp file: %v\n", err)
		}
	}()

	// 3. サービス層でアップロードを処理
	url, err := h.uploadService.UploadToS3(r.Context(), tempFilePath, config.AppConfig.S3BucketName, config.AppConfig.S3PrefixUploadFile)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to upload file to S3: "+err.Error())
		return
	}

	// 4. JSON形式でレスポンスを返す
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"url": url})
}

// saveFileToTempDirectory は、ファイルを一時ディレクトリに保存し、そのパスを返します。
func (h *S3UploadHandler) saveFileToTempDirectory(file io.Reader, fileName string) (string, error) {
	tempFilePath := filepath.Join(os.TempDir(), fileName)
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}
	defer func() {
		if err := tempFile.Close(); err != nil {
			fmt.Printf("Failed to close file: %v\n", err)
		}
	}()

	// ファイルを一時ディレクトリにコピー
	if _, err := io.Copy(tempFile, file); err != nil {
		return "", fmt.Errorf("failed to save file to temp directory: %v", err)
	}

	return tempFilePath, nil
}
