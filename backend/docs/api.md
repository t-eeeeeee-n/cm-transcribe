# API 使用説明書

## エンドポイント一覧

### 1. `/api/transcriptions/create` [POST]

- **説明**: 音声ファイルをAmazon Transcribeで文字起こしするためのジョブを作成します。
- **リクエストボディ**:
    - `media_uri` (必須): S3バケットの音声ファイルへのURI。
    - `language` (オプション): 文字起こしの言語コード。デフォルトは `en-US`。
- **リクエスト例**:

```bash
curl -X POST "http://localhost:8080/api/transcriptions/start" \
-H "Content-Type: application/json" \
-d @- <<'EOF'
{
  "media_uri": "https://transcribe-test-a.s3.amazonaws.com/doda.mp3",
  "language_code": "ja-JP",
  "custom_vocabulary_name": "doda"
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

### 2. `/api/custom_vocabulary` [POST]

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
curl -X POST "http://localhost:8080/api/custom_vocabulary/create" \
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

### 3. `/api/custom_vocabulary/update` [POST]

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
curl -X POST "http://localhost:8080/api/custom_vocabulary/update" \
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