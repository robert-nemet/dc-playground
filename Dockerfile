FROM golang:latest as builder

WORKDIR /app
COPY go.mod /app/
COPY main.go /app/
ADD internal /app/


RUN go mod download
RUN go build -o app

FROM golang:buster

WORKDIR /app

COPY --from=builder /app/app /app/

EXPOSE 9999

ENTRYPOINT [ "/app/app" ]
