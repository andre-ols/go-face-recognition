package usecases

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/andre-ols/go-face-recognition/internal/entity"
)

type LoadPersonsUseCase interface {
	Execute(personsDir string) (map[int]*entity.Person, error)
}

type LoadPersonsUseCaseImpl struct {
}

// This function `Execute` is the implementation of the `LoadPersonsUseCase` interface.
// Execute reads people's images from a directory of folders.
// The root directory must be provided as input. Each person should have a subfolder within the root directory
// with the person's name being the folder name. Inside the folder should contain photos of the person.
// This function will dynamically load the persons and relate them to their photos.
func (l *LoadPersonsUseCaseImpl) Execute(personsDir string) (map[int]*entity.Person, error) {

	loadPersonsTime := time.Now()

	defer func() {
		fmt.Println("Time to load persons: ", time.Since(loadPersonsTime))
	}()

	// Open persons directory
	dir, err := os.Open(personsDir)
	if err != nil {
		fmt.Println("Error on open persons directory:", err)
	}
	defer dir.Close()

	// Read the directory content
	subFolders, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error on read persons directory:", err)
		return nil, err
	}

	persons := map[int]*entity.Person{}

	for index, subFolder := range subFolders {
		if !subFolder.IsDir() {
			continue
		}

		personName := subFolder.Name()

		personDir := filepath.Join(personsDir, personName)

		files, err := os.ReadDir(personDir)
		if err != nil {
			fmt.Printf("Error on read person directory: %s\n", personDir)
			return nil, err
		}

		var imagesPath []string
		for _, file := range files {
			imagesPath = append(imagesPath, filepath.Join(personDir, file.Name()))
		}

		persons[index] = entity.NewPerson(index, personName, imagesPath)

	}

	return persons, nil
}

func NewLoadPersonsUseCase() LoadPersonsUseCase {
	return &LoadPersonsUseCaseImpl{}
}
