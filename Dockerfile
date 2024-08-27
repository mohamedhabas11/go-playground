### Stage 1: Build the go binary
FROM golang:1.20-alpine AS builder

#### Set working directory inside the container
WORKDIR /app

#### Copy Go tracking files
COPY go.mod ./

#### Install Dependancies
RUN go mod download

#### Copy the source code into the container
COPY . .

