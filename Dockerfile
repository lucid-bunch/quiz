FROM golang:1.12-alpine as builder
WORKDIR /build/
COPY . .
RUN apk add --update git && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w -extldflags "-static"' -o app

FROM scratch
USER 1000
COPY --from=builder /build/app .
COPY --from=builder /build/quiz.csv .
CMD ["./app"]
