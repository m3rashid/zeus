package utils

import "time"

type Endpoint struct {
	URL          string `json:"url"`
	Method       string `json:"method"`
	Description  string `json:"description"`
	AuthRequired bool   `json:"auth_required"`
}

type Endpoints []Endpoint

type ServerDetails struct {
	Name        string    `json:"name"`
	BaseUrl     string    `json:"url"`
	Endpoints   Endpoints `json:"endpoints"`
	LastUpdated time.Time `json:"last_updated"`
}

type ResourceServersMap = map[string]ServerDetails
