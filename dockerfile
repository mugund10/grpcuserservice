FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o userservice ./cmd/server
EXPOSE 9292
CMD ["./userservice"]