package flow

import (
	"auth/config"
	"auth/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleLogin(ctx *fiber.Ctx) error {
	var loginBody LoginRequestBody
	if err := ctx.BodyParser(&loginBody); err != nil {
		return err
	}

	if err := validator.New().Struct(loginBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	db, err := config.GetDb()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	var user models.User
	if err := db.Where("email = ?", loginBody.Email).First(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	if user.ID == 0 || !user.ComparePassword(loginBody.Password) {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	addUserIDToCookie(ctx, user.ID)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}
