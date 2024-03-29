package repo

import (
	"context"
	"fmt"

	"github.com/lugavin/go-scaffold/internal/pkg/entity"
	"github.com/lugavin/go-scaffold/pkg/mysql"
)

var (
	historyTableName = "history"
	historyColumns   = []string{"source", "destination", "original", "translation"}
)

const _defEntityCap = 64

// TranslationRepo -.
type TranslationRepo struct {
	*mysql.Mysql
}

// NewTranslationRepo -.
func NewTranslationRepo(ms *mysql.Mysql) *TranslationRepo {
	return &TranslationRepo{ms}
}

// GetHistory -.
func (r *TranslationRepo) GetHistory(ctx context.Context) ([]entity.Translation, error) {
	sql, _, err := r.Builder.
		Select(historyColumns...).
		From(historyTableName).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Builder: %w", err)
	}

	rows, err := r.Pool.QueryContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Translation, 0, _defEntityCap)
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
func (r *TranslationRepo) Store(ctx context.Context, t entity.Translation) error {
	sql, args, err := r.Builder.
		Insert(historyTableName).
		Columns(historyColumns...).
		Values(t.Source, t.Destination, t.Original, t.Translation).
		ToSql()
	if err != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Builder: %w", err)
	}

	_, err = r.Pool.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("TranslationRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}
