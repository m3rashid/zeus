package main

import (
	"log"
	"notifications/config"
	"os"
	"utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Server = utils.Server{
		Name:       "notifications",
		ServerIP:   "localhost",
		Modules:    []utils.Module{},
		ServerPort: os.Getenv("PORT"),
		CorsConfig: cors.Config{},
	}

	config.Server.SetupStructuredLogger()

	app := fiber.New(config.Server.GetDefaultFiberConfig(fiber.Config{}))
	app.Use(cors.New(config.Server.CorsConfig))

	app.Get(utils.DISCOVERY_ENDPOINT, config.Server.DiscoveryHandler)

	log.Printf("Starting %s server ...\n", config.Server.Name)
	app.Listen(":" + config.Server.ServerPort)
}
