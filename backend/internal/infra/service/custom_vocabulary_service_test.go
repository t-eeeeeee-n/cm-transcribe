package service

import (
	"cmTranscribe/internal/domain/model"
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomVocabulary_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// 生成されたモックを使用
	mockService := NewMockTranscribeServiceClient(mockCtrl)

	// モックの設定
	mockService.EXPECT().CreateVocabulary(gomock.Any()).Return(&transcribeservice.CreateVocabularyOutput{}, nil)

	// サービスを初期化
	service := &CustomVocabularyService{
		client: mockService,
	}

	// テストするカスタムボキャブラリ
	vocabulary := model.CustomVocabulary{
		VocabularyName: "test-vocab",
		LanguageCode:   "en-US",
		FileUri:        "s3://bucket/test-vocab.csv",
	}

	// 実行
	err := service.CreateCustomVocabulary(vocabulary)

	// アサーション
	assert.NoError(t, err)
}

func TestCreateCustomVocabulary_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// 生成されたモックを使用
	mockService := NewMockTranscribeServiceClient(mockCtrl)

	// モックの設定
	mockService.EXPECT().CreateVocabulary(gomock.Any()).Return(nil, errors.New("AWS error: Failed to create vocabulary"))

	// サービスを初期化
	service := &CustomVocabularyService{
		client: mockService,
	}

	// テストするカスタムボキャブラリ
	vocabulary := model.CustomVocabulary{
		VocabularyName: "test-vocab",
		LanguageCode:   "en-US",
		FileUri:        "s3://bucket/test-vocab.csv",
	}

	// 実行
	err := service.CreateCustomVocabulary(vocabulary)

	// アサーション
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create custom vocabulary")
}

func TestUpdateCustomVocabulary_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// 生成されたモックを使用
	mockService := NewMockTranscribeServiceClient(mockCtrl)

	// モックの設定
	mockService.EXPECT().UpdateVocabulary(gomock.Any()).Return(&transcribeservice.UpdateVocabularyOutput{}, nil)

	// サービスを初期化
	service := &CustomVocabularyService{
		client: mockService,
	}

	// テストするカスタムボキャブラリ
	vocabulary := model.CustomVocabulary{
		VocabularyName: "test-vocab",
		LanguageCode:   "en-US",
		FileUri:        "s3://bucket/test-vocab.csv",
	}

	// 実行
	err := service.UpdateCustomVocabulary(vocabulary)

	// アサーション
	assert.NoError(t, err)
}

func TestUpdateCustomVocabulary_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// 生成されたモックを使用
	mockService := NewMockTranscribeServiceClient(mockCtrl)

	// モックの設定
	mockService.EXPECT().UpdateVocabulary(gomock.Any()).Return(nil, errors.New("AWS error: Failed to update vocabulary"))

	// サービスを初期化
	service := &CustomVocabularyService{
		client: mockService,
	}

	// テストするカスタムボキャブラリ
	vocabulary := model.CustomVocabulary{
		VocabularyName: "test-vocab",
		LanguageCode:   "en-US",
		FileUri:        "s3://bucket/test-vocab.csv",
	}

	// 実行
	err := service.UpdateCustomVocabulary(vocabulary)

	// アサーション
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to update custom vocabulary")
}

func TestGetCustomVocabularyByName_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// 生成されたモックを使用
	mockService := NewMockTranscribeServiceClient(mockCtrl)

	// モックの設定
	mockService.EXPECT().GetVocabulary(gomock.Any()).Return(&transcribeservice.GetVocabularyOutput{
		VocabularyName:   aws.String("test-vocab"),
		LanguageCode:     aws.String("en-US"),
		DownloadUri:      aws.String("s3://bucket/test-vocab.csv"),
		VocabularyState:  aws.String("READY"),
		VocabularyLastModifiedTime: aws.Time(time.Now()),
	}, nil)

	// サービスを初期化
	service := &CustomVocabularyService{
		client: mockService,
	}

	// 実行
	result, err := service.GetCustomVocabularyByName("test-vocab")

	// アサーション
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test-vocab", result.VocabularyName)
	assert.Equal(t, "en-US", result.LanguageCode)
	assert.Equal(t, "s3://bucket/test-vocab.csv", result.FileUri)
	assert.Equal(t, "READY", result.VocabularyState)
	assert.Equal(t, time.Now(), result.VocabularyLastModifiedTime)
}

func TestGetCustomVocabularyByName_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// 生成されたモックを使用
	mockService := NewMockTranscribeServiceClient(mockCtrl)

	// モックの設定
	mockService.EXPECT().GetVocabulary(gomock.Any()).Return(nil, errors.New("AWS error: Vocabulary not found"))

	// サービスを初期化
	service := &CustomVocabularyService{
		client: mockService,
	}

	// 実行
	result, err := service.GetCustomVocabularyByName("nonexistent-vocab")

	// アサーション
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to get custom vocabulary")
}
