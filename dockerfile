# syntax=docker/dockerfile:1
FROM golang:1.20-alpine AS builder

WORKDIR /src/

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/api ./cmd/api/api.go



FROM alpine:latest  

RUN apk --no-cache  add ca-certificates

WORKDIR /root/

# copy over binary build from the first builder stage above into ./
COPY --from=builder /src/bin/api ./

# migration files
COPY ./db ./

ENV DB_MIGRATIONS_PATH=./db/migrations

ENV LISTEN_PORT=9080

EXPOSE  9980

# Run
CMD ["./api"]

# ENTRYPOINT [ "/api" ]