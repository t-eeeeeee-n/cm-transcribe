# API 使用説明書

## エンドポイント一覧

### 1. `/api/transcriptions/start` [POST]

- **説明**: 音声ファイルをAmazon Transcribeで文字起こしするためのジョブを作成します。
- **リクエストボディ**:
    - `jobName` (必須): ジョブ名
    - `mediaUri` (必須): S3バケットの音声ファイルへのURI。
    - `languageCode` (オプション): 文字起こしの言語コード。デフォルトは `en-US`。
- **リクエスト例**:

```bash
curl -X POST "http://localhost:8080/api/transcriptions/start" \
-H "Content-Type: application/json" \
-d @- <<'EOF'
{
  "jobName": "test",
  "mediaUri": "https://transcribe-test-a.s3.amazonaws.com/doda.mp3",
  "languageCode": "ja-JP",
  "customVocabularyName": "doda"
}
EOF
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

### 2. `/api/custom/vocabulary` [POST]

- **説明**: カスタムボキャブラリーを作成し、Amazon Transcribeに登録します。送信されたボキャブラリー情報をCSVファイルとして一時的に保存し、そのファイルをS3にアップロードした後、AWS Transcribeに登録します。
- **リクエストボディ**:
  - `name` (必須): 作成するカスタムボキャブラリーの名前。
  - `language` (必須): ボキャブラリーの言語コード。例: ja-JP。
  - `vocabularies` (必須): ボキャブラリーの語彙リスト。各語彙は以下のプロパティを持つ。
    - `phrase` (必須): フレーズ。
    - `soundsLike` (任意): みたいに聞こえるオプション。
    - `ipa` (任意): IPA（国際音声記号）オプション。
    - `displayAs` (任意): 表示オプション。
- **リクエスト例**:

```bash
curl -X POST "http://localhost:8080/api/custom/vocabulary" \
-H "Content-Type: application/json" \
-d @- <<'EOF'
{
  "name": "doda2",
  "language_code": "ja-JP",
  "vocabularies": [
    {
      "phrase": "パーソル",
      "soundsLike": "パアソル",
      "ipa": "",
      "displayAs": "PERSOL"
    },
    {
      "phrase": "デューダ",
      "soundsLike": "デュウダ",
      "ipa": "",
      "displayAs": "doda"
    }
  ]
}
EOF
```

- **レスポンス**:

```bash
{
  "message": "Custom vocabulary created successfully"
}
```

- **エラーレスポンス**:
  - **400 Bad Request**: 必要なパラメータが不足している場合や、JSONの形式が正しくない場合。
  - **500 Internal Server Error**: サーバー内部のエラーやAWSへのリクエストが失敗した場合。

### 3. `/api/custom/vocabulary` [PUT]

- **説明**: 既存のカスタムボキャブラリーに新しい語彙を上書きします。
- **リクエストボディ**:
  - `name` (必須): 更新するカスタムボキャブラリーの名前。
  - `language` (必須): ボキャブラリーの言語コード。例: ja-JP。
  - `vocabularies` (必須): 追加するボキャブラリーの語彙リスト。各語彙は以下のプロパティを持つ。
    - `phrase` (必須): フレーズ。
    - `soundsLike` (任意): みたいに聞こえるオプション。
    - `ipa` (任意): IPA（国際音声記号）オプション。
    - `displayAs` (任意): 表示オプション。
- **リクエスト例**:

```bash
curl -X PUT "http://localhost:8080/api/custom/vocabulary" \
-H "Content-Type: application/json" \
-d @- <<'EOF'
{
  "name": "MyVocabulary01",
  "language_code": "ja-JP",
  "vocabularies": [
    {
      "phrase": "新しいフレーズ",
      "soundsLike": "アタライイフレーズ",
      "ipa": "",
      "displayAs": "新しいフレーズ"
    }
  ]
}
EOF
```

- **レスポンス**:

```bash
{
  "message": "Custom vocabulary updated successfully"
}
```

- **エラーレスポンス**:
  - **400 Bad Request**: 必要なパラメータが不足している場合や、JSONの形式が正しくない場合。
  - **500 Internal Server Error**: サーバー内部のエラーやAWSへのリクエストが失敗した場合。

### 4. `/api/custom/vocabulary` [GET]

- **説明**: 指定した名前のカスタムボキャブラリーを取得します。Amazon Transcribeに登録されているカスタムボキャブラリーの詳細情報を返します。
- **クエリパラメータ**:
  - `name` (必須): 取得したいカスタムボキャブラリーの名前。
- **リクエスト例**:

```bash
curl -X GET "http://localhost:8080/api/custom/vocabulary?name=MyVocabulary01" \
-H "Content-Type: application/json"
```

- **レスポンス**:

```bash
{
  "VocabularyName": "MyVocabulary01",
  "LanguageCode": "ja-JP",
  "FileUri": "s3://bucket-name/path/to/vocabulary.csv",
  "VocabularyState": "READY"
}
```

- **エラーレスポンス**:
  - **400 Bad Request**: 必要なパラメータが不足している場合や、クエリパラメータが正しくない場合。
  - **404 Not Found**: 指定された名前のカスタムボキャブラリーが見つからない場合。
  - **500 Internal Server Error**:  サーバー内部のエラーやAWSへのリクエストが失敗した場合。


### 5. `/api/transcriptions/list` [GET]

- **説明**: 文字起こしジョブの一覧を取得します。Amazon Transcribeで実行したジョブのリストを返します。
- **リクエスト例**:

```bash
curl -X GET "http://localhost:8080/api/transcriptions/list" \
-H "Content-Type: application/json"
```

- **レスポンス**:

```bash
{
  "Jobs": [
    {
      "JobName": "transcription-job-id-1",
      "CreationTime": "2023-09-13T12:00:00Z",
      "CompletionTime": "2023-09-13T13:00:00Z",
      "LanguageCode": "ja-JP",
      "TranscriptionJobStatus": "COMPLETED",
      "OutputLocationType": "S3_BUCKET"
    },
    {
      "JobName": "transcription-job-id-2",
      "CreationTime": "2023-09-14T10:00:00Z",
      "CompletionTime": null,
      "LanguageCode": "en-US",
      "TranscriptionJobStatus": "IN_PROGRESS",
      "OutputLocationType": "S3_BUCKET"
    }
  ]
}
```

- **エラーレスポンス**:
  - **500 Internal Server Error**:  サーバー内部のエラーやAWSへのリクエストが失敗した場合。