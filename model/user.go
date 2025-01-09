package model

import "time"

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	FullName    string    `json:"full_name"`
	PhoneNumber string    `json:"phone_number"`
	Nickname    string    `json:"nickname"`
	QQ          string    `json:"qq"`
	Avatar      string    `json:"avatar"`
	Gender      string    `json:"gender"`
	Bio         string    `json:"bio"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Role        string    `json:"role"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTKey struct {
	Key string `json:"key"`
}

type ChangePasswordAndUsernameUser struct {
	OldPassword string `json:"old_password"`
	NewUsername string `json:"new_username"`
	NewPassword string `json:"new_password"`
}

type ChangeInfoUser struct {
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Nickname    string `json:"nickname"`
	QQ          string `json:"qq"`
	Avatar      string `json:"avatar"`
	Gender      string `json:"gender"`
	Bio         string `json:"bio"`
}
