package entity

type Role struct {
	ID     int64  `db:"id"`
	Code   string `db:"code"`
	Name   string `db:"name"`
	Remark string `db:"remark"`
}
