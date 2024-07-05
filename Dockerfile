# Use the official Golang image as the base image
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /barber_book

# Copy the local package files to the container's workspace
COPY . /barber_book

# Build the Go application
RUN go build -o barber-book ./main.go

EXPOSE 8080:8080

CMD [ "./barber-book" ]