package flow

import (
	"auth/models"
	"common/server"
)

var FlowModule = server.Module{
	Name:      "flow",
	SchemaMap: server.SchemaMap{},
	Models: server.Models{
		models.Acl{},
		models.Group{},
		models.Tenant{},
		models.User{},
		models.Client{},
	},
	ProtectedRoutes: server.ProtectedRoutes{
		"logout": {
			HttpMethod:  "POST",
			Controller:  HandleLogout,
			Description: "Handle the logout request",
		},
		"select-user": {
			HttpMethod:  "POST",
			Controller:  HandleSelectUser,
			Description: "Select user to login as",
		},
		"users": {
			HttpMethod:  "GET",
			Controller:  GetUsersToSelect,
			Description: "Get list of users to select from",
		},
	},
	AnonymousRoutes: server.AnonymousRoutes{
		"register": {
			HttpMethod:  "POST",
			Controller:  HandleRegister,
			Description: "Handle the register request",
		},
		"login": {
			HttpMethod:  "POST",
			Controller:  HandleLogin,
			Description: "Handle the login request",
		},
	},
}
