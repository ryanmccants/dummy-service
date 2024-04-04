# base go image
FROM golang:1.21.4 as builder
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o dummy-service ./cmd/api
RUN chmod +x /app/dummy-service

# build a tiny docker image
FROM alpine:latest
RUN mkdir /app
COPY --from=builder /app/dummy-service /app
EXPOSE 9292
CMD [ "/app/dummy-service" ]