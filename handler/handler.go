package handler

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/nadirbasalamah/go-simple-api/model"
	"github.com/nadirbasalamah/go-simple-api/service"
)

// GetProducts return all product data
func GetProducts(c *fiber.Ctx) {
	result, err := service.GetProducts()
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}

	if len(result) == 0 {
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

	product, err := service.GetProduct(id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
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

	res, err := service.CreateProduct(*p)
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

	res, err := service.EditProduct(*p, id)
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
	err := service.DeleteProduct(id)
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
