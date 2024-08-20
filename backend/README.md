### ディレクトリ構成

````markdown
backend/
├── cmd/                         # エントリーポイント（アプリケーションの起動コード）
│   └── server/
│       └── main.go              # サーバーの起動、依存関係の設定
├── internal/                    # 内部実装用（Go特有のディレクトリ、外部からのアクセスは不可）
│   ├── app/                     # アプリケーション層（ユースケース、サービス）
│   │   ├── service/             # アプリケーションサービス
│   │   │   └── transcription.go # 文字起こしに関連するユースケースロジック
│   │   ├── dto/                 # データ転送オブジェクト（DTO）
│   │   └── command/             # ユースケースコマンド
│   ├── domain/                  # ドメイン層（ビジネスロジック、エンティティ、リポジトリインターフェース）
│   │   ├── model/               # ドメインモデル（エンティティ、値オブジェクト）
│   │   │   └── transcription.go # 文字起こしに関連するドメインモデル
│   │   ├── repository/          # リポジトリインターフェース
│   │   │   └── transcription_repository.go
│   │   └── service/             # ドメインサービス
│   │       └── transcription_service.go
│   ├── infra/                   # インフラストラクチャ層（DB、外部サービス、リポジトリの実装）
│   │   ├── persistence/         # 永続化（リポジトリの実装）
│   │   │   └── transcription_repository.go
│   │   ├── external/            # 外部APIやサービスへの接続
│   │   └── config/              # 設定管理
│   │       └── config.go
│   ├── interface/               # インターフェース層（API、ハンドラ、コントローラ、プレゼンテーション）
│   │   ├── api/                 # Web APIエンドポイント
│   │   │   └── transcription_handler.go
│   │   ├── grpc/                # gRPCサービス
│   │   └── repository/          # リポジトリのインターフェース（実装はinfra層）
│   └── shared/                  # 共通のユーティリティ、ヘルパー、共通モジュール
│       ├── utils/               # 汎用的なユーティリティ
│       └── middleware/          # ミドルウェア
├── pkg/                         # 外部に公開可能なライブラリ（内部ロジックから分離）
│   └── somepkg/                 # パッケージ例
├── api/                         # OpenAPI/SwaggerやgRPCプロトコルの定義
│   ├── openapi.yaml             # API定義ファイル
│   └── protobuf/                # gRPCのprotobufファイル
├── scripts/                     # スクリプト（セットアップ、デプロイ、メンテナンス）
│   ├── setup.sh                 # 初期セットアップスクリプト
│   └── deploy.sh                # デプロイスクリプト
├── docs/                        # ドキュメント（設計書、使用説明書）
│   ├── architecture.md          # アーキテクチャの概要
│   └── api.md                   # APIの使用説明書
├── .gitignore                   # Gitで無視するファイルリスト
├── go.mod                       # Goモジュールファイル
└── go.sum                       # 依存関係のバージョン管理ファイル

````    

### セットアップ

1. Goの依存関係をインストール
    ```bash
    go mod tidy
    ```

2. サーバー起動
    ```bash
    go run ./cmd/server/main.go
    ```