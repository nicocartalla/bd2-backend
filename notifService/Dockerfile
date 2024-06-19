FROM golang:1.21-alpine as builder

WORKDIR /app
ARG opts
COPY go.mod ./
COPY go.sum ./
COPY ./src ./src

RUN go mod download
RUN env ${opts} go build -o /notif ./src/main.go

FROM alpine:latest
RUN apk add -U tzdata
ENV TZ=America/Montevideo
RUN cp /usr/share/zoneinfo/America/Montevideo /etc/localtime

COPY --from=builder /notif /app/bin/notif
COPY docker.env /app.env
CMD [ "/app/bin/notif" ]
