FROM golang:1.21-alpine as builder

WORKDIR /app
ARG opts
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./src ./src

RUN env ${opts} go build -o /penca ./src/main.go

FROM alpine:latest
RUN apk add -U tzdata
ENV TZ=America/Montevideo
RUN cp /usr/share/zoneinfo/America/Montevideo /etc/localtime

COPY --from=builder /penca /app/bin/penca
COPY ./src/app.env /app.env
EXPOSE 8080
CMD [ "/app/bin/penca" ]
