package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Endpoint struct {
	Path         string                     `json:"path"`
	Method       string                     `json:"method"`
	Description  string                     `json:"description"`
	AuthRequired bool                       `json:"auth_required"`
	Controller   func(ctx *fiber.Ctx) error `json:"-"`
}

type Endpoints []Endpoint

type ServerDetails struct {
	Name        string    `json:"name"`
	BaseUrl     string    `json:"url"`
	Endpoints   Endpoints `json:"endpoints"`
	LastUpdated time.Time `json:"last_updated"`
}

type ResourceServersMap = map[string]ServerDetails
