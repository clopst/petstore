package postgres

import (
	"database/sql"
	"errors"

	"github.com/clopst/petstore/internal/database"
	"github.com/clopst/petstore/pkg/model"
)

type Category model.Category
type CategoryService model.CategoryService

func (c *Category) CreateCategory() (*model.Category, error) {
	row := database.DBClient.QueryRow(
		"INSERT INTO categories (name) VALUES ($1) RETURNING *;",
		c.Name,
	)

	err := row.Scan(&c.ID, &c.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("insert failed")
		}

		return nil, err
	}

	category := model.Category(*c)
	return &category, nil
}

func (c Category) ShowCategory(id int) (*model.Category, error) {
	row := database.DBClient.QueryRow("SELECT * FROM categories WHERE id=$1", id)

	err := row.Scan(&c.ID, &c.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("empty result")
		}

		return nil, err
	}

	category := model.Category(c)
	return &category, nil
}

func (c Category) GetAllCategory() ([]*model.Category, error) {
	var categories []*model.Category

	rows, err := database.DBClient.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}

		category := model.Category(c)
		categories = append(categories, &category)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
