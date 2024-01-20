package main

import (
	"auth/config"
	"auth/discovery"
	"auth/flow"
	"auth/templates"
	"log"
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
		Name:       "auth",
		ServerIP:   "localhost",
		Modules:    []utils.Module{},
		ServerPort: os.Getenv("PORT"),
		CorsConfig: cors.Config{},
	}

	go discovery.DiscoverServers()

	config.Server.SetupStructuredLogger()

	app := fiber.New(config.Server.GetDefaultFiberConfig(fiber.Config{}))
	app.Use(cors.New(config.Server.CorsConfig))

	app.Get(utils.DISCOVERY_ENDPOINT, config.Server.DiscoveryHandler)
	app.Get("/api/discover", discovery.DiscoverResourceServers)
	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Set("Content-Type", "text/html")
		return templates.LoginOrRegister(templates.LoginOrRegisterProps{
			IsLogin:     true,
			LoginUrl:    "/login",
			RegisterUrl: "/register",
		}).Render(ctx.Context(), ctx.Response().BodyWriter())
	})

	flow.Setup(app)

	log.Printf("Starting %s server ...\n", config.Server.Name)
	app.Listen(":" + config.Server.ServerPort)
}
