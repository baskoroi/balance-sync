package dto

type OKResponse struct{}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type DepositRequest struct {
	UserID uint `json:"user_id"`
	Amount int  `json:"amount"`
}

type SpendRequest struct {
	UserID uint `json:"user_id"`
	Amount int  `json:"amount"`
}
