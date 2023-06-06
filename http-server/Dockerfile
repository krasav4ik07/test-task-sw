FROM golang:1.18

WORKDIR ./src/http-server
COPY ./ ./

RUN go mod download
RUN go build -o http-server ./internal/main.go

ENTRYPOINT ["./http-server"]