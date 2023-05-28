package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/lugavin/go-scaffold/internal/pkg/entity"
)

// AuthTokenUseCase -.
type AuthTokenUseCase struct {
	repo AuthTokenRepo
}

// NewAuthTokenUseCase -.
func NewAuthTokenUseCase(r AuthTokenRepo) *AuthTokenUseCase {
	return &AuthTokenUseCase{
		repo: r,
	}
}

func (uc *AuthTokenUseCase) CreateAuthToken(ctx context.Context, uid int64, clientIP string) (entity.AuthTokenDTO, error) {
	refreshToken := uuid.NewString()
	e := entity.AuthToken{
		UID:          uid,
		ClientIP:     clientIP,
		RefreshToken: refreshToken,
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(30 * 24 * time.Hour),
	}

	ret := entity.AuthTokenDTO{}
	if err := uc.repo.Store(ctx, e); err != nil {
		return ret, fmt.Errorf("AuthTokenUseCase - CreateAuthToken - uc.repo.Store: %w", err)
	}
	ret.RefreshToken = refreshToken
	ret.ExpiresIn = e.ExpiredAt.Unix()
	ret.AccessToken = "" // todo create jwt

	return ret, nil
}
