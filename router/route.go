package router

import (
	"github.com/gofiber/fiber"
	"github.com/nadirbasalamah/go-simple-api/handler"
	// "github.com/nadirbasalamah/go-simple-api/middleware"
)

// SetupRoutes func to define some routes
func SetupRoutes(app *fiber.App) {
	// Create a simple middleware
	// api := app.Group("/api/", middleware.AuthReq())

	// Create a routes
	app.Get("/products", handler.GetProducts)
	app.Get("/product/:id", handler.GetProduct)
	app.Post("/product", handler.CreateProduct)
	app.Put("/product/:id", handler.EditProduct)
	app.Delete("/product/:id", handler.DeleteProduct)
}
