FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o app ./cmd/web

FROM gcr.io/distroless/static-debian12
COPY --from=build /src/app /app
COPY --from=build /src/web /web
EXPOSE 8080
ENTRYPOINT [ "/app" ]