package discovery

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"utils"

	"github.com/gofiber/fiber/v2"
)

var DiscoveryEndpoints = map[string]string{
	"notifications": "http://localhost:5001" + utils.DISCOVERY_ENDPOINT,
	// ...
}

var ResourceServerDetailsMap = utils.ResourceServersMap{}

func DiscoverServers() {
	timeTick := time.NewTicker(time.Second * 5)

	for range timeTick.C {
		for name, url := range DiscoveryEndpoints {
			log.Printf("Discovering %s server at %s", name, url)
			response, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error in getting response from %s server\n", name)
				break
			}

			byteArr, err := io.ReadAll(response.Body)
			if err != nil {
				break
			}

			var resourceServerDetails utils.ServerDetails
			if err := json.Unmarshal(byteArr, &resourceServerDetails); err != nil {
				fmt.Printf("Error in parsing response from %s server\n", name)
				break
			}

			resourceServerDetails.LastUpdated = time.Now()
			ResourceServerDetailsMap[name] = resourceServerDetails
		}
	}
}

func DiscoverResourceServers(ctx *fiber.Ctx) error {
	scope := ctx.Query("scope", "all")
	if scope == "all" {
		return ctx.Status(fiber.StatusOK).JSON(ResourceServerDetailsMap)
	}

	details, resourceServerExists := ResourceServerDetailsMap[scope]
	if !resourceServerExists {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No such resource server exists",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		scope: details,
	})
}
