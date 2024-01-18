package utils

import (
	"common"
	"models"
	"strings"

	"github.com/gofiber/fiber/v2"
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
	AclRelationNeeded models.AclRelation
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
			URL:          strings.Join([]string{"/api/anonymous", module.Name, route}, "/"),
			Method:       common.Ternary[HTTPMethod](routeDetails.HttpMethod == "", DEFAULT_HTTP_METHOD, routeDetails.HttpMethod),
			Description:  common.Ternary[string](routeDetails.Description == "", "No description provided", routeDetails.Description),
		})
	}

	for route, routeDetails := range module.ProtectedRoutes {
		endpointsMap = append(endpointsMap, Endpoint{
			AuthRequired: true,
			URL:          strings.Join([]string{"/api", module.Name, route}, "/"),
			Method:       common.Ternary[HTTPMethod](routeDetails.HttpMethod == "", DEFAULT_HTTP_METHOD, routeDetails.HttpMethod),
			Description:  common.Ternary[string](routeDetails.Description == "", "No description provided", routeDetails.Description),
		})
	}

	return endpointsMap
}
