package repo

import (
	"context"
	"fmt"

	"github.com/lugavin/go-scaffold/internal/pkg/entity"
	"github.com/lugavin/go-scaffold/pkg/mysql"
)

// RoleRepo -.
type RoleRepo struct {
	*mysql.Mysql
}

// NewRoleRepo -.
func NewRoleRepo(ms *mysql.Mysql) *RoleRepo {
	return &RoleRepo{ms}
}

func (r *RoleRepo) GetRoles(ctx context.Context, uid int64) ([]entity.Role, error) {
	query := "SELECT a.id, a.code, a.name, a.remark FROM sys_role a JOIN sys_user_role b ON a.id = b.role_id WHERE b.user_id = ?"
	//rows, err := r.Pool.QueryContext(ctx, query, uid)
	//if err != nil {
	//	return nil, fmt.Errorf("RoleRepo - GetRoles - r.Pool.Query: %w", err)
	//}
	//defer rows.Close()

	entities := make([]entity.Role, 0, _defEntityCap)
	//for rows.Next() {
	//	e := entity.Role{}
	//	if err = rows.Scan(&e.ID, &e.Code, &e.Name, &e.Remark); err != nil {
	//		return nil, fmt.Errorf("RoleRepo - GetRoles - rows.Scan: %w", err)
	//	}
	//	entities = append(entities, e)
	//}
	if err := r.Pool.SelectContext(ctx, &entities, query, uid); err != nil {
		return nil, fmt.Errorf("RoleRepo - GetRoles - r.Pool.SelectContext: %w", err)
	}

	return entities, nil
}
