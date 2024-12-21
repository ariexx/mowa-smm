package services

import (
	"context"
	"go.uber.org/zap/zapcore"
	"mowa-backend/api/utils"
	db "mowa-backend/db/sqlc"
	"mowa-backend/internal/database"

	"github.com/go-resty/resty/v2"
)

type JAPProviderService struct {
	APIKey string
	APIURL string
	*db.Queries
	restyClient *resty.Client
	logService  LoggerService
}

func NewJAPProviderService(apiKey string) *JAPProviderService {
	return &JAPProviderService{
		APIKey:      "95af236fed2b3365abfe7b567dd5d395",
		APIURL:      "https://justanotherpanel.com/api/v2xxx",
		Queries:     db.New(database.New().DB()),
		restyClient: resty.New(),
		logService:  NewLoggerService(),
	}
}

func (j *JAPProviderService) HandleOrder(order Order) (ProviderResponse, error) {
	response, err := j.restyClient.R().
		SetBody(map[string]interface{}{
			"key":      j.APIKey,
			"action":   "add",
			"service":  order.ServiceID,
			"link":     order.Link,
			"quantity": order.Quantity,
			// Add other order fields as needed
		}).
		SetResult(&ProviderResponse{}).
		Post(j.APIURL)
	if err != nil {
		return ProviderResponse{}, err
	}

	result := response.Result().(*ProviderResponse)
	return utils.UnPtr(result), nil
}

func (j *JAPProviderService) Status(orderID string) (JAPStatusResponse, error) {
	response, err := j.restyClient.R().
		SetBody(map[string]interface{}{
			"key":    j.APIKey,
			"action": "status",
			"order":  orderID,
		}).
		SetResult(&ProviderResponse{}).
		Post(j.APIURL)
	if err != nil {
		return JAPStatusResponse{}, err
	}

	result := response.Result().(*JAPStatusResponse)
	return utils.UnPtr(result), nil
}

func (j *JAPProviderService) MultiStatus(orderIDs []string) (map[string]JAPStatusResponse, error) {
	response, err := j.restyClient.R().
		SetBody(map[string]interface{}{
			"key":    j.APIKey,
			"action": "status",
			"orders": orderIDs,
		}).
		SetResult(&map[string]JAPStatusResponse{}).
		Post(j.APIURL)
	if err != nil {
		return nil, err
	}

	result := response.Result().(*map[string]JAPStatusResponse)
	return utils.UnPtr(result), nil
}

func (j *JAPProviderService) Services() ([]JAPServicesResponse, error) {
	response, err := j.restyClient.R().
		SetBody(map[string]interface{}{
			"key":    j.APIKey,
			"action": "services",
		}).
		SetResult(&[]JAPServicesResponse{}).
		Post(j.APIURL)
	if err != nil {
		return nil, nil
	}

	result := response.Result().(*[]JAPServicesResponse)

	return utils.UnPtr(result), nil
}

func (j *JAPProviderService) Refill(orderID string) (JAPRefillResponse, error) {
	response, err := j.restyClient.R().
		SetBody(map[string]interface{}{
			"key":    j.APIKey,
			"action": "refill",
			"order":  orderID,
		}).
		SetResult(&JAPRefillResponse{}).
		Post(j.APIURL)
	if err != nil {
		return JAPRefillResponse{}, err
	}

	result := response.Result().(*JAPRefillResponse)
	return utils.UnPtr(result), nil
}

func (j *JAPProviderService) MultiRefill(orderIDs []string) ([]JAPRefillResponse, error) {
	response, err := j.restyClient.R().
		SetBody(map[string]interface{}{
			"key":    j.APIKey,
			"action": "refill",
			"orders": orderIDs,
		}).
		SetResult(&[]JAPRefillResponse{}).
		Post(j.APIURL)
	if err != nil {
		return nil, err
	}

	result := response.Result().(*[]JAPRefillResponse)
	return utils.UnPtr(result), nil
}

func (j *JAPProviderService) RefillStatus(refillID string) (JAPRefillStatusResponse, error) {
	response, err := j.restyClient.R().
		SetBody(map[string]interface{}{
			"key":    j.APIKey,
			"action": "refill_status",
			"refill": refillID,
		}).
		SetResult(&JAPRefillStatusResponse{}).
		Post(j.APIURL)
	if err != nil {
		return JAPRefillStatusResponse{}, err
	}

	result := response.Result().(*JAPRefillStatusResponse)
	return utils.UnPtr(result), nil
}

func (j *JAPProviderService) MultiRefillStatus(refillIDs []string) ([]JAPRefillStatusResponse, error) {
	response, err := j.restyClient.R().
		SetBody(map[string]interface{}{
			"key":     j.APIKey,
			"action":  "refill_status",
			"refills": refillIDs,
		}).
		SetResult(&[]JAPRefillStatusResponse{}).
		Post(j.APIURL)
	if err != nil {
		return nil, err
	}

	result := response.Result().(*[]JAPRefillStatusResponse)
	return utils.UnPtr(result), nil
}

func (j *JAPProviderService) Cancel(orderIDs []string) (JAPCancelResponse, error) {
	response, err := j.restyClient.R().
		SetBody(map[string]interface{}{
			"key":    j.APIKey,
			"action": "cancel",
			"orders": orderIDs,
		}).
		SetResult(&JAPCancelResponse{}).
		Post(j.APIURL)
	if err != nil {
		return JAPCancelResponse{}, err
	}

	result := response.Result().(*JAPCancelResponse)
	return utils.UnPtr(result), nil
}

func (j *JAPProviderService) Balance() (JAPBalanceResponse, error) {
	response, err := j.restyClient.R().
		SetBody(map[string]interface{}{
			"key":    j.APIKey,
			"action": "balance",
		}).
		SetResult(&JAPBalanceResponse{}).
		Post(j.APIURL)
	if err != nil {
		j.logService.PrintStdout(context.Background(), zapcore.ErrorLevel, "Error getting balance", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return JAPBalanceResponse{}, err
	}

	result := response.Result().(*JAPBalanceResponse)
	return utils.UnPtr(result), nil
}
