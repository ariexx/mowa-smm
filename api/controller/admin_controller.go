package controller

import (
	"mowa-backend/api/dto"
	"mowa-backend/api/services"
	"mowa-backend/api/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AdminController interface {
	Route(router fiber.Router)
	DashboardStatistics(c *fiber.Ctx) error
	GetAdmins(c *fiber.Ctx) error
	GetLastOrders(c *fiber.Ctx) error

	// JAP Handler
	GetJAPBalance(c *fiber.Ctx) error
}

type adminController struct {
	adminService services.AdminService
	japProvider  services.JAPProviderService
}

func NewAdminController(adminService services.AdminService) AdminController {
	return &adminController{
		adminService: adminService,
		japProvider:  utils.UnPtr(services.NewJAPProviderService("")),
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

	japRouter := router.Group("/jap")
	{
		japRouter.Get("/balance", a.GetJAPBalance)
	}
}

// GetAdmins - get all admins role
func (a *adminController) GetAdmins(c *fiber.Ctx) error {
	// Get current page and page size from query parameters
	currentPage, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || currentPage < 1 {
		currentPage = 1
	}

	pageSize, err := strconv.Atoi(c.Query("page_size", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// Fetch admins from the service
	admins, err := a.adminService.GetAdmins(c, currentPage, pageSize)
	if err != nil {
		return err
	}

	// Calculate total pages
	totalRecords := len(admins)
	totalPages := (totalRecords + pageSize - 1) / pageSize

	// Create pagination metadata
	pagination := &utils.PaginationMetadata{
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		CurrentPage:  currentPage,
		PageSize:     pageSize,
	}

	// Return paginated response
	return c.JSON(utils.ApiResponseSuccessWithPagination("success", admins, pagination))
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

func (a *adminController) GetJAPBalance(c *fiber.Ctx) error {
	balance, err := a.japProvider.Balance()
	if err != nil {
		return err
	}

	return c.JSON(utils.ApiResponseSuccess("success", balance))
}
