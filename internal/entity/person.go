package entity

type Person struct {
	ID         int
	Name       string
	ImagesPath []string
}

func NewPerson(id int, name string, imagesPath []string) *Person {
	return &Person{
		ID:         id,
		Name:       name,
		ImagesPath: imagesPath,
	}
}
