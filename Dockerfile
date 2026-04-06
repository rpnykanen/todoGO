### Build
FROM golang:1.26-alpine AS build-stage
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY go.mod ./

COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /app .

### Run
FROM alpine:3.23.3 AS final

COPY --from=build-stage /app /bin/app

EXPOSE 8080

CMD ["bin/app"]