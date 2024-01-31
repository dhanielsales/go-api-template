FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o service cmd/service/main.go

FROM scratch

COPY --from=builder ["/app/service", "/service"]
COPY --from=builder ["/app/app.env", "/app.env"]

ENV GO_ENV=production
ENV PORT $PORT

EXPOSE $PORT

ENTRYPOINT ["./service"]
