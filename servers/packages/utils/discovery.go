package utils

import "time"

type Endpoint struct {
	Method      string `json:"method"`
	Description string `json:"description"`
}

type EndpointsMap map[string]Endpoint

type AllowedScopes map[string][]string

type ResourceServerDetails struct {
	BaseUrl       string        `json:"url"`
	Endpoints     EndpointsMap  `json:"endpoints"`
	LastUpdated   time.Time     `json:"last_updated"`
	AllowedScopes AllowedScopes `json:"scopes"`
}

type ResourceServersMap = map[string]ResourceServerDetails

type IncomingScopeRequest = string

func (scopes *AllowedScopes) ToString() (string, error) {

	return "", nil
}

type ModuleHeartBeat struct {
	ModuleName   string       `json:"module_name"`
	EndpointsMap EndpointsMap `json:"endpoints_map"`
}
