package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	Name       string      `json:"name"`
	IP         string      `json:"ip"`
	Port       string      `json:"port"`
	CorsConfig cors.Config `json:"cors_config"`
}

func (s *Server) GetDefaultFiberConfig(overrides fiber.Config) fiber.Config {
	return fiber.Config{
		CaseSensitive:         true,
		PassLocalsToViews:     true,
		DisableStartupMessage: true,
		Concurrency:           256 * 1024 * 1024,
		AppName:               s.Name,
		ServerHeader:          s.Name,
		RequestMethods:        []string{"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			log.Println(err.Error())
			return ctx.Status(code).JSON(fiber.Map{"success": false})
		},
	}
}
