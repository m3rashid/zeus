package utils

import (
	"models"

	"github.com/gofiber/fiber/v2"
)

type ProtectedRoutes = map[string]struct {
	HttpMethod        string // default POST
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
