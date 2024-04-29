package model

type RegisterResponse struct {
	Message string         `json:"message"`
	Data    UserWithAccess `json:"data"`
}
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserWithAccess struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
