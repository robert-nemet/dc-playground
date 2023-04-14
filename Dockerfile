ARG GO_VERSION=1.20.3
FROM golang:${GO_VERSION}-buster as builder

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download -x

COPY . /app/
RUN go build -o app

FROM gcr.io/distroless/base-debian11

USER nonroot:nonroot
WORKDIR /app

COPY --chown=nonroot:nonroot --from=builder /app/app /app/

ENTRYPOINT [ "/app/app" ]