//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/google/wire"
	"mowa-backend/api/controller"
	"mowa-backend/api/services"
	db "mowa-backend/db/sqlc"
)

func InitializeUserController() controller.UserController {
	wire.Build(
		controller.NewUserController(),
		services.NewUserService(),
		db.New(),
	)

	return nil
}
