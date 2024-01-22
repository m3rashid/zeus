package main

import (
	"common/server"
	"log"
	"notifications/config"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Server = server.Server{
		Name:       "notifications",
		ServerIP:   "localhost",
		Modules:    []server.Module{},
		ServerPort: os.Getenv("PORT"),
		CorsConfig: cors.Config{},
		DBUri:      os.Getenv("DB_URI"),
		RedisUri:   os.Getenv("REDIS_URI"),
	}

	config.Server.SetupStructuredLogger()

	app := fiber.New(config.Server.GetDefaultFiberConfig(fiber.Config{}))
	app.Use(cors.New(config.Server.CorsConfig))

	app.Get(server.DISCOVERY_ENDPOINT, config.Server.DiscoveryHandler)

	log.Printf("Starting %s server ...\n", config.Server.Name)
	app.Listen(":" + config.Server.ServerPort)
}
