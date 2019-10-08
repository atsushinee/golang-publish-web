package db

import (
	"github.com/atsushinee/golang-publish-web/data"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/toolkits/file"
	"log"
)

const dbName = "data/publish.db"

var x *xorm.Engine

func init() {

	var err error

	if !file.IsExist(dbName) {
		//err = os.MkdirAll(path.Dir(dbName), os.ModePerm)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//_, err = os.Create(dbName)
		//		//if err != nil {
		//		//	log.Fatal(err)
		//		//}
		err = data.RestoreAsset("", dbName)
		if err != nil {
			log.Fatal(err)
		}
	}
	x, err = xorm.NewEngine("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	err = x.Sync(new(models.User),
		new(models.Project),
		new(models.Product),
		new(models.Application),
		new(models.ApplicationDownloadLog),
		new(models.Session),
		new(models.Role),
		new(models.UserRole),
		new(models.RolePermission),
		new(models.Menu),
		new(models.RoleMenu),
		new(models.AdminMenu),
		new(models.Doc),
		new(models.DocViewLog),
	)
	if err != nil {
		log.Fatal(err)
	}
}
