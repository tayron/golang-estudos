FROM golang:1.18
WORKDIR /app

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install upx -y

COPY . ./

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.53.3
RUN ./bin/golangci-lint --version
RUN ./bin/golangci-lint run ./...

RUN go test ./...

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -o golangci -ldflags "-s -w" ./cmd/main.go
RUN upx crawler-vtweb

RUN ./golangci