package repo

import (
	"context"
	"fmt"

	"github.com/lugavin/go-scaffold/internal/entity"
	"github.com/lugavin/go-scaffold/pkg/postgres"
)

const _defaultEntityCap = 64

// TranslationPgRepo -.
type TranslationPgRepo struct {
	*postgres.Postgres
}

// NewTranslationPgRepo -.
func NewTranslationPgRepo(pg *postgres.Postgres) *TranslationPgRepo {
	return &TranslationPgRepo{pg}
}

// GetHistory -.
func (r *TranslationPgRepo) GetHistory(ctx context.Context) ([]entity.Translation, error) {
	sql, _, err := r.Builder.
		Select("source, destination, original, translation").
		From("history").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Translation, 0, _defaultEntityCap)
	for rows.Next() {
		e := entity.Translation{}
		if err = rows.Scan(&e.Source, &e.Destination, &e.Original, &e.Translation); err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetHistory - rows.Scan: %w", err)
		}
		entities = append(entities, e)
	}

	return entities, nil
}

// Store -.
func (r *TranslationPgRepo) Store(ctx context.Context, t entity.Translation) error {
	sql, args, err := r.Builder.
		Insert("history").
		Columns("source, destination, original, translation").
		Values(t.Source, t.Destination, t.Original, t.Translation).
		ToSql()
	if err != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}
