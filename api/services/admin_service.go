package services

import (
	"mowa-backend/api/dto"
	"mowa-backend/api/utils"
	db "mowa-backend/db/sqlc"
	"mowa-backend/internal/database"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap/zapcore"
)

type AdminService interface {
	DashboardStatistics(ctx *fiber.Ctx) (dto.DashboardResponse, error)
	GetLastOrder(ctx *fiber.Ctx, limit int32) ([]dto.GetLastOrderResponse, error)
	GetAdmins(ctx *fiber.Ctx, currentPage int, pageSize int) ([]db.GetAdminsRow, error)
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
		TotalUsers:      int64(stats),
		TotalOrders:     int64(0),
		TotalRevenue:    int64(0),
		TotalTicketOpen: int64(0),
	}, nil
}

// GetLastOrder implements AdminService.
func (a *adminService) GetLastOrder(ctx *fiber.Ctx, limit int32) ([]dto.GetLastOrderResponse, error) {
	orders, err := a.Queries.GetLastOrders(ctx.Context(), limit)
	if err != nil {
		a.log.PrintStdout(ctx.Context(), zapcore.ErrorLevel, "Failed to get last order", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return nil, nil
	}

	result := make([]dto.GetLastOrderResponse, 0)
	for _, v := range orders {
		result = append(result, dto.GetLastOrderResponse{
			Name:    v.FullName,
			Product: v.Name.String,
			Total:   int64(v.Total),
			Price:   int64(v.Price.Float64),
			Status:  utils.Status(v.Status),
			Date:    v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return result, nil
}

// GetAdmins implements AdminService.
// Subtle: this method shadows the method (*Queries).GetAdmins of adminService.Queries.
func (a *adminService) GetAdmins(ctx *fiber.Ctx, currentPage int, pageSize int) ([]db.GetAdminsRow, error) {
	admins, err := a.Queries.GetAdmins(ctx.Context(), db.GetAdminsParams{
		Limit:  int32(pageSize),
		Offset: int32((currentPage - 1) * pageSize),
	})

	if err != nil {
		a.log.PrintStdout(ctx.Context(), zapcore.ErrorLevel, "Failed to get admins", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return nil, err
	}

	return admins, nil
}
