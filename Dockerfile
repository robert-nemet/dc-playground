FROM golang:latest as builder

WORKDIR /app
COPY . /app/

RUN go mod tidy
RUN go build -o app

FROM golang:buster

WORKDIR /app

COPY --from=builder /app/app /app/

ENTRYPOINT [ "/app/app" ]
