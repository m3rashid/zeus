package flow

import "github.com/gofiber/fiber/v2"

type GetSelectUserRequestBody struct {
}

func GetUsersToSelect(ctx *fiber.Ctx) error {
	return nil
}

type SelectUserRequestBody struct {
	UserID uint `json:"userId"`
}

func HandleSelectUser(ctx *fiber.Ctx) error {
	var body SelectUserRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	return nil
}
