package flow

import "github.com/gofiber/fiber/v2"

type RegisterRequestBody struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func HandleRegister(ctx *fiber.Ctx) error {
	var body RegisterRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	return nil
}
