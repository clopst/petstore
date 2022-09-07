package model

type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LocationService interface {
	GetAllLocation() ([]*Location, error)
	ShowLocation(id int) (*Location, error)
	CreateLocation() error
}
