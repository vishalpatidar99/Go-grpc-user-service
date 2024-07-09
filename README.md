# Golang gRPC User Service with Search

## Introduction

This project implements a Golang gRPC service that retrieves user details and provides search capabilities based on specified criteria.

## Objective
- Retrieve user details by user ID.
- Retrieve a list of user details based on a list of user IDs.
- Implement a search functionality to find user details based on criteria such as city, phone number, marital status, etc.

## Prerequisites
Ensure you have the following installed on your development environment:
- Go programming language (prefer official documentation to [install](https://go.dev/doc/install)) 
- Docker (if you plan to containerize the application)
- grpcurl (for testing gRPC endpoints [official site](https://github.com/fullstorydev/grpcurl))

### I followed these simple steps to download grpcurl in my wsl2, will work for linux environment too
1. Download repository:
    ```bash
    wget https://github.com/fullstorydev/grpcurl/releases/download/v1.9.1/grpcurl_1.9.1_linux_x86_64.tar.gz
    ```

2. Extract file:
    ```bash
    tar -zxvf grpcurl_1.9.1_linux_x86_64.tar.gz
    ```

3. Move to bin folder for being executalbe:
    ```bash
    sudo mv grpcurl /usr/local/bin/
    ```

4. Check the installation:
    ```bash
    grpcurl --version
   ```


## Docker Setup

### Build and Run Docker Container

1. Clone the repository:
   ```bash
   git clone https://github.com/vishalpatidar99/Go-grpc-user-service.git
   cd Go-grpc-user-service
   ```
2. Run docker compose:
    ```bash
    docker-compose up -d --build
    ```

## Local Setup
 1. Clone the repository:
    follow same step mention in docker setup

2. Build binary:
    ```bash
    go build -o user-service main.go
    ```
3. Run grpc service:
   ```bash
    ./user-service
    ```
  
 ## Testing Details
 ### Run Unit Test
 ```bash
 go test -v ./tests
 ```
 
 ### Test using gRPC client written in Go
 you can update request params as needed
 ```bash
 go run tests/mock-client/client.go
 ```
 
 ### Test using grpcurl with some simple commands
 1. GetUserByID Endpoint
    To retrieve a user by requested ID:

 ```bash
 grpcurl -plaintext -d '{"id": 1}' localhost:50051 user.UserService/GetUserByID
 ```
 
 2. GetUsersByIDs Endpoint
    To retrieve a list of users by their IDs:

 ```bash
 grpcurl -plaintext -d '{"ids": [1, 2]}' localhost:50051 user.UserService/GetUsersByIDs
 ```
 
 3. SearchUser Endpoint
    To search a user by request criteria:

```bash
grpcurl -plaintext -d '{"married": "true"}' localhost:50051 user.UserService/SearchUsers
```
