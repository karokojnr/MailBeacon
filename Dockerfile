FROM golang:1.22.4 AS build-stage
WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /api /app/cmd/api/main.go


FROM scratch AS production-stage
WORKDIR /app

COPY --from=build-stage /api /api

EXPOSE 8080

CMD [ "/api" ]