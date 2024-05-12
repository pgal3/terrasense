FROM golang:1.22 AS build

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o terrasense_server cmd/main.go

# # Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:latest AS release

WORKDIR /

COPY --from=build /app/terrasense_server /terrasense_server

# EXPOSE 3000

ENTRYPOINT ["/terrasense_server", "--prod"]