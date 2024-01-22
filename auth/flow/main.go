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
			HttpMethod:  "POST",
			Controller:  HandleLogout,
			Description: "Handle the logout request",
		},
		"/select-user": {
			HttpMethod:  "POST",
			Controller:  HandleSelectUser,
			Description: "Select user to login as",
		},
		"/get-select-users": {
			HttpMethod:  "POST",
			Controller:  GetUsersToSelect,
			Description: "Get list of users to select from",
		},
	},
	AnonymousRoutes: server.AnonymousRoutes{
		"/register": {
			HttpMethod:  "POST",
			Controller:  HandleRegister,
			Description: "Handle the register request",
		},
		"/login": {
			HttpMethod:  "POST",
			Controller:  HandleLogin,
			Description: "Handle the login request",
		},
	},
}

func Setup(app *fiber.App) {}
