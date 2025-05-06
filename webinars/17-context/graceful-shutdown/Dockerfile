FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /graceful-shutdown

EXPOSE 8080

# Run
CMD ["/graceful-shutdown"]
