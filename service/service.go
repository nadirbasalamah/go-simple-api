package service

import (
	"database/sql"
	"errors"

	"github.com/nadirbasalamah/go-simple-api/database"
	"github.com/nadirbasalamah/go-simple-api/model"
)

// GetProducts returns all product data
func GetProducts() ([]model.Product, error) {
	rows, err := database.DB.Query("SELECT id, name, description, category, amount FROM products ORDER BY name")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	result := model.Products{}
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Category, &product.Amount)
		if err != nil {
			return nil, err
		}
		result.Products = append(result.Products, product)
	}

	return result.Products, nil
}

// GetProduct returns product data
func GetProduct(id string) (model.Product, error) {
	product := model.Product{}

	row, err := database.DB.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return model.Product{}, err
	}
	defer row.Close()

	for row.Next() {
		switch err := row.Scan(&product.Id, &product.Amount, &product.Name, &product.Description, &product.Category); err {
		case sql.ErrNoRows:
			return model.Product{}, errors.New("Data not found")
		case nil:
			return product, nil
		default:
			return model.Product{}, err
		}
	}

	return product, nil

}

// CreateProduct returns inserted product data
func CreateProduct(product model.Product) (model.Product, error) {
	_, err := database.DB.Query("INSERT INTO products (name, description, category, amount) VALUES ($1, $2, $3, $4) ", product.Name, product.Description, product.Category, product.Amount)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

// EditProduct returns edited product data
func EditProduct(product model.Product, id string) (model.Product, error) {
	_, err := database.DB.Query("UPDATE products SET name=$1, description=$2, category=$3, amount=$4 WHERE id=$5", product.Name, product.Description, product.Category, product.Amount, id)
	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

// DeleteProduct returns deleted product data
func DeleteProduct(id string) error {
	_, err := database.DB.Query("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
