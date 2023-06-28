FROM ubuntu:20.04

# Set destination for COPY
WORKDIR /app/bin

# specif location to copy go binary
COPY ./bin/api ./
COPY ./db ./
ENV DB_MIGRATIONS_PATH=./db/migrations

ENV LISTEN_PORT=9081
EXPOSE  9980

# Run
CMD ["/app/bin/api" ]
