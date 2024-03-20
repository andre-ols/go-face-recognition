package usecases

import (
	"fmt"
	"time"

	"github.com/Kagami/go-face"
)

type ClassifyPersonsUseCase interface {
	Execute(unkImagePath string, threshold float32) ([]*RecognizedFaces, error)
}

type ClassifyPersonsUseCaseImpl struct {
	rec *face.Recognizer
}

type RecognizedFaces struct {
	ID   int
	Face *face.Face
}

// This function `Execute` is a method of the `ClassifyPersonsUseCaseImpl` struct that implements the
// `ClassifyPersonsUseCase` interface.
// Execute performs the recognition of unknown faces against known faces and classifies them.
// The unknown image path and a threshold for classification confidence must be provided as input.
// This function recognizes faces in the unknown image and classifies them against known faces.
// It returns a slice of RecognizedFaces, each containing the ID of the recognized person and the recognized face,
// and an error if the recognition process fails.
// It is important to note that each photo should contain only one face, which is the face of the person.
// Also, it's ideal for each person to have more than one photo to improve recognition accuracy.
func (c *ClassifyPersonsUseCaseImpl) Execute(unkImagePath string, threshold float32) ([]*RecognizedFaces, error) {
	unknownFacesTime := time.Now()

	defer func() {
		fmt.Println("Time to recognize unknown faces: ", time.Since(unknownFacesTime))
	}()

	unkFaces, err := c.rec.RecognizeFile(unkImagePath)

	if err != nil {
		return nil, fmt.Errorf("Can't recognize: %v", err)
	}

	if len(unkFaces) == 0 {
		return nil, fmt.Errorf("Unable to recognize people in the image")
	}

	// Classify the unknown faces
	knowFaces := []*RecognizedFaces{}
	for _, unkFace := range unkFaces {
		catID := c.rec.ClassifyThreshold(unkFace.Descriptor, threshold)

		if catID < 0 {
			continue
		}
		knowFaces = append(knowFaces, &RecognizedFaces{ID: int(catID), Face: &unkFace})
	}

	return knowFaces, nil
}

func NewClassifyPersonsUseCase(rec *face.Recognizer) ClassifyPersonsUseCase {
	return &ClassifyPersonsUseCaseImpl{rec: rec}
}
