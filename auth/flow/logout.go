package flow

import (
	"auth/config"
	"auth/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LogoutRequestBody struct {
	UserID uint `json:"userId" validate:"required"`
}

func HandleLogout(ctx *fiber.Ctx) error {
	var logoutBody LogoutRequestBody
	if err := ctx.BodyParser(&logoutBody); err != nil {
		return err
	}

	if err := validator.New().Struct(logoutBody); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	db, err := config.GetDb()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	var user models.User
	if err := db.Where("id = ?", logoutBody.UserID).Error; err != nil || user.ID == 0 {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	removeUserIDFromCookie(ctx, user.ID)
	return ctx.SendStatus(fiber.StatusOK)
}
