package usecases

import (
	"fmt"
	"time"

	"github.com/Kagami/go-face"
	"github.com/andre-ols/go-face-recognition/internal/entity"
)

type RecognizePersonsUseCase interface {
	Execute(persons map[int]*entity.Person) error
}

type RecognizePersonsUseCaseImpl struct {
	rec *face.Recognizer
}

// This function `Execute` in the `RecognizePersonsUseCaseImpl` struct is responsible for recognizing
// faces of persons based on the images provided in the `persons` map.
// Execute performs recognition of known faces for the provided persons.
// Each photo should contain only one face, which is the face of the person.
// It is ideal for each person to have more than one photo to improve recognition.
func (r *RecognizePersonsUseCaseImpl) Execute(persons map[int]*entity.Person) error {

	knowFacesTime := time.Now()

	defer func() {
		fmt.Println("Time to recognize known faces: ", time.Since(knowFacesTime))
	}()

	faces := []face.Descriptor{}
	categories := []int32{}

	for _, p := range persons {
		for _, imagePath := range p.ImagesPath {
			face, err := r.rec.RecognizeSingleFile(imagePath)
			if err != nil {
				return fmt.Errorf("Can't recognize: %v", err)
			}
			if face == nil {
				return fmt.Errorf("Unable to recognize people in the image")
			}
			faces = append(faces, face.Descriptor)
			categories = append(categories, int32(p.ID))
		}
	}

	// Pass the samples to the recognizer.
	r.rec.SetSamples(faces, categories)

	return nil
}

func NewRecognizePersonsUseCase(rec *face.Recognizer) RecognizePersonsUseCase {
	return &RecognizePersonsUseCaseImpl{rec: rec}
}
