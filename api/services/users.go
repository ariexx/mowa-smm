package services

import (
	"context"
	"go.uber.org/zap/zapcore"
	"mowa-backend/api/utils"
	db "mowa-backend/db/sqlc"
	"mowa-backend/internal/database"
)

type UserService interface {
	GetUser(ctx context.Context, id uint32) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	IsUserActive(ctx context.Context, id uint32) bool
}

type userService struct {
	*db.Queries
	log LoggerService
}

func NewUserService() UserService {
	return &userService{
		Queries: db.New(database.New().DB()),
		log:     NewLoggerService(),
	}
}

func (u *userService) GetUser(ctx context.Context, id uint32) (db.User, error) {
	//check if user is admin
	isAdmin := ctx.Value("user").(db.User).Role
	if isAdmin != "admin" {
		return db.User{}, utils.ErrUserNotFound
	}

	user, err := u.Queries.GetUser(ctx, id)
	if err != nil {
		u.log.PrintStdout(ctx, zapcore.ErrorLevel, "GetUser", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return db.User{}, utils.MapError(err)
	}

	return user, nil
}

func (u *userService) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	user, err := u.Queries.GetUserByEmail(ctx, email)
	if err != nil {
		u.log.PrintStdout(ctx, zapcore.ErrorLevel, "GetUserByEmail", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return db.User{}, err
	}

	return user, nil
}

func (u *userService) IsUserActive(ctx context.Context, id uint32) bool {
	user, err := u.GetUser(ctx, id)
	if err != nil {
		return false
	}

	return !user.DeletedAt.Valid
}
