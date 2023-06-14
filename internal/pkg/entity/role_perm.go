package entity

type RolePerm struct {
	ID     int64 `db:"id"`
	RoleID int64 `db:"role_id"`
	PermID int64 `db:"perm_id"`
}
