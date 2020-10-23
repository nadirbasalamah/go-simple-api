package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	_ "github.com/lib/pq"
	"github.com/nadirbasalamah/go-simple-api/database"
	"github.com/nadirbasalamah/go-simple-api/router"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(middleware.Logger())

	router.SetupRoutes(app)

	app.Listen(3000)
}
