package dto

type GetUserByIdRequest struct {
	Id uint32 `json:"id" params:"id"`
}
