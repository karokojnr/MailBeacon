FROM golang:1.22-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest &&\
    apk add make npm nodejs &&\
    make



FROM alpine:latest as run

WORKDIR /app
COPY --from=build /app/main ./run

RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["./run"]