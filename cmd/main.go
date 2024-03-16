package main

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

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

	initialTime := time.Now()

	// Init the recognizer.
	rec, err := face.NewRecognizer(modelsDir)
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)
	}
	// Free the resources when you're finished.
	defer rec.Close()

	fmt.Println("Time to init recognizer: ", time.Since(initialTime))

	knowFacesTime := time.Now()

	trump := entity.NewPerson(1, "Trump", []string{"trump.jpg", "trump-2.jpg"})
	biden := entity.NewPerson(2, "Biden", []string{"biden.jpg", "biden-2.jpg"})

	persons := map[int]*entity.Person{}

	persons[trump.ID] = trump
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

	fmt.Println("Time to recognize know faces: ", time.Since(knowFacesTime))

	unknownFacesTime := time.Now()

	// Pass the samples to the recognizer.
	rec.SetSamples(faces, categories)

	// Now let's try to classify some not yet known image.
	testImage := filepath.Join(imagesDir, "unknown.jpg")

	unkFaces, err := rec.RecognizeFile(testImage)

	fmt.Println("Faces: ", len(unkFaces))

	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	if len(unkFaces) == 0 {
		log.Fatalf("Don't have faces on the image")
	}

	fmt.Println("Time to recognize unknown faces: ", time.Since(unknownFacesTime))

	classifyTime := time.Now()

	// Classify the unkown faces
	knowFacesID := []int{}
	for _, unkFace := range unkFaces {
		catID := rec.ClassifyThreshold(unkFace.Descriptor, 0.3)

		if catID < 0 {
			continue
		}
		knowFacesID = append(knowFacesID, catID)
	}

	fmt.Println("Time to classify unknown faces: ", time.Since(classifyTime))

	fmt.Println("Total time: ", time.Since(initialTime))

	// print the result
	fmt.Println("knowFaces")
	for _, knowFaceID := range knowFacesID {
		fmt.Println(persons[int(knowFaceID)].Name)
	}
}
