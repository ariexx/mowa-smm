package controller

import (
	"github.com/gofiber/fiber/v2"
	"mowa-backend/api/services"
	"mowa-backend/api/utils"
)

type UserController interface {
	Route(router fiber.Router)
	GetUser(ctx *fiber.Ctx) error
}

type userController struct {
	service services.UserService
}

func (u *userController) Route(router fiber.Router) {
	router.Get("/:id", u.GetUser)
}

func NewUserController(service services.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (u *userController) GetUser(ctx *fiber.Ctx) error {
	userId := uint32(1)
	user, err := u.service.GetUser(ctx.Context(), userId)
	if err != nil {
		return err
	}

	return ctx.JSON(utils.ApiResponseSuccess("User fetched successfully", user))
}
