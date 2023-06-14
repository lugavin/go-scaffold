package errors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lugavin/go-scaffold/internal/pkg/errors"
)

func TestAppError(t *testing.T) {
	appErr := errors.New(errors.ParamsValidationErrCode, "params invalid")
	assert.True(t, appErr.Is(errors.ErrInvalidParams))
}
