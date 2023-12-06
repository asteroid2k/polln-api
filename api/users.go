package api

type CreateUserRequest struct {
	Username   string `json:"username" validate:"required,alphanum,min=4"`
	Passphrase string `json:"passphrase" validate:"required,min=8"`
	Email      string `json:"email" validate:"required,email"`
}

type UserResponse struct {
	UserId   string
	Username string
}
