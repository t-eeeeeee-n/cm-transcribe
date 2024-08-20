# API 使用説明書

## エンドポイント一覧

### 1. `/api/transcribe` [GET]

- **説明**: 音声ファイルをAmazon Transcribeで文字起こしするためのジョブを作成します。
- **パラメータ**:
    - `media_uri` (必須): S3バケットの音声ファイルへのURI。
    - `language` (オプション): 文字起こしの言語コード。デフォルトは `en-US`。
- **リクエスト例**:

```bash
curl -X GET "http://localhost:8080/api/transcribe?media_uri=https://transcribe-test-ten.s3.amazonaws.com/doda.mp3"
```

- **レスポンス**:

```bash
{
  "ID": "job-id",
  "MediaURI": "https://transcribe-test-ten.s3.amazonaws.com/doda.mp3",
  "Language": "en-US",
  "Status": "Pending",
  "CreatedAt": "2024-08-20T14:00:00Z"
}
```

- **エラーレスポンス**:
    - **400 Bad Request**: パラメータが不正または不足している場合。
    - **500 Internal Server Error**: サーバー内部のエラー。


## エラーコード
- **400**: リクエストのパラメータが不足または無効です。
- **500**: サーバー内部でエラーが発生しました。

