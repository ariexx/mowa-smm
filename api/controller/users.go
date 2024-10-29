package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"mowa-backend/api/dto"
	"mowa-backend/api/middlewares"
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
	router.Get("/:id", middlewares.JwtMiddleware(), u.GetUser)
}

func NewUserController(service services.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (u *userController) GetUser(ctx *fiber.Ctx) error {
	var request dto.GetUserByIdRequest
	if err := ctx.ParamsParser(&request); err != nil {
		return fmt.Errorf("%s", "Invalid request")
	}

	user, err := u.service.GetUser(ctx.Context(), request.Id)
	if err != nil {
		return err
	}

	return ctx.JSON(utils.ApiResponseSuccess("User fetched successfully", user))
}
