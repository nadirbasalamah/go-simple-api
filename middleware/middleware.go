package middleware

import (
	"github.com/gofiber/basicauth"
	"github.com/gofiber/fiber"
	"github.com/nadirbasalamah/go-simple-api/config"
)

// AuthReq func as a middleware using basic authentication
func AuthReq() func(*fiber.Ctx) {
	cfg := basicauth.Config{
		Users: map[string]string{
			config.Config("USERNAME"): config.Config("PASSWORD"),
		},
	}
	err := basicauth.New(cfg)
	return err
}
