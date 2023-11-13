package api

type RegisterRequest struct {
	Username   string
	Passphrase string
	Email      string
}

type ValidationError struct {
	Field string
	Error string
}
type AppError struct {
	Message   string
	ErrorCode string
	Status    int
	Errors    []ValidationError
}
