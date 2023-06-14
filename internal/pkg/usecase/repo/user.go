package repo

import "github.com/lugavin/go-scaffold/pkg/mysql"

// UserRepo -.
type UserRepo struct {
	*mysql.Mysql
}

// NewUserRepo -.
func NewUserRepo(ms *mysql.Mysql) *UserRepo {
	return &UserRepo{ms}
}
