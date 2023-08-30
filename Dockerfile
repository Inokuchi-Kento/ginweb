# ビルドステージ
FROM golang:1.20-alpine as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o ginweb

# 実行ステージ
FROM alpine:latest
COPY --from=build /app/ginweb /ginweb
CMD ["/ginweb"]