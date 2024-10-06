FROM golang:1.23-alpine AS builder

WORKDIR /go/src/github.com/yyh-gl/gatling

ARG VERSION

ENV TZ="Asia/Tokyo"
ENV GOOS="linux"
ENV GOARCH="amd64"
ENV CGO_ENABLED=0

COPY . .

RUN go build -ldflags '-X main.version=$(version)' -o ./gatling . 

FROM gcr.io/distroless/base

COPY --from=builder /go/src/github.com/yyh-gl/gatling/gatling /app/gatling

CMD ["/app/gatling"]
