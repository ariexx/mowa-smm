package dto

// DashboardResponse - response for the admin dashboard
type DashboardResponse struct {
	TotalUsers    int64 `json:"total_users"`
	TotalOrders	 int64 `json:"total_orders"`
	TotalRevenue	int64 `json:"total_revenue"`
	TotalTicketOpen int64 `json:"total_ticket_open"`
}

// GetLastOrderResponse - response for the last order
type GetLastOrderResponse struct {
	Name string `json:"name"`
	Product string `json:"product"`
	Total int64 `json:"total"`
	Price int64 `json:"price"`
	Status string `json:"status"` //Completed, Pending, Cancelled
	Date string `json:"date"`
}

type GetLastOrdersRequest struct {
	Limit int32 `json:"limit" query:"limit"`
}
