package services

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap/zapcore"
	db "mowa-backend/db/sqlc"
	"os"
	"time"
)

type JwtService interface {
	GenerateToken(userId uint32) (access, refresh string)
	ValidateToken(token string) bool
	GetUserFromToken(token string) (db.User, error)
}

type jwtService struct {
	secret      string
	userService UserService
	logService  LoggerService
}

func NewJwtService() JwtService {
	return &jwtService{
		secret:      os.Getenv("JWT_SECRET"),
		userService: NewUserService(),
		logService:  NewLoggerService(),
	}
}

func (j *jwtService) GenerateToken(userId uint32) (access, refresh string) {
	j.logService.PrintStdout(context.Background(), zapcore.InfoLevel, fmt.Sprintf("GenerateToken for user %d", userId))
	//if user is not found, return empty string
	userData, err := j.userService.GetUser(context.Background(), userId)
	if err != nil {
		j.logService.PrintStdout(context.Background(), zapcore.ErrorLevel, "GenerateToken.GetUser", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return "", ""
	}

	//if user is found, generate token
	if userData.ID != 0 {
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": userData,
			"type": "access",
			"iat":  time.Now().Unix(),
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})

		accessTokenString, err := accessToken.SignedString([]byte(j.secret))
		if err != nil {
			j.logService.PrintStdout(context.Background(), zapcore.ErrorLevel, "GenerateToken.accessTokenString", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
			return "", ""
		}

		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": userData,
			"type": "refresh",
			"iat":  time.Now().Unix(),
			"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
		})

		refreshTokenString, err := refreshToken.SignedString([]byte(j.secret))
		if err != nil {
			j.logService.PrintStdout(context.Background(), zapcore.ErrorLevel, "GenerateToken.refreshTokenString", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
			return "", ""
		}

		return accessTokenString, refreshTokenString
	}

	return "", ""
}

func (j *jwtService) ValidateToken(token string) bool {
	j.logService.PrintStdout(context.Background(), zapcore.InfoLevel, "ValidateToken")
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		j.logService.PrintStdout(context.Background(), zapcore.ErrorLevel, "ValidateToken.Parse", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return false
	}

	//check token expired or not
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		j.logService.PrintStdout(context.Background(), zapcore.ErrorLevel, "ValidateToken.ParseWithClaims", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return false
	}

	if claims["exp"] == nil {
		j.logService.PrintStdout(context.Background(), zapcore.InfoLevel, "ValidateToken.Claims", zapcore.Field{Key: "error", Type: zapcore.StringType, String: "exp is nil"})
		return false
	}

	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		j.logService.PrintStdout(context.Background(), zapcore.InfoLevel, "ValidateToken.Claims", zapcore.Field{Key: "error", Type: zapcore.StringType, String: "token expired"})
		return false
	}

	return true
}

func (j *jwtService) GetUserFromToken(token string) (db.User, error) {
	j.logService.PrintStdout(context.Background(), zapcore.InfoLevel, "GetUserFromToken")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		j.logService.PrintStdout(context.Background(), zapcore.ErrorLevel, "GetUserFromToken.ParseWithClaims", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return db.User{}, err
	}

	userData := db.User{}
	if claims["user"] != nil {
		userData = claims["user"].(db.User)
	}

	return userData, nil
}
