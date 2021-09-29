FROM golang:1.16-stretch AS builder

WORKDIR /app
COPY . ./
RUN go mod tidy
RUN go build -o /mainrun

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /mainrun /app/
WORKDIR /app

STOPSIGNAL SIGINT
EXPOSE 8080
CMD [ "/mainrun" ]


