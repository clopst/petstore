package model

type Breed struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	CategoryID int      `json:"category_id"`
	Category   Category `json:"category"`
}

type BreedService interface {
	GetAllBreed() ([]*Breed, error)
	ShowBreed(id int) (*Breed, error)
	CreateBreed() error
}
