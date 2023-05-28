// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/lugavin/go-scaffold/internal/pkg/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Translation -.
	Translation interface {
		Translate(context.Context, entity.Translation) (entity.Translation, error)
		History(context.Context) ([]entity.Translation, error)
	}

	// TranslationRepo -.
	TranslationRepo interface {
		Store(context.Context, entity.Translation) error
		GetHistory(context.Context) ([]entity.Translation, error)
	}

	// TranslationWebAPI -.
	TranslationWebAPI interface {
		Translate(entity.Translation) (entity.Translation, error)
	}
)

type (
	// AuthToken -.
	AuthToken interface {
		CreateAuthToken(ctx context.Context, uid int64, clientIP string) (entity.AuthTokenDTO, error)
	}

	// AuthTokenRepo -.
	AuthTokenRepo interface {
		Store(context.Context, entity.AuthToken) error
	}
)
