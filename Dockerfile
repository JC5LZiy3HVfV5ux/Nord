# syntax=docker/dockerfile:1
FROM golang:1.17.8 as build
WORKDIR /root
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/cache-server/main.go

FROM alpine:latest as dev
WORKDIR /root
COPY --from=build /root/bin/app .
EXPOSE 5000
CMD ["./app"]