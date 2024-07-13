FROM golang:1.22-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest &&\
    apk add --no-cache make npm nodejs ca-certificates &&\
    make



FROM alpine:latest as run

WORKDIR /app

COPY --from=build /app/main ./run

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["./run"]