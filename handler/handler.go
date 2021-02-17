package handler

import (
	"errors"
	"log"

	"github.com/gofiber/fiber"
	"github.com/nadirbasalamah/go-simple-api/model"
	"github.com/nadirbasalamah/go-simple-api/service"
)

// GetProducts return all product data
func GetProducts(c *fiber.Ctx) {
	result, err := service.GetProducts()
	if err != nil {
		failedResponse(c, err, 500)
		return
	}

	if len(result) == 0 {
		failedResponse(c, errors.New("Product data not found"), 404)
		return
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"product": result,
		"message": "All products data found",
	}); err != nil {
		failedResponse(c, err, 500)
		return
	}
}

// GetProduct return product data by product id
func GetProduct(c *fiber.Ctx) {
	id := c.Params("id")
	product := model.Product{}

	product, err := service.GetProduct(id)
	if err != nil {
		failedResponse(c, err, 500)
		return
	}

	if product.Id == 0 {
		failedResponse(c, errors.New("Product data not found"), 404)
		return
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Product Data Found",
		"product": product,
	}); err != nil {
		failedResponse(c, err, 500)
		return
	}
}

// CreateProduct func to add new product
func CreateProduct(c *fiber.Ctx) {
	p := new(model.Product)
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		failedResponse(c, err, 400)
		return
	}

	err := p.ValidateRequest()

	if err != nil {
		failedResponse(c, err, 400)
		return
	}

	res, err := service.CreateProduct(*p)
	if err != nil {
		failedResponse(c, err, 500)
		return
	}

	log.Println(res)

	if err := c.Status(201).JSON(&fiber.Map{
		"success": true,
		"message": "Product created",
	}); err != nil {
		failedResponse(c, err, 500)
		return
	}
}

// EditProduct func to edit a product data
func EditProduct(c *fiber.Ctx) {
	id := c.Params("id")
	p := new(model.Product)
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		failedResponse(c, err, 400)
		return
	}

	err := p.ValidateRequest()

	if err != nil {
		failedResponse(c, err, 400)
		return
	}

	res, err := service.EditProduct(*p, id)
	if err != nil {
		failedResponse(c, err, 500)
		return
	}

	log.Println(res)

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Product updated",
		"product": p,
	}); err != nil {
		failedResponse(c, err, 500)
		return
	}
}

// DeleteProduct func to delete a product data
func DeleteProduct(c *fiber.Ctx) {
	id := c.Params("id")
	err := service.DeleteProduct(id)
	if err != nil {
		failedResponse(c, err, 500)
		return
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Product deleted",
	}); err != nil {
		failedResponse(c, err, 500)
		return
	}
}

func failedResponse(c *fiber.Ctx, err error, code int) {
	c.Status(code).JSON(&fiber.Map{
		"success": false,
		"message": err.Error(),
	})
}
