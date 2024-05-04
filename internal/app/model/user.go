package model

type RegisterResponse struct {
	Message string         `json:"message"`
	Data    UserWithAccess `json:"data"`
}
type User struct {
	IDStr     string `json:"id" db:"-"`
	ID        int64  `json:"-" db:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt" db:"createdAt"`
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

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt" db:"createdAt"`
}
