#Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

#Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY db/migration ./migration
COPY startup.sh .
COPY wait-for.sh .
EXPOSE 8080
# If CMD works with ENTRYPOINT, CMD will pass to ENRTYPOINT as an extra command, equal to : ENTRYPOINT ["/app/main", "/app/main"]
CMD ["/app/main"]
ENTRYPOINT ["/app/startup.sh"]
