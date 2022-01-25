FROM golang:1.17

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", "air.toml"]