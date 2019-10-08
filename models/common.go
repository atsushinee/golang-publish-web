package models

type User struct {
	Id            int64
	Username      string `xorm:"unique notnull"`
	Name          string
	Password      string
	CreateTime    string
	LastLoginTime string
	LastLoginIp   string
}

type Session struct {
	Id        int64
	SessionId string
	Uid       int64
	TTL       string
}
type UserRole struct {
	Uid int64 `xorm:"pk"`
	Rid int64 `xorm:"pk"`
}

type Role struct {
	Id   int64
	Name string `xorm:"unique notnull"`
}

type RolePermission struct {
	Id              int64
	Rid             int64 `xorm:"notnull"`
	AllowAccessPath string
	Operation       string
}

type Menu struct {
	Id   int64
	Name string
	Path string
	Xh   int
}

type RoleMenu struct {
	Rid int64 `xorm:"pk"`
	Mid int64 `xorm:"pk"`
}

type AdminMenu struct {
	Id   int64
	Name string `xorm:"unique notnull"`
	Path string `xorm:"unique notnull"`
}
