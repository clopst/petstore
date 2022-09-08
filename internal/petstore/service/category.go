package service

import (
	db "github.com/clopst/petstore/internal/petstore/repo/postgres"
	"github.com/clopst/petstore/pkg/model"
)

var (
	category        db.Category
	categoryService db.CategoryService
)

func GetAllCategory() ([]*model.Category, error) {
	categoryService = &category

	categories, err := categoryService.GetAllCategory()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func CreateCategory(c *db.Category) (int, error) {
	categoryService = c

	id, err := categoryService.CreateCategory()
	if err != nil {
		return id, err
	}

	return id, nil
}

func ShowCategory(id int) (*model.Category, error) {
	categoryService = &category

	category, err := categoryService.ShowCategory(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}