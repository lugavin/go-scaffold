package usecase

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/lugavin/go-scaffold/config"
	"github.com/lugavin/go-scaffold/internal/pkg/entity"
)

// AuthTokenUseCase -.
type AuthTokenUseCase struct {
	repo   AuthTokenRepo
	cfg    config.JWT
	priKey *rsa.PrivateKey
	pubKey *rsa.PublicKey
}

// NewAuthTokenUseCase -.
func NewAuthTokenUseCase(r AuthTokenRepo, c config.JWT) (*AuthTokenUseCase, error) {
	priKey, pubKey, err := resolveSignKey(c.PrivateKey, c.PublicKey)
	if err != nil {
		return nil, err
	}
	return &AuthTokenUseCase{r, c, priKey, pubKey}, nil
}

func (uc *AuthTokenUseCase) CreateAuthToken(ctx context.Context, uid int64, clientIP string) (*entity.AuthTokenDTO, error) {
	token, err := uc.generateToken(uid)
	if err != nil {
		return nil, fmt.Errorf("AuthTokenUseCase - CreateAuthToken - uc.generateAccessToken: %w", err)
	}

	ent := entity.AuthToken{
		UID:          uid,
		ClientIP:     clientIP,
		RefreshToken: uuid.NewString(),
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(time.Duration(uc.cfg.RefreshTokenExp) * time.Second),
	}
	if err = uc.repo.Store(ctx, ent); err != nil {
		return nil, fmt.Errorf("AuthTokenUseCase - CreateAuthToken - uc.repo.Store: %w", err)
	}

	return &entity.AuthTokenDTO{
		AccessToken:  token,
		RefreshToken: ent.RefreshToken,
		ExpiresIn:    ent.ExpiredAt.Unix(),
	}, nil
}

func (uc *AuthTokenUseCase) RenewAuthToken(ctx context.Context, uid int64, refreshToken string) (*entity.AuthTokenDTO, error) {
	row, err := uc.repo.GetAuthToken(ctx, refreshToken)
	if err != nil {
		return nil, fmt.Errorf("AuthTokenUseCase - RenewAuthToken - uc.repo.GetAuthToken: %w", err)
	}
	if row == nil {
		return nil, fmt.Errorf("AuthTokenUseCase - RenewAuthToken - refreshTokenInvalid: %s", refreshToken)
	}
	if row.UID != uid || row.ExpiredAt.Before(time.Now()) {
		return nil, fmt.Errorf("AuthTokenUseCase - RenewAuthToken - refreshTokenExpired: %s", refreshToken)
	}

	token, err := uc.generateToken(uid)
	if err != nil {
		return nil, fmt.Errorf("AuthTokenUseCase - RenewAuthToken - uc.generateAccessToken: %w", err)
	}

	return &entity.AuthTokenDTO{
		AccessToken:  token,
		RefreshToken: row.RefreshToken,
		ExpiresIn:    row.ExpiredAt.Unix(),
	}, nil
}

func (uc *AuthTokenUseCase) generateToken(uid int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Duration(uc.cfg.AccessTokenExp) * time.Second).Unix(),
		"uid": uid,
	})
	return token.SignedString(uc.priKey)
}

func resolveSignKey(priKeyBase64, pubKeyBase64 string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	priKeyBytes, err := base64.StdEncoding.DecodeString(priKeyBase64)
	if err != nil {
		return nil, nil, err
	}
	priKey, err := jwt.ParseRSAPrivateKeyFromPEM(priKeyBytes)
	if err != nil {
		return nil, nil, err
	}

	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyBase64)
	if err != nil {
		return nil, nil, err
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		return nil, nil, err
	}

	return priKey, pubKey, nil
}
