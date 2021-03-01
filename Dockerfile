FROM golang:1.15.8-buster as builder
ENV DATA_DIRECTORY /go/src/github.com/laybatin/bodymirror-app-backend
WORKDIR $DATA_DIRECTORY
ARG APP_VERSION
ARG CGO_ENABLED=0
COPY . .
RUN go build -o cmd/bodymirror-server -ldflags="-X github.com/laybatin/bodymirror-app-backend/internal/config.Version=$APP_VERSION" github.com/laybatin/bodymirror-app-backend/cmd/server

FROM alpine:3.10
ENV DATA_DIRECTORY /go/src/github.com/laybatin/bodymirror-app-backend/
RUN apk add --update --no-cache \
    ca-certificates
COPY --from=builder ${DATA_DIRECTORY}/cmd/bodymirror-server /bodymirror-app-backend

ENTRYPOINT ["/bodymirror-app-backend"]
