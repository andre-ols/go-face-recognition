# Facial Recognition in Golang using go-face

This is a project in Go that demonstrates facial recognition using the go-face library. The entire project runs in a Docker container to ensure portability and ease of deployment.

## Requirements

Make sure you have Docker installed on your machine before proceeding.

## Installation and Usage

1. Clone this repository:

```bash
git clone https://github.com/andre-ols/go-face-recognition.git
```

2. Navigate to the project directory:

```bash
cd go-face-recognition
```

3. Build the Docker image:

```bash
docker compose up  --build -d
```

4. Run the Docker container:

```bash
docker exec go-face-recognition go run cmd/main.go go run internal/cmd/main.go
```

This will execute the project inside the Docker container and demonstrate facial recognition.

## Implementation Details

The project uses the go-face library to perform facial recognition. The library is a wrapper around the dlib library, which is a toolkit for machine learning and computer vision. The go-face library provides a simple API for facial recognition and is easy to use.

## Project Structure

go-face-recognition/
│
├── testdata/
│ ├── images/
│ │ ├── obama.jpg
│ │ ├── biden.jpg
│ │ └── unknown.jpg
│ └── models/
│ └── ...
├── internal/
│ |── entity/
│ └──── person.go
├── cmd/
│ └── main.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md

`testdata/` contains images used for testing the facial recognition model. The `images/` directory contains images of Barack Obama, Joe Biden, and an unknown person. The `models/` directory contains the pre-trained model used for facial recognition.

`Dockerfile` contains the instructions for building the Docker image for the project. `docker-compose.yml` contains the configuration for running the Docker container.

## Legal Disclaimer

This project is for educational purposes only and should not be used for any illegal activities. The author is not responsible for any misuse of the information provided.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
