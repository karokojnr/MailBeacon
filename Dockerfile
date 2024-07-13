FROM golang:1.22-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -v -o /api cmd/api/main.go

RUN go install github.com/a-h/templ/cmd/templ@latest &&\
    apk add make npm nodejs &&\
    make



FROM scratch as run

WORKDIR /app
COPY --from=build /app/main ./run

EXPOSE 8080

CMD ["./run"]


# FROM scratch AS production-stage

# WORKDIR /app

# COPY --from=build /api /api


# EXPOSE 8080

# CMD ["/api"]