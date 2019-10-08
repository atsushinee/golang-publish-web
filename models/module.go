package models

type Project struct {
	Id       int64
	Name     string     `xorm:"unique notnull"`
	Products []*Product `xorm:"-"`
	Docs     []*Doc     `xorm:"-"`
}

type Application struct {
	Id                   int64
	ProductId            int64 `xorm:"notnull"`
	Uid                  int64
	VersionName          string
	VersionCode          int64
	Desc                 string
	Url                  string
	CreateTime           string
	Name                 string `xorm:"-"`
	DownloadCount        int64  `xorm:"-"`
	LastDownloadUserName string `xorm:"-"`
	Author               string `xorm:"-"`
	FileName             string `xorm:"-"`
	FileSize             string `xorm:"-"`
}

type Product struct {
	Id           int64
	ProjectId    int64
	Name         string         `xorm:"notnull"`
	Applications []*Application `xorm:"-"`
}

type ApplicationDownloadLog struct {
	Id               int64
	ApplicationId    int64
	DownloadUserId   int64
	DownloadTime     string
	DownloadUserName string `xorm:"-"`
}

type Doc struct {
	Id         int64
	Name       string
	Pid        int64
	Uid        int64
	Url        string
	Type       string
	CreateTime string
	Author     string `xorm:"-"`
	ViewCount  int64  `xorm:"-"`
}

type DocViewLog struct {
	Id           int64
	DocId        int64
	ViewUserId   int64
	ViewTime     string
	ViewUserName string `xorm:"-"`
}
