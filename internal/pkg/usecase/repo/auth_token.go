package repo

import (
	"context"
	"fmt"

	"github.com/lugavin/go-scaffold/internal/pkg/entity"
	"github.com/lugavin/go-scaffold/pkg/mysql"
)

// AuthTokenRepo -.
type AuthTokenRepo struct {
	*mysql.Mysql
}

// NewAuthTokenRepo -.
func NewAuthTokenRepo(ms *mysql.Mysql) *AuthTokenRepo {
	return &AuthTokenRepo{ms}
}

// Store -.
func (r *AuthTokenRepo) Store(ctx context.Context, t entity.AuthToken) error {
	sql, args, err := r.Builder.
		Insert("auth_token").
		Columns("uid, client_ip, refresh_token, created_at, expired_at").
		Values(t.UID, t.ClientIP, t.RefreshToken, t.CreatedAt, t.ExpiredAt).
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthTokenRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthTokenRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}
