package flow

import (
	"auth/config"
	"auth/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LoginRequestBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func HandleLogin(ctx *fiber.Ctx) error {
	var loginBody LoginRequestBody
	if err := ctx.BodyParser(&loginBody); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := validator.New().Struct(loginBody); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	db, err := config.GetDb()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	var user models.User
	if err := db.Where("email = ?", loginBody.Email).First(&user).Error; err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	if user.ID == 0 || !user.ComparePassword(loginBody.Password) {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	addUserIDToCookie(ctx, user.ID)
	return ctx.SendStatus(fiber.StatusOK)
}
