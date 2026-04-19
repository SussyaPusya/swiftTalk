package models

import "time"

type CreateUserDTO struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordDTO struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type GetUserDTO struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
}

type ChangeNameDTO struct {
	OldName string `json:"oldName"`
	NewName string `json:"newName"`
}
