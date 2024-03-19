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

	defer func() {
		fmt.Printf("\033[0;33mTotal time: %s\033[0;33m", time.Since(initialTime))
	}()

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

	faces := []face.Descriptor{}
	categories := []int32{}

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

	fmt.Println("Time to recognize known faces: ", time.Since(knowFacesTime))

	unknownFacesTime := time.Now()

	// Now let's try to classify some not yet known image.
	unkImage := filepath.Join(imagesDir, "unknown.jpg")

	unkFaces, err := rec.RecognizeFile(unkImage)

	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	if len(unkFaces) == 0 {
		log.Fatalf("Don't have faces on the image")
	}

	fmt.Println("Time to recognize unknown faces: ", time.Since(unknownFacesTime))

	classifyTime := time.Now()

	// Classify the unknown faces
	knowFaces := []face.Face{}
	knowFacesID := []int{}
	for _, unkFace := range unkFaces {
		catID := rec.ClassifyThreshold(unkFace.Descriptor, 0.3)

		if catID < 0 {
			continue
		}
		knowFaces = append(knowFaces, unkFace)
		knowFacesID = append(knowFacesID, int(catID))
	}

	fmt.Println("Time to classify unknown faces: ", time.Since(classifyTime))

	// print the result
	fmt.Printf("\033[0;32mFound %d faces\033[0m\n", len(knowFaces))
	for _, knowFaceID := range knowFacesID {
		fmt.Printf("\033[0;32mPerson: %s\033[0m\n", persons[int(knowFaceID)].Name)
	}

	drawer := entity.NewDrawer(unkImage, filepath.Join("fonts", "Roboto-Regular.ttf"))

	for index, knowFace := range knowFaces {
		drawer.DrawFace(knowFace.Rectangle, persons[knowFacesID[index]].Name)
	}

	drawer.SaveImage(filepath.Join(imagesDir, "result.jpg"))

}
