package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Kagami/go-face"
	"github.com/andre-ols/go-face-recognition/internal/entity"
)

var (
	modelsDir = "models"
	imagesDir = "images"
)

// This example shows the basic usage of the package: create an
// recognizer, recognize faces, classify them using few known ones.
func main() {

	// Init the recognizer.
	rec, err := face.NewRecognizer(modelsDir)
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)
	}
	// Free the resources when you're finished.
	defer rec.Close()

	obama := entity.NewPerson(1, "Obama", []string{"obama.jpg", "obama-2.jpg"})
	biden := entity.NewPerson(2, "Biden", []string{"biden.jpg", "biden-2.jpg"})

	persons := map[int]*entity.Person{}

	persons[obama.ID] = obama
	persons[biden.ID] = biden

	var faces []face.Descriptor
	var categories []int32

	for _, p := range persons {
		for _, imagePath := range p.ImagesPath {
			face, err := rec.RecognizeSingleFile(filepath.Join(imagesDir, imagePath))
			if err != nil {
				log.Fatalf("Can't recognize: %v", err)
			}
			if face == nil {
				log.Fatalf("Not a single face on the image")
			}
			faces = append(faces, face.Descriptor)
			categories = append(categories, int32(p.ID))
		}
	}

	// Pass the samples to the recognizer.
	rec.SetSamples(faces, categories)

	// Now let's try to classify some not yet known image.
	testImage := filepath.Join(imagesDir, "unknown.jpg")

	unkFace, err := rec.RecognizeSingleFile(testImage)

	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	if unkFace == nil {
		log.Fatalf("Not a single face on the image")
	}

	// Find the closest person.
	catID := rec.ClassifyThreshold(unkFace.Descriptor, 0.3)

	if catID < 0 {
		log.Fatalf("Can't classify")
	}

	// print the result
	fmt.Println("Found person: ", persons[catID].Name)
}
