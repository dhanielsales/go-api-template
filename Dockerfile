FROM golang:1.20-alpine as builder

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
COPY . .
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o go-api-template cmd/go-api-template/main.go

FROM scratch

COPY --from=builder ["/etc/ssl/certs/ca-certificates.crt", "/etc/ssl/certs/"]
COPY --from=builder ["/app/go-api-template", "/go-api-template"]

ENV GO_ENV=production
ENV PORT $PORT

EXPOSE $PORT

CMD ["./go-api-template"]
