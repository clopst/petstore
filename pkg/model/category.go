package model

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryService interface {
	GetAllCategory() ([]*Category, error)
	ShowCategory(id int) (*Category, error)
	CreateCategory() error
}
