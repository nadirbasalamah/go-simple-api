package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber"
	"github.com/nadirbasalamah/go-simple-api/database"
	"github.com/nadirbasalamah/go-simple-api/model"
)

// GetProducts return all product data
func GetProducts(c *fiber.Ctx) {
	rows, err := database.DB.Query("SELECT id, name, description, category, amount FROM products ORDER BY name")
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	defer rows.Close()
	result := model.Products{}
	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Category, &product.Amount)
		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
			return
		}
		result.Products = append(result.Products, product)
	}

	if len(result.Products) == 0 {
		c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "Product data not found",
		})
		return
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"product": result,
		"message": "All products data found",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
}

// GetProduct return product data by product id
func GetProduct(c *fiber.Ctx) {
	id := c.Params("id")
	product := model.Product{}

	row, err := database.DB.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	defer row.Close()

	for row.Next() {
		switch err := row.Scan(&product.Id, &product.Amount, &product.Name, &product.Description, &product.Category); err {
		case sql.ErrNoRows:
			log.Println("Data not found")
			c.Status(404).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		case nil:
			log.Println(product.Name, product.Description, product.Category, product.Amount)
		default:
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}
	}

	if product.Id == 0 {
		c.Status(404).JSON(&fiber.Map{
			"success": false,
			"message": "Product data not found",
		})
		return
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Product Data Found",
		"product": product,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
}

// CreateProduct func to add new product
func CreateProduct(c *fiber.Ctx) {
	p := new(model.Product)
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	err := p.ValidateRequest()

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	res, err := database.DB.Query("INSERT INTO products (name, description, category, amount) VALUES ($1, $2, $3, $4) ", p.Name, p.Description, p.Category, p.Amount)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	log.Println(res)

	if err := c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Product created",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Product failed to create",
		})
		return
	}
}

// EditProduct func to edit a product data
func EditProduct(c *fiber.Ctx) {
	id := c.Params("id")
	p := new(model.Product)
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	err := p.ValidateRequest()

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	res, err := database.DB.Query("UPDATE products SET name=$1, description=$2, category=$3, amount=$4 WHERE id=$5", p.Name, p.Description, p.Category, p.Amount, id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	log.Println(res)

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Product updated",
		"product": p,
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Product failed to update",
		})
		return
	}
}

// DeleteProduct func to delete a product data
func DeleteProduct(c *fiber.Ctx) {
	id := c.Params("id")
	_, err := database.DB.Query("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Product deleted",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
}
