package controller

import (
	"mowa-backend/api/dto"
	"mowa-backend/api/services"
	"mowa-backend/api/utils"

	"github.com/gofiber/fiber/v2"
)

type AdminController interface {
	Route(router fiber.Router)
	DashboardStatistics(c *fiber.Ctx) error
	GetAdmins(c *fiber.Ctx) error
	GetLastOrders(c *fiber.Ctx) error
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
	router.Get("/admins", a.GetAdmins)
	router.Get("/last-orders", a.GetLastOrders)
}

// GetAdmins - get all admins role
func (a *adminController) GetAdmins(c *fiber.Ctx) error {
	admins, err := a.adminService.GetAdmins(c)
	if err != nil {
		return err
	}

	return c.JSON(utils.ApiResponseSuccess("success", admins))
}

// GetLastOrders - get last orders with limit.
func (a *adminController) GetLastOrders(c *fiber.Ctx) error {
	var request dto.GetLastOrdersRequest
	if err := c.QueryParser(&request); err != nil {
		return err
	}

	orders, err := a.adminService.GetLastOrder(c, request.Limit)
	if err != nil {
		return err
	}

	return c.JSON(utils.ApiResponseSuccess("success", orders))
}
