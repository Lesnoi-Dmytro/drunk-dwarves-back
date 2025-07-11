FROM golang:1.24.4-alpine AS bulder

WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/go-jet/jet/v2/cmd/jet@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.* ./

RUN go mod download

COPY . .

RUN goose up
RUN CGO_ENABLED=0 GOOS=linux go build -o /gen-jet ./cmd/generate/jet
RUN /gen-docs

RUN CGO_ENABLED=0 GOOS=linux go build -o /gen-docs ./cmd/generate/docs
RUN /gen-docs

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api

EXPOSE 8080

CMD ["/api"]
