FROM golang:1.20-alpine as builder

ENV APP_NAME $APP_NAME

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest 
RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN go build -o service cmd/service/main.go

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN task build

FROM scratch

COPY --from=builder ["/app/service", "/service"]

ENV GO_ENV=production

CMD ["/service"]

