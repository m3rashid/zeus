package flow

import (
	"auth/config"
	"auth/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetUsersToSelect(ctx *fiber.Ctx) error {
	userIds := getLocalUserIDsFromCookie(ctx)
	users := []models.User{}

	if len(userIds) > 0 {
		db, err := config.GetDb()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(users)
		}

		if err := db.Where("id in (?)", userIds).Find(&users).Error; err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(users)
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}

type SelectUserRequestBody struct {
	UserID   uint   `json:"userId" validate:"required"`
	ClientID string `json:"clientId" validate:"required"`
}

func HandleSelectUser(ctx *fiber.Ctx) error {
	var selectUserBody SelectUserRequestBody
	if err := ctx.BodyParser(&selectUserBody); err != nil {
		return err
	}

	if err := validator.New().Struct(selectUserBody); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	db, err := config.GetDb()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	var client models.Client
	if err := db.Where("client_id = ?", selectUserBody.ClientID).First(&client).Error; err != nil || client.ID == 0 {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	var user models.User
	if err := db.Where("id = ?", selectUserBody.UserID).First(&user).Error; err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	connectApp := models.ConnectedClients{
		ClientID: client.ID,
		UserID:   user.ID,
	}

	if err := db.Create(&connectApp).Error; err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"authorization_code": "",
		//
	})
}
