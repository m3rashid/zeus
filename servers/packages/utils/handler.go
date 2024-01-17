package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RequestContextData struct {
	RequestID string `json:"request_id"`
	UserID    uint   `json:"user_id"`
	HopCount  int    `json:"hop_count"`
}

func NewRequestID() string {
	return uuid.New().String()
}

func QueryParamsToUrl(baseUrl string, queryMap map[string]string) string {
	url := baseUrl + "?"
	for key, value := range queryMap {
		if value != "" && key != "" {
			url += key + "=" + value + "&"
		}
	}

	return strings.TrimSuffix(url, "&")
}

func GetQueryFomUrl[ReturnType interface{}](ctx *fiber.Ctx, returnType *ReturnType) {
	ctx.QueryParser(returnType)
}
