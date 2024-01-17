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

type Module struct {
	Name            string
	SchemaMap       SchemaMap
	Models          []interface{}
	AnonymousRoutes AnonymousRoutes
	ProtectedRoutes ProtectedRoutes
}

func (module *Module) HeartBeat() ModuleHeartBeat {
	var endpointsMap = make(EndpointsMap)
	for route, routeDetails := range module.AnonymousRoutes {
		endpointsMap[strings.Join([]string{"/api/anonymous", module.Name, route}, "/")] = Endpoint{
			Method:      common.Ternary[HTTPMethod](routeDetails.HttpMethod == "", DEFAULT_HTTP_METHOD, routeDetails.HttpMethod),
			Description: common.Ternary[string](routeDetails.Description == "", "No description provided", routeDetails.Description),
		}
	}

	for route, routeDetails := range module.ProtectedRoutes {
		endpointsMap[strings.Join([]string{"/api", module.Name, route}, "/")] = Endpoint{
			Method:      common.Ternary[HTTPMethod](routeDetails.HttpMethod == "", DEFAULT_HTTP_METHOD, routeDetails.HttpMethod),
			Description: common.Ternary[string](routeDetails.Description == "", "No description provided", routeDetails.Description),
		}
	}

	return ModuleHeartBeat{
		ModuleName:   module.Name,
		EndpointsMap: endpointsMap,
	}
}
