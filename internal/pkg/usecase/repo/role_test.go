package repo_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lugavin/go-scaffold/internal/pkg/usecase/repo"
	"github.com/lugavin/go-scaffold/pkg/mysql"
)

func TestGetRoles(t *testing.T) {
	ms, err := mysql.New(os.Getenv("OP_DB_UNIT_TEST"))
	assert.NoError(t, err)

	var uid int64 = 101
	rows, err := repo.NewRoleRepo(ms).
		GetRoles(context.Background(), uid)
	assert.NoError(t, err)
	assert.True(t, len(rows) > 0)
}
