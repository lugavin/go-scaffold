package repo_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/lugavin/go-scaffold/internal/pkg/entity"
	"github.com/lugavin/go-scaffold/internal/pkg/usecase/repo"
	"github.com/lugavin/go-scaffold/pkg/mysql"
)

func TestStoreAuthToken(t *testing.T) {
	t.Skip()
	ms, err := mysql.New(os.Getenv("OP_DB_UNIT_TEST"))
	assert.NoError(t, err)

	err = repo.NewAuthTokenRepo(ms).Store(context.Background(), entity.AuthToken{
		UID:          102,
		ClientIP:     "172.10.10.1",
		RefreshToken: uuid.NewString(),
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(30 * 24 * time.Hour),
	})
	assert.NoError(t, err)
}

func TestGetAuthToken(t *testing.T) {
	ms, err := mysql.New(os.Getenv("OP_DB_UNIT_TEST"))
	assert.NoError(t, err)

	refreshToken := "943f0e7a-fad2-11ed-a6c7-0242ac120002"
	row, err := repo.NewAuthTokenRepo(ms).
		GetAuthToken(context.Background(), refreshToken)
	assert.NoError(t, err)
	assert.NotNil(t, row)
}
