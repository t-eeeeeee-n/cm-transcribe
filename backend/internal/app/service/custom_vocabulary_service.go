package service

import (
	"cmTranscribe/internal/app/dto"
	"cmTranscribe/internal/domain/model"
	"cmTranscribe/internal/domain/service"
	"cmTranscribe/internal/infra/config"
	"cmTranscribe/internal/shared/validator"
	"fmt"
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
func (s *CustomVocabularyService) CreateCustomVocabulary(request dto.CreateVocabularyDto) error {
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
	s3File := model.NewS3File(filePath, config.AppConfig.S3BucketName, "custom-vocabularies")
	// バリデーションの実行
	if err := validator.Validate(s3File); err != nil {
		// エラーハンドリングのみ行う
		return fmt.Errorf("error processing s3File: %v", err)
	}

	// S3にファイルをアップロード
	s3Uri, err := s.S3StorageService.UploadToS3(*s3File)
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
	return s.CustomVocabularyService.CreateCustomVocabulary(*customVocabulary)
}

// UpdateCustomVocabulary 既存のカスタムボキャブラリを更新します。
func (s *CustomVocabularyService) UpdateCustomVocabulary(request dto.UpdateVocabularyDto) error {
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
	s3File := model.NewS3File(filePath, config.AppConfig.S3BucketName, "custom-vocabularies")
	// バリデーションの実行
	if err := validator.Validate(s3File); err != nil {
		// エラーハンドリングのみ行う
		return fmt.Errorf("error processing s3File: %v", err)
	}

	// S3にファイルをアップロード
	s3Uri, err := s.S3StorageService.UploadToS3(*s3File)
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
	return s.CustomVocabularyService.UpdateCustomVocabulary(*customVocabulary)
}
