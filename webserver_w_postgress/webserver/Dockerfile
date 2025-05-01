ARG GOBUILD_IMAGE=golang:1.24.2

FROM ${GOBUILD_IMAGE} AS builder
# These build args are passed by buildx automatically
ARG TARGETARCH

WORKDIR /go/src/workdir

COPY . .
RUN go mod download && go mod verify

RUN if [ "$TARGETARCH" = "arm64" ]; then \
      make build-arm64 ;\
    else \
      make build-amd64 ;\
    fi

FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /go/src/workdir/database/migrations database/migrations

COPY --from=builder /go/src/workdir/static static


COPY --from=builder /go/src/workdir/webserver webserver

EXPOSE 8080

CMD ["./webserver"]
