# goのバージョン指定
FROM golang:1.21.5-alpine3.18
# go moduleモード
ENV GO111MODULE on
# 必要なものをインストール
RUN apk update && \
  apk --no-cache add git && \
  apk --no-cache add bash
# 作業ディレクトリ変更
WORKDIR /go/src/app
# go.mod作成
RUN go mod init backend
# go.mod更新＆Air（ホットリロードライブラリ）インストール
RUN go mod tidy && \
  go install github.com/cosmtrek/air@v1.40.4
# Air実行
CMD ["air", "-c", ".air.toml"]
