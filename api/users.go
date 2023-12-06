package api

type CreateUserRequest struct {
	Username   string `json:"username"`
	Passphrase string `json:"passphrase"`
	Email      string `json:"email"`
}

type UserResponse struct {
	UserId   string
	Username string
}
