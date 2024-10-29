package server

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"mowa-backend/api/utils"
	"mowa-backend/internal/database"
	"strconv"
)

type FiberServer struct {
	*fiber.App
	context.Context
	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader:  "mowa-backend",
			AppName:       "mowa-backend",
			StrictRouting: true,
			ErrorHandler:  ErrorHandler,
		}),
		Context: context.Background(),
		db:      database.New(),
	}

	return server
}

func (f *FiberServer) RegisterMiddlewares() {
	f.App.Use(f.compress(), f.cors(), f.helmet(), f.idempotency(), f.limiter(), f.logger(), f.requestid(), f.recover())
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if ok {
		return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponseFail(ValidationErrors(err), "Validation Error"))
	}

	var errorFiber *fiber.Error
	if ok = errors.As(err, &errorFiber); ok {
		switch errorFiber.Code {
		case fiber.StatusNotFound:
			return ctx.Status(fiber.StatusNotFound).JSON(utils.ApiResponseError("Not found", strconv.Itoa(errorFiber.Code)))
		case fiber.StatusInternalServerError:
			return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ApiResponseError("Something wrong, please contact admin", "500"))
		default:
			return ctx.Status(errorFiber.Code).JSON(utils.ApiResponseError(errorFiber.Message, strconv.Itoa(errorFiber.Code)))
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(utils.ApiResponseFail(nil, err.Error()))
}

func ValidationErrors(err error) any {
	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if !ok {
		return nil
	}

	errorsMap := make(map[string]string)
	for _, field := range validationErrors {
		errorsMap[field.Field()] = field.Tag()
	}

	return errorsMap
}
