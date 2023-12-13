FROM golang:1.21.3-bullseye
WORKDIR /builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

FROM alpine:latest
RUN apk update && apk add bash
WORKDIR /app
COPY --from=0 /builder/main ./
COPY ./config/config.yaml ./config/
COPY ./wait-for-it.sh .
CMD ["./main"]