# syntax=docker/dockerfile:1
FROM golang:1.21-bookworm AS builder

WORKDIR /src/

COPY go.mod go.sum ./

ENV GOPRIVATE=github.com/Iknite-Space
ENV HOME=/root
RUN apt-get update && apt-get install -y ca-certificates git-core ssh
RUN git config --global url."git@github.com:Iknite-Space/".insteadOf "https://github.com/Iknite-Space/"
ADD ./temp_home/.ssh/id_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "\n\nHost github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

RUN go mod download -x

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/api ./cmd/api/api.go



FROM golang:1.21-bookworm

RUN apt update && apt install ca-certificates -y

WORKDIR /app/bin

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
