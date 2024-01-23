package server

import (
	"common/utils"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HTTPMethod = string

const (
	HTTP_GET     HTTPMethod = "GET"
	HTTP_POST    HTTPMethod = "POST"
	HTTP_PUT     HTTPMethod = "PUT"
	HTTP_DELETE  HTTPMethod = "DELETE"
	HTTP_PATCH   HTTPMethod = "PATCH"
	HTTP_OPTIONS HTTPMethod = "OPTIONS"
	HTTP_HEAD    HTTPMethod = "HEAD"
)

const DEFAULT_HTTP_METHOD HTTPMethod = HTTP_GET

type ProtectedRoutes = map[string]struct {
	HttpMethod        HTTPMethod
	Controller        fiber.Handler
	Description       string
	AclRelationNeeded int
}

type AnonymousRoutes = map[string]struct {
	HttpMethod  string
	Controller  fiber.Handler
	Description string
}

type Schema = map[string]string
type SchemaMap = map[string]Schema
type Models = []interface{}

type Module struct {
	Name            string
	SchemaMap       SchemaMap
	Models          Models
	AnonymousRoutes AnonymousRoutes
	ProtectedRoutes ProtectedRoutes
}

type Modules = []Module

func (module *Module) GetEndpoints() Endpoints {
	var endpointsMap = Endpoints{}

	for route, routeDetails := range module.AnonymousRoutes {
		endpointsMap = append(endpointsMap, Endpoint{
			AuthRequired: false,
			Path:         strings.Join([]string{"/api/anonymous", module.Name, route}, "/"),
			Method:       utils.Ternary[HTTPMethod](routeDetails.HttpMethod == "", DEFAULT_HTTP_METHOD, routeDetails.HttpMethod),
			Description:  utils.Ternary[string](routeDetails.Description == "", "No description provided", routeDetails.Description),
			Controller:   routeDetails.Controller,
		})
	}

	for route, routeDetails := range module.ProtectedRoutes {
		endpointsMap = append(endpointsMap, Endpoint{
			AuthRequired: true,
			Path:         strings.Join([]string{"/api", module.Name, route}, "/"),
			Method:       utils.Ternary[HTTPMethod](routeDetails.HttpMethod == "", DEFAULT_HTTP_METHOD, routeDetails.HttpMethod),
			Description:  utils.Ternary[string](routeDetails.Description == "", "No description provided", routeDetails.Description),
			Controller:   routeDetails.Controller,
		})
	}

	return endpointsMap
}

func (server *Server) SetupEndpoints(app *fiber.App) {
	for _, module := range server.Modules {
		endpoints := module.GetEndpoints()
		for _, endpoint := range endpoints {
			app.Add(endpoint.Method, endpoint.Path, endpoint.Controller)
		}
	}
}

func (server *Server) MigrateDbModels(db *gorm.DB) {
	models := []interface{}{}
	for _, module := range server.Modules {
		models = append(models, module.Models...)
	}
	log.Println("Migrating Models for ", server.Name)
	db.AutoMigrate(models...)
}
