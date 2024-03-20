package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/Kagami/go-face"
	"github.com/andre-ols/go-face-recognition/internal/entity"
	"github.com/andre-ols/go-face-recognition/internal/usecases"
)

var (
	modelsDir  = "models"
	imagesDir  = "images"
	personsDir = "persons"
)

// This example shows the basic usage of the package: create an
// recognizer, recognize faces, classify them using few known ones.
func main() {

	initialTime := time.Now()

	defer func() {
		fmt.Printf("\x1b[34mTotal time: %s\x1b[0m\n", time.Since(initialTime))
	}()

	loadPersonsUseCase := usecases.NewLoadPersonsUseCase()

	// Load persons from fs.
	persons, err := loadPersonsUseCase.Execute(personsDir)

	initRecognizerTime := time.Now()

	// Init the recognizer.
	rec, err := face.NewRecognizer(modelsDir)
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)
	}
	// Free the resources when you're finished.
	defer rec.Close()

	fmt.Println("Time to init recognizer: ", time.Since(initRecognizerTime))

	// Recognize faces from a few images and assign them to persons.
	recognizePersonsUseCase := usecases.NewRecognizePersonsUseCase(rec)

	err = recognizePersonsUseCase.Execute(persons)

	if err != nil {
		log.Fatalf("Can't recognize persons: %v", err)
	}

	classifyPersonUseCase := usecases.NewClassifyPersonsUseCase(rec)

	// Classify the unknown faces.
	unkImagePath := filepath.Join(imagesDir, "unknown.jpg")

	recognizedFaces, err := classifyPersonUseCase.Execute(unkImagePath, 0.3)

	if err != nil {
		log.Fatalf("Can't classify persons: %v", err)
	}

	// print the result
	fmt.Printf("\033[0;32mFound %d faces\033[0m\n", len(recognizedFaces))
	for _, recognizedFace := range recognizedFaces {
		fmt.Printf("\033[0;32mPerson: %s\033[0m\n", persons[recognizedFace.ID].Name)
	}

	drawer := entity.NewDrawer(unkImagePath, filepath.Join("fonts", "Roboto-Regular.ttf"))

	for _, recognizedFace := range recognizedFaces {
		drawer.DrawFace(recognizedFace.Face.Rectangle, persons[recognizedFace.ID].Name)
	}

	drawer.SaveImage(filepath.Join(imagesDir, "result.jpg"))

}
