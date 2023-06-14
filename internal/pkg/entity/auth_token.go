package entity

import "time"

// AuthToken -.
type AuthToken struct {
	ID           int64     `db:"id"`
	UID          int64     `db:"uid"`
	ClientIP     string    `db:"client_ip"`
	RefreshToken string    `db:"refresh_token"`
	CreatedAt    time.Time `db:"created_at"`
	ExpiredAt    time.Time `db:"expired_at"`
}

// AuthTokenDTO -.
type AuthTokenDTO struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}
