package api

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
