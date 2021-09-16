package dto

type AccountCreationResponse struct {
	Account  *Account
	ApiError *ApiError
}

type TransferMoneyRequest struct {
	SenderAccountID    int64
	RecipientAccountID int64
	Cents              int64
}

type TransferMoneyResponse struct {
	Success  bool
	ApiError *ApiError
}

type ReplenishRequest struct {
	AccountID int64
	Cents     int64
}

type ReplenishResponse struct {
	Cents    int64
	ApiError *ApiError
}

type Account struct {
	ID    int64
	Cents int64
}

type ApiError struct {
	Message string
}
