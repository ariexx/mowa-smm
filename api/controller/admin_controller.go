package controller

import (
	"mowa-backend/api/services"
	"mowa-backend/api/utils"

	"github.com/gofiber/fiber/v2"
)

type AdminController interface {
	Route(router fiber.Router)
	DashboardStatistics(c *fiber.Ctx) error
}

type adminController struct {
	adminService services.AdminService
}


func NewAdminController(adminService services.AdminService) AdminController {
	return &adminController{
		adminService: adminService,
	}
}

// DashboardStatistics implements AdminController.
func (a *adminController) DashboardStatistics(c *fiber.Ctx) error {
	stats, err := a.adminService.DashboardStatistics(c)
	if err != nil {
		return err
	}

	return c.JSON(utils.ApiResponseSuccess("success", stats))
}

// Route implements AdminController.
func (a *adminController) Route(router fiber.Router) {
	router.Get("/dashboard", a.DashboardStatistics)
}
