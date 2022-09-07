package model

type Pet struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Age         float64  `json:"age"`
	Description string   `json:"description"`
	ImageURL    string   `json:"image_url"`
	LocationID  int      `json:"location_id"`
	BreedID     int      `json:"breed_id"`
	Location    Location `json:"location"`
	Breed       Breed    `json:"breed"`
}

type PetService interface {
	GetAllPet() ([]*Pet, error)
	GetPetByCategory(categoryID int) ([]*Pet, error)
	ShowPet(id int) (*Pet, error)
	CreatePet() error
	DeletePet(id int) error
}
