package models

import "time"

type Session struct {
	Token     string    `json:"token"`
	UserID    string    `json:"user_id"`
	ExpireAt  time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
}
