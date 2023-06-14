package repo

import "github.com/lugavin/go-scaffold/pkg/mysql"

// PermRepo -.
type PermRepo struct {
	*mysql.Mysql
}

// NewPermRepo -.
func NewPermRepo(ms *mysql.Mysql) *PermRepo {
	return &PermRepo{ms}
}
