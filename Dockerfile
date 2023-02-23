
FROM golang:1.19-alpine
# アップデートとgitのインストール
RUN apk update && apk add git

# ホットリロード
RUN go install github.com/cosmtrek/air@latest

# goのREPL
RUN go install github.com/x-motemen/gore/cmd/gore@latest

# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

# air & apiサーバー起動
CMD ["air", "-c", ".air.toml"]
