package service

import (
	"cmTranscribe/internal/app/dto"
	"cmTranscribe/internal/domain/model"
	"cmTranscribe/internal/domain/service"
	"cmTranscribe/internal/infra/config"
	"cmTranscribe/internal/shared/validator"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
)

// CustomVocabularyService アプリケーション層のサービス
type CustomVocabularyService struct {
	CustomVocabularyService service.CustomVocabularyService
	FileService             service.FileService
	S3StorageService        service.S3StorageService
}

// NewCustomVocabularyService 新しい CustomVocabularyService を作成します
func NewCustomVocabularyService(
	customVocabularyService service.CustomVocabularyService,
	fileService service.FileService,
	s3StorageService service.S3StorageService,
) *CustomVocabularyService {
	return &CustomVocabularyService{
		CustomVocabularyService: customVocabularyService,
		FileService:             fileService,
		S3StorageService:        s3StorageService,
	}
}

// CreateCustomVocabulary カスタムボキャブラリを作成します
func (s *CustomVocabularyService) CreateCustomVocabulary(ctx context.Context, request dto.CreateVocabularyDto) error {
	// DTOからドメインモデルに変換
	entries := model.ConvertEntriesToContent(request.Vocabularies)
	// ファイルパス作成
	filePath := model.GenerateFilePath(request.VocabularyName)

	// ドメインモデルを作成
	csvFile := model.NewCSVFile(request.VocabularyName, filePath, entries)
	// バリデーションの実行
	if err := validator.Validate(csvFile); err != nil {
		// エラーハンドリングのみ行う
		return fmt.Errorf("error processing csvFile: %v", err)
	}

	// CSVファイルを作成
	filePath, fileHandle, err := s.FileService.CreateCSV(*csvFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := fileHandle.Close(); err != nil {
			fmt.Printf("Failed to close file: %v\n", err)
		}
	}()
	defer func() {
		if err := s.FileService.Cleanup(filePath); err != nil {
			fmt.Printf("Failed to clean up file: %v\n", err)
		}
	}()

	// ドメインモデルを作成
	s3File := model.NewS3File(filePath, config.AppConfig.S3BucketName, config.AppConfig.S3PrefixVocabulary)
	// バリデーションの実行
	if err := validator.Validate(s3File); err != nil {
		// エラーハンドリングのみ行う
		return fmt.Errorf("error processing s3File: %v", err)
	}

	// S3にファイルをアップロード
	s3Uri, err := s.S3StorageService.UploadToS3(ctx, *s3File)
	if err != nil {
		return fmt.Errorf("failed to create custom vocabulary: %v", err)
	}

	// ドメインモデルを作成
	customVocabulary := model.NewCustomVocabulary(request.VocabularyName, request.LanguageCode, s3Uri)
	// バリデーションの実行
	if err := validator.Validate(customVocabulary); err != nil {
		// エラーハンドリングのみ行う
		return fmt.Errorf("error processing customVocabulary: %v", err)
	}

	// ドメインサービスを使ってカスタムボキャブラリを作成
	err = s.CustomVocabularyService.CreateCustomVocabulary(ctx, *customVocabulary)
	if err != nil {
		return fmt.Errorf("failed to create custom vocabulary: %v", err)
	}
	return nil
}

// UpdateCustomVocabulary 既存のカスタムボキャブラリを更新します。
func (s *CustomVocabularyService) UpdateCustomVocabulary(ctx context.Context, request dto.UpdateVocabularyDto) error {
	// DTOからドメインモデルに変換
	entries := model.ConvertEntriesToContent(request.Vocabularies)
	// ファイルパス作成
	filePath := model.GenerateFilePath(request.VocabularyName)

	// ドメインモデルを作成
	csvFile := model.NewCSVFile(request.VocabularyName, filePath, entries)
	// バリデーションの実行
	if err := validator.Validate(csvFile); err != nil {
		// エラーハンドリングのみ行う
		return fmt.Errorf("error processing csvFile: %v", err)
	}

	// CSVファイルを作成
	filePath, fileHandle, err := s.FileService.CreateCSV(*csvFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := fileHandle.Close(); err != nil {
			fmt.Printf("Failed to close file: %v\n", err)
		}
	}()
	defer func() {
		if err := s.FileService.Cleanup(filePath); err != nil {
			fmt.Printf("Failed to clean up file: %v\n", err)
		}
	}()

	// ドメインモデルを作成
	s3File := model.NewS3File(filePath, config.AppConfig.S3BucketName, config.AppConfig.S3PrefixVocabulary)
	// バリデーションの実行
	if err := validator.Validate(s3File); err != nil {
		// エラーハンドリングのみ行う
		return fmt.Errorf("error processing s3File: %v", err)
	}

	// S3にファイルをアップロード
	s3Uri, err := s.S3StorageService.UploadToS3(ctx, *s3File)
	if err != nil {
		return fmt.Errorf("failed to update custom vocabulary: %v", err)
	}

	// ドメインモデルを作成
	customVocabulary := model.NewCustomVocabulary(request.VocabularyName, request.LanguageCode, s3Uri)
	// バリデーションの実行
	if err := validator.Validate(customVocabulary); err != nil {
		// エラーハンドリングのみ行う
		return fmt.Errorf("error processing customVocabulary: %v", err)
	}

	// ドメインサービスを使ってカスタムボキャブラリを作成
	err = s.CustomVocabularyService.UpdateCustomVocabulary(ctx, *customVocabulary)
	if err != nil {
		return fmt.Errorf("failed to update custom vocabulary: %v", err)
	}
	return nil
}

// GetCustomVocabularyByName 名前でカスタムボキャブラリを取得し、クライアントに返す形式に変換します
func (s *CustomVocabularyService) GetCustomVocabularyByName(ctx context.Context, name string) (*dto.CustomVocabularyResponse, error) {
	// ドメインサービスを使ってデータを取得
	customVocab, err := s.CustomVocabularyService.GetCustomVocabularyByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("failed to get custom vocabulary: %v", err)
	}

	// DownloadUriから内容をダウンロード
	vocabularies, err := s.downloadAndParseVocabularyFile(customVocab.FileUri)
	if err != nil {
		return nil, fmt.Errorf("failed to download and parse vocabulary file: %v", err)
	}

	// 結果の構築
	response := &dto.CustomVocabularyResponse{
		VocabularyName:             customVocab.VocabularyName,
		LanguageCode:               customVocab.LanguageCode,
		Vocabularies:               vocabularies,
		VocabularyState:            customVocab.VocabularyState,
		VocabularyLastModifiedTime: customVocab.VocabularyLastModifiedTime,
	}

	return response, nil
}

// downloadAndParseVocabularyFile ダウンロードしてパースする
func (s *CustomVocabularyService) downloadAndParseVocabularyFile(uri string) ([]dto.Vocabulary, error) {
	// HTTPリクエストを使用してファイルをダウンロード
	resp, err := http.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Failed to close file: %v\n", err)
		}
	}()

	// ダウンロードした内容をパース
	var vocabularies []dto.Vocabulary
	reader := csv.NewReader(resp.Body)
	reader.Comma = '\t' // ファイルの区切り文字がタブであることを想定

	// ヘッダーをスキップ
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("failed to read header: %v", err)
	}

	// 各行をパースして構造体にマッピング
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to parse CSV: %v", err)
		}

		vocabulary := dto.Vocabulary{
			Phrase:     record[0],
			IPA:        record[1],
			SoundsLike: record[2],
			DisplayAs:  record[3],
		}
		vocabularies = append(vocabularies, vocabulary)
	}

	return vocabularies, nil
}
