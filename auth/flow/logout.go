package flow

import "github.com/gofiber/fiber/v2"

type LogoutRequestBody struct{}

func HandleLogout(ctx *fiber.Ctx) error {
	var body LogoutRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	return nil
}
