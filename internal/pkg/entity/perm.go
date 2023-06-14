package entity

type Perm struct {
	ID       int64  `db:"id"`
	Code     string `db:"code"`
	Name     string `db:"name"`
	Type     string `db:"type"`
	URL      string `db:"url"`
	Method   string `db:"method"`
	Seq      int    `db:"seq"`
	Icon     string `db:"icon"`
	IsParent bool   `db:"is_parent"`
	ParentID int64  `db:"parent_id"`
}
