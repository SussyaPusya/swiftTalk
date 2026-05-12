package models

import "time"

type User struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

const (
	AccessTokenTTL  = 15 * time.Minute
	RefreshTokenTTL = 24 * time.Hour
)
