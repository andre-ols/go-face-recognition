FROM golang:1.22.1

# Install go-face dependencies
RUN apt-get update && apt-get -y install \
    libdlib-dev \
    libblas-dev \
    libatlas-base-dev \
    liblapack-dev \
    libjpeg62-turbo-dev

# Set the working directory
WORKDIR /app

# Copy go modules files
COPY ./go.mod .
COPY ./go.sum .

# Copy the source code
COPY ./internal/ internal/
COPY ./models/ models/
COPY ./images/ images/
COPY ./cmd/ cmd/

WORKDIR /app/cmd

RUN go mod tidy

# WORKDIR /app

# Compile for Linux
# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o face-recognition ./cmd/main.go

# Keep the container running
CMD ["tail", "-f", "/dev/null"]
