ARG GOBUILD_IMAGE

FROM ${GOBUILD_IMAGE} as builder

WORKDIR /go/src/workdir

COPY . .

RUN CGO_ENABLED=0 go build -a -o webserver main.go


FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /go/src/workdir/webserver webserver

CMD ["./webserver"]
