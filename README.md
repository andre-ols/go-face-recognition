# Facial Recognition in Golang

This is a facial recognition project that attempts to predict whether a person is Barack Obama, Joe Biden, or an unknown person. The project was built in Go and utilizes the go-face library for facial recognition. Since the go-face library is a wrapper for the dlib library, which is a tool for machine learning and computer vision developed in C++, to facilitate deployment and ensure portability, the entire project runs in a Docker container.

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
docker build -t go-face-recognition .
```

4. Run the Docker container:

```bash
docker run go-face-recognition go run cmd/main.go
```

This will execute the project inside the Docker container and demonstrate facial recognition.

## Implementation Details

The project uses the go-face library to perform facial recognition. The library is a wrapper around the dlib library, which is a toolkit for machine learning and computer vision. The go-face library provides a simple API for facial recognition and is easy to use.

## Project Structure

go-face-recognition/
│
├── images/
│ ├── obama.jpg
│ ├── biden.jpg
│ └── unknown.jpg
└── models/
│ └── ...
├── internal/
│ |── entity/
│ └──── person.go
├── cmd/
│ └── main.go
├── Dockerfile
├── go.mod
├── go.sum
└── README.md

The `images/` directory contains images of Barack Obama, Joe Biden, and an unknown person. The `models/` directory contains the pre-trained model used for facial recognition.

`Dockerfile` contains the instructions for building the Docker image for the project.

## Legal Disclaimer

This project is for educational purposes only and should not be used for any illegal activities. The author is not responsible for any misuse of the information provided.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
