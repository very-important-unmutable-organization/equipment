FROM golang:1.17-rc-buster as builder

WORKDIR /go/src/app

RUN wget -O swag_linux_amd64.tar.gz https://github.com/swaggo/swag/releases/download/v1.7.1/swag_linux_amd64.tar.gz \
 && tar -xzf swag_linux_amd64.tar.gz \
 && mv swag_linux_amd64/swag  /usr/local/bin \
 && rm swag_linux_amd64.tar.gz

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN swag init -g ./cmd/main.go

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

CMD ["./equipment-app"]
