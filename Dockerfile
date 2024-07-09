FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o go-grpc-user-service main.go

EXPOSE 50051
CMD ["./go-grpc-user-service"]
