package flow

import (
	"common/server"

	"github.com/gofiber/fiber/v2"
)

var FlowModule = server.Module{
	Name:      "flow",
	SchemaMap: server.SchemaMap{},
	Models:    server.Models{},
	ProtectedRoutes: server.ProtectedRoutes{
		"/logout": {
			HttpMethod:  "GET",
			Controller:  ServeLogoutPage,
			Description: "Serve the logout page",
		},
		"/logout/handle": {
			HttpMethod:  "POST",
			Controller:  HandleLogout,
			Description: "Handle the logout request",
		},
	},
	AnonymousRoutes: server.AnonymousRoutes{
		"/register": {
			HttpMethod:  "GET",
			Controller:  ServeRegisterPage,
			Description: "Serve the register page",
		},
		"/register/handle": {
			HttpMethod:  "POST",
			Controller:  HandleRegister,
			Description: "Handle the register request",
		},
		"/login": {
			HttpMethod:  "GET",
			Controller:  ServeLoginPage,
			Description: "Serve the login page",
		},
		"/login/handle": {
			HttpMethod:  "POST",
			Controller:  HandleLogin,
			Description: "Handle the login request",
		},
	},
}

func Setup(app *fiber.App) {}
