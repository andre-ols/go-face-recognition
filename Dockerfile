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
COPY ./fonts fonts
COPY ./persons persons

WORKDIR /app/cmd

RUN go mod tidy

WORKDIR /app

# Keep the container running
CMD ["tail", "-f", "/dev/null"]
