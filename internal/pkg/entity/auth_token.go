package entity

import "time"

// AuthToken -.
type AuthToken struct {
	ID           int64
	UID          int64
	ClientIP     string
	RefreshToken string
	CreatedAt    time.Time
	ExpiredAt    time.Time
}

// AuthTokenDTO -.
type AuthTokenDTO struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}
