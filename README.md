# Go-Face-Recognition

Go-Face-Recognition is an facial recognition system, based on the principles of FaceNet and developed entirely in Go language. It leverages cutting-edge technology and utilizes the go-face library, which is built upon the powerful dlib C++ library for high-performance facial analysis.

## Table of Contents

- [Overview](#overview)
  - [About FaceNet](#about-facenet)
  - [About dlib](#about-dlib)
- [Key Features](#key-features)
- [Usage](#usage)
  - [Dynamic Loading of People](#dynamic-loading-of-people)
  - [Recognition of Faces](#recognition-of-faces)
  - [Output Generation](#output-generation)
  - [Capabilities](#capabilities)
- [Installation and Usage](#installation-and-usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## About FaceNet

Go-Face-Recognition is based on the principles of [FaceNet](https://arxiv.org/abs/1503.03832), a groundbreaking facial recognition system developed by Google. FaceNet employs a deep neural network to directly learn a mapping from facial images to a compact Euclidean space, where distances between embeddings correspond directly to a measure of facial similarity. By leveraging this learned embedding space, tasks such as facial recognition, verification, and clustering become straightforward, as FaceNet embeddings serve as feature vectors that capture essential facial characteristics. This integration enables Go-Face-Recognition to achieve state-of-the-art performance in facial recognition tasks, making it a versatile and powerful tool for various applications.

### About dlib

[dlib](http://dlib.net/) is a modern C++ toolkit containing machine learning algorithms and tools for creating complex software in C++ to solve real-world problems. It is renowned for its robustness, efficiency, and versatility in various applications, including computer vision, machine learning, and artificial intelligence.

## Key Features

- **Dynamic Person Loading:** Go-Face-Recognition dynamically loads people from within the 'persons' directory, enhancing flexibility.

- **Precise Face Recognition:** Utilizing go-face library powered by dlib, the system performs accurate and reliable face recognition, even in complex scenarios.

- **Easy Deployment with Docker:** To streamline dependency management and deployment, the project is encapsulated within a Docker environment, ensuring seamless integration into any development or production environment.

## Usage

### Dynamic Loading of People

This project dynamically loads people from within the `persons/` directory. Each person should have a subfolder with the person's name, containing images of that person to be used in the model. It is ideal to provide more than one image per person to improve classification accuracy. The images provided for each person should contain only one face, which is the face of the person.

### Recognition of Faces

After loading the people, the software reads an image from the `images/` directory. By default, it searches for an image named `unknown.jpg`. It then recognizes the faces in the image based on the provided people. The input image can contain multiple people, and the software attempts to recognize all of them.

### Output Generation

The output of the system is a new image with the faces marked and the name of each identified person. The generated image will be saved in the `images/` directory with the name `result.jpg`.

### Capabilities

The system enables effortless recognition of faces within images, empowering users with a powerful tool for various applications, including security, authentication, access control, and more.

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
docker compose up --build -d
```

4. Run the Docker container:

```bash
docker exec go-face-recognition go run cmd/main.go
```

This will execute the project inside the Docker container and demonstrate facial recognition.

## Project Structure

```
go-face-recognition/
├── images/
│ ├── result.jpg
│ └── unknown.jpg
├── persons/
│ ├── donald_trump/
│ │ └── ...
│ └── joe_biden/
│ └── ...
├── models/
│ └── ...
├── internal/
│ ├── entity/
│ └── usecases/
├── cmd/
│ └── main.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

The `images/` directory contains the input and output images. The `persons/` directory contains sub-folders for each person, with images of that person to be used in the model. The `models/` directory contains the trained model for facial recognition. The `internal/` directory contains the core logic of the system, including entities and use cases. The `cmd/` directory contains the main entry point of the system.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request on the GitHub repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
