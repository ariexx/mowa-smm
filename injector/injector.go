//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"mowa-backend/api/controller"
	"mowa-backend/api/services"
)

func InitializeUserController() controller.UserController {
	wire.Build(
		controller.NewUserController,
		services.NewUserService,
	)

	return nil
}
