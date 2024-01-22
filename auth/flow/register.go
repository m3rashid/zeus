package flow

import (
	"auth/config"
	"auth/models"
	"common/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string, hash *string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	*hash = string(bytes)
	return nil
}

type RegisterRequestBody struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}

func HandleRegister(ctx *fiber.Ctx) error {
	var registerBody RegisterRequestBody
	if err := ctx.BodyParser(&registerBody); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if err := validator.New().Struct(registerBody); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	if registerBody.Password != registerBody.ConfirmPassword {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	var hashedPassword string
	if err := hashPassword(registerBody.Password, &hashedPassword); err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	db, err := config.GetDb()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	var user models.User
	if err := db.Where("email = ?", registerBody.Email).First(&user).Error; err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	if user.ID != 0 {
		return ctx.SendStatus(fiber.StatusConflict)
	}

	newUser := models.User{
		Name:     registerBody.Name,
		Email:    registerBody.Email,
		Password: hashedPassword,
		OTP:      utils.GenerateOtp(),
	}

	if err := db.Create(&newUser).Error; err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	addUserIDToCookie(ctx, newUser.ID)
	return ctx.SendStatus(fiber.StatusOK)
}
