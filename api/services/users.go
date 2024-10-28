package services

import (
	"context"
	"go.uber.org/zap/zapcore"
	db "mowa-backend/db/sqlc"
)

type UserService interface {
	GetUser(ctx context.Context, id uint32) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	IsUserActive(ctx context.Context, id uint32) bool
}

type userService struct {
	queries *db.Queries
	log     LoggerService
}

func NewUserService(queries *db.Queries) UserService {
	return &userService{
		queries: queries,
		log:     NewLoggerService(),
	}
}

func (u *userService) GetUser(ctx context.Context, id uint32) (db.User, error) {
	user, err := u.queries.GetUser(ctx, id)
	if err != nil {
		u.log.PrintStdout(ctx, zapcore.ErrorLevel, "GetUser", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return db.User{}, err
	}

	return user, nil
}

func (u *userService) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) IsUserActive(ctx context.Context, id uint32) bool {
	//TODO implement me
	panic("implement me")
}
