package services

type Order struct {
	Quantity  int
	ServiceID int
	Action    *string
	Link      string
}

type ProviderResponse struct {
	Order int `json:"order"`
}

type JAPServicesResponse struct {
	Service  int    `json:"service"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Rate     string `json:"rate"`
	Min      string `json:"min"`
	Max      string `json:"max"`
	Refill   bool   `json:"refill"`
	Cancel   bool   `json:"cancel"`
}

type Provider interface {
	HandleOrder(order Order) ProviderResponse
}

type JAPStatusResponse struct {
	Charge     string `json:"charge"`
	StartCount string `json:"start_count"`
	Status     string `json:"status"`
	Remains    string `json:"remains"`
	Currency   string `json:"currency"`
}

type JAPRefillResponse struct {
	Order  string `json:"order,omitempty"`
	Refill string `json:"refill"`
}

type JAPRefillStatusResponse struct {
	Status interface{} `json:"status"` //it can be string or object with key "error"
	Refill string      `json:"refill"`
}

type JAPCancelResponse struct {
	Order  int         `json:"order"`
	Cancel interface{} `json:"cancel"` // it can be string or object with key "error"
}

type JAPBalanceResponse struct {
	Balance  string `json:"balance"`
	Currency string `json:"currency"`
}
