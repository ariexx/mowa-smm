package services

import (
	"mowa-backend/api/dto"
	db "mowa-backend/db/sqlc"
	"mowa-backend/internal/database"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap/zapcore"
)

type AdminService interface {
	DashboardStatistics(ctx *fiber.Ctx) (dto.DashboardResponse, error)
}

type adminService struct {
	*db.Queries
	log LoggerService
}

func NewAdminService() AdminService {
	return &adminService{
		Queries: db.New(database.New().DB()),
		log:     NewLoggerService(),
	}
}

// DashboardStatistics - get statistics for the admin dashboard
func (a *adminService) DashboardStatistics(ctx *fiber.Ctx) (dto.DashboardResponse, error) {

	// get the statistics
	stats, err := a.Queries.GetStatistics(ctx.Context())
	if err != nil {
		a.log.PrintStdout(ctx.Context(), zapcore.ErrorLevel, "Failed to get statistics", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return dto.DashboardResponse{}, nil
	}

	return dto.DashboardResponse{
		TotalUsers: int64(stats),
		TotalOrders: int64(0),
		TotalRevenue: int64(0),
		TotalTicketOpen: int64(0),
	}, nil
}
