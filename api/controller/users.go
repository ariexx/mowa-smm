package controller

import (
	"github.com/gofiber/fiber/v2"
	"mowa-backend/api/services"
	"mowa-backend/api/utils"
)

type UserController interface {
	GetUser(ctx *fiber.Ctx, id uint32) error
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (u *userController) GetUser(ctx *fiber.Ctx, id uint32) error {
	userId := uint32(1)
	user, err := u.service.GetUser(ctx.Context(), userId)
	if err != nil {
		return err
	}

	return ctx.JSON(utils.ApiResponseSuccess("User fetched successfully", user))
}
