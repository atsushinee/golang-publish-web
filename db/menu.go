package db

import (
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/jeanphorn/log4go"
)

func CreateMenu(menu *models.Menu) bool {

	_, err := x.InsertOne(menu)
	if err != nil {
		log4go.Error("CreateMenu error:", err)
		return false
	}
	return true
}

func GetRoleMenusByRid(rid int64) []*models.RoleMenu {
	var roleMenus []*models.RoleMenu
	rows, err := x.Where("rid=?", rid).Rows(new(models.RoleMenu))
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		roleMenu := &models.RoleMenu{}
		err = rows.Scan(roleMenu)
		if err != nil {
			panic(err)
		}
		roleMenus = append(roleMenus, roleMenu)
	}
	return roleMenus
}

func GetMenuByMid(mid int64) *models.Menu {
	menu := &models.Menu{}
	b, err := x.ID(mid).Get(menu)
	if err != nil {
		log4go.Error("GetMenuById error: ", err.Error())
		return nil
	}
	if !b {
		return nil
	}
	return menu
}
