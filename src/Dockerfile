FROM golang:1.17-rc-buster as builder

WORKDIR /go/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 go build -v -o equipment-app ./cmd/main.go

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait && chmod +x ./equipment-app

FROM debian:stable-slim

RUN apt-get update && \
    apt-get install -y ca-certificates apache2-utils && \
    rm -rf /var/cache/apt/*

WORKDIR /go/src/app

COPY --from=builder /go/src/app/ /go/src/app/
COPY --from=builder /wait /wait
COPY start.sh /go/src/app
CMD ["./start.sh"]
