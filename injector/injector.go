//go:build wireinject
// +build wireinject

package injector

import (
	"mowa-backend/api/controller"
	"mowa-backend/api/services"

	"github.com/google/wire"
)

func InitializeUserController() controller.UserController {
	wire.Build(
		controller.NewUserController,
		services.NewUserService,
	)

	return nil
}

func InitializeAdminController() controller.AdminController {
	wire.Build(
		controller.NewAdminController,
		services.NewAdminService,
	)

	return nil
}
