package entity

import "time"

// User -.
type User struct {
	ID            int64     `db:"id"`
	Username      string    `db:"username"`
	Password      string    `db:"password"`
	Nickname      string    `db:"nickname"`
	Salt          string    `db:"salt"`
	Phone         string    `db:"phone"`
	Email         string    `db:"email"`
	Avatar        string    `db:"avatar"`
	LangKey       string    `db:"lang_key"`
	Activated     bool      `db:"activated"`
	ActivationKey string    `db:"activation_key"`
	ResetKey      string    `db:"reset_key"`
	ResetDate     time.Time `db:"reset_date"`
	CreatedAt     time.Time `db:"created_at"`
	CreatedBy     string    `db:"created_by"`
	UpdatedAt     time.Time `db:"updated_at"`
	UpdatedBy     string    `db:"updated_by"`
}

// ActiveUser -.
type ActiveUser struct {
	UID      int64
	Username string
	ClientIP string
	Roles    []string
}
