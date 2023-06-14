package repo

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/lugavin/go-scaffold/internal/pkg/entity"
	"github.com/lugavin/go-scaffold/pkg/mysql"
)

var (
	authTokenTableName = "auth_token"
	authTokenColumns   = []string{"uid", "client_ip", "refresh_token", "created_at", "expired_at"}
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
	//query, args, err := r.Builder.
	//	Insert(authTokenTableName).
	//	Columns(authTokenColumns...).
	//	Values(t.UID, t.ClientIP, t.RefreshToken, t.CreatedAt, t.ExpiredAt).
	//	ToSql()
	//if err != nil {
	//	return fmt.Errorf("AuthTokenRepo - Store - r.Builder: %w", err)
	//}
	query := "INSERT INTO auth_token (uid, client_ip, refresh_token, created_at, expired_at) VALUES (:uid, :client_ip, :refresh_token, :created_at, :expired_at)"

	//_, err = r.Pool.ExecContext(ctx, query, args...)
	_, err := r.Pool.NamedExecContext(ctx, query, &t)
	if err != nil {
		return fmt.Errorf("AuthTokenRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}

// GetAuthToken -.
func (r *AuthTokenRepo) GetAuthToken(ctx context.Context, refreshToken string) (*entity.AuthToken, error) {
	query, args, err := r.Builder.
		Select(authTokenColumns...).
		From(authTokenTableName).
		Where(sq.Eq{"refresh_token": refreshToken}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("AuthTokenRepo - SelectByRefreshToken - r.Builder: %w", err)
	}

	ent := entity.AuthToken{}
	//err = r.Pool.QueryRowContext(ctx, query, args...).
	//	Scan(&ent.UID, &ent.ClientIP, &ent.RefreshToken, &ent.CreatedAt, &ent.ExpiredAt)
	err = r.Pool.GetContext(ctx, &ent, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("AuthTokenRepo - SelectByRefreshToken - row.Scan: %w", err)
	}

	return &ent, nil
}
