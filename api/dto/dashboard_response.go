package dto

// DashboardResponse - response for the admin dashboard
type DashboardResponse struct {
	TotalUsers    int64 `json:"total_users"`
	TotalOrders	 int64 `json:"total_orders"`
	TotalRevenue	int64 `json:"total_revenue"`
	TotalTicketOpen int64 `json:"total_ticket_open"`
}
