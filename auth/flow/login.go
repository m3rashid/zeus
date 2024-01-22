package flow

import "github.com/gofiber/fiber/v2"

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleLogin(ctx *fiber.Ctx) error {
	var body LoginRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	return nil
}
