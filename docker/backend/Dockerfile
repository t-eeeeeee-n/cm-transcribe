FROM golang:1.20 as build

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.29.0

# 依存関係ファイルをコピー
COPY ./backend/go.mod ./backend/go.sum /app/
# 依存関係をダウンロード
RUN go mod download

# ソースコードをコンテナにコピー
COPY ./backend /app

# 開発サーバーの起動
CMD ["air", "-c", ".air.toml"]