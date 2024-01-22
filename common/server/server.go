package server

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/utils"
)

type Server struct {
	Name       string      `json:"name"`
	Modules    []Module    `json:"modules"`
	ServerIP   string      `json:"network_ip"`
	ServerPort string      `json:"network_port"`
	CorsConfig cors.Config `json:"cors_config"`
	DBUri      string      `json:"db_uri"`
	RedisUri   string      `json:"redis_uri"`
	Logger     *slog.Logger
}

func (s *Server) SetupStructuredLogger() {
	// TODO: handle logging to a file
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	s.Logger = logger
}

func (s *Server) SetupDefaultCaches() cache.Config {
	return cache.Config{
		Expiration:   10 * time.Minute,
		CacheControl: true,
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.Path())
		},
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
	}
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

const DISCOVERY_ENDPOINT = "/api/discovery"

func (s *Server) DiscoveryHandler(ctx *fiber.Ctx) error {
	endpoints := Endpoints{}
	for _, module := range s.Modules {
		endpoints = append(endpoints, module.GetEndpoints()...)
	}

	return ctx.Status(fiber.StatusOK).JSON(ServerDetails{
		Name:        s.Name,
		Endpoints:   endpoints,
		LastUpdated: time.Now(),
		BaseUrl:     "http://" + s.ServerIP + ":" + s.ServerPort,
	})
}
