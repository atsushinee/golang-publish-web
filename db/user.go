package db

import (
	"fmt"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/utils"
	"github.com/jeanphorn/log4go"
	"strconv"
	"strings"
)

func GetUser(username string) *models.User {
	user := &models.User{
		Username: username,
	}
	b, err := x.Get(user)
	if err != nil {
		log4go.Error("GetUser error:", err)
		return nil
	}
	if !b {
		return nil
	}
	return user
}

func GetAllUsers() ([]*models.User, error) {
	users := []*models.User{}
	rows, err := x.Rows(new(models.User))
	if err != nil {
		log4go.Error("GetAllUsers Rows error:", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		i := new(models.User)
		err := rows.Scan(i)
		if err != nil {
			log4go.Error("GetAllUsers rows.Scan error:", err)
			return nil, err
		}
		users = append(users, i)
	}
	return users, nil
}

func GetUserById(id int64) *models.User {
	user := &models.User{}
	b, err := x.ID(id).Get(user)
	if err != nil {
		log4go.Error("GetUserById error:", err)
		return nil
	}
	if !b {
		return nil
	}
	return user
}

func CreateUser(username, name, pwd string) bool {
	user := &models.User{
		Username:   username,
		Name:       name,
		Password:   pwd,
		CreateTime: utils.NowTimeString(),
	}
	_, err := x.InsertOne(user)
	if err != nil {
		log4go.Error("CreateUser error:", err)
		return false
	}
	return true
}

func CreateNewUser(username, name, roleList string) bool {
	session := x.NewSession()
	defer session.Close()

	user := &models.User{
		Username:   username,
		Name:       name,
		Password:   utils.DefaultPassword(),
		CreateTime: utils.NowTimeString(),
	}

	_, err := session.InsertOne(user)
	if err != nil {
		session.Rollback()
		log4go.Error("CreateUser error:", err)
		return false
	}
	_, err = session.Get(user)
	if err != nil {
		session.Rollback()
		return false
	}
	roles := strings.Split(roleList, "#")
	for i := 0; i < len(roles)-1; i++ {
		rid, err := strconv.ParseInt(roles[i], 10, 16)
		if err != nil {
			session.Rollback()
			return false
		}
		_, err = session.InsertOne(&models.UserRole{
			Uid: user.Id,
			Rid: rid,
		})
		if err != nil {
			session.Rollback()
			return false
		}
	}
	session.Commit()
	return true
}

func UpdateUser(user *models.User, ip string) {
	_, err := x.Exec("update user set last_login_ip=? , last_login_time=? where id=?",
		ip, utils.NowTimeString(), user.Id)
	if err != nil {
		log4go.Error("UpdateUser error:", err)
		panic(err)
	}
}

func UpdateUserPassword(username string, password string) error {
	_, err := x.Exec("update user set password=?  where username=?",
		password, username)
	if err != nil {
		log4go.Error(fmt.Sprintf("UpdateUserPassword error: %s", err))
		return err
	}
	return nil
}

func GetUserPermission(uid int64) map[string]map[string]bool {
	userPermissionMap := make(map[string]map[string]bool)
	userRoles := GetUserRolesByUid(uid)

	for _, userRole := range userRoles {
		rolePermissions := GetRolePermissionsByRid(userRole.Rid)
		for _, permission := range rolePermissions {
			if userPermissionMap[permission.AllowAccessPath] == nil {
				userPermissionMap[permission.AllowAccessPath] = make(map[string]bool)
			}
			if len(permission.Operation) > 0 {
				operations := strings.Split(permission.Operation, ";")
				for _, operation := range operations {
					userPermissionMap[permission.AllowAccessPath][operation] = true
				}
			}
		}
	}
	return userPermissionMap
}

func GetUserRoleMenu(uid int64) []*models.Menu {
	menuMap := make(map[string]*models.Menu)
	userRoles := GetUserRolesByUid(uid)
	for _, userRole := range userRoles {
		roleMenus := GetRoleMenusByRid(userRole.Rid)
		for _, roleMenu := range roleMenus {
			menu := GetMenuByMid(roleMenu.Mid)
			if menu != nil {
				menuMap[menu.Name] = menu
			}
		}
	}
	var menus []*models.Menu
	for _, v := range menuMap {
		menus = append(menus, v)
	}
	for i := 0; i < len(menus); i++ {
		for j := 1; j < len(menus)-i; j++ {
			if menus[j-1].Xh < menus[j].Xh {
				menus[j-1], menus[j] = menus[j], menus[j-1]
			}
		}
	}
	return menus
}

func GetDefaultMenu() []*models.Menu {
	return []*models.Menu{{Path: "/doc", Name: "文档"}, {Path: "/download", Name: "下载"}}
}

func CreateUserRole(uid, pid int64) bool {
	userRole := &models.UserRole{
		Uid: uid,
		Rid: pid,
	}
	_, err := x.InsertOne(userRole)
	if err != nil {
		log4go.Error("CreateUserRole error:", err)
		return false
	}
	return true
}

func ResetPassword(uid int64) error {
	_, err := x.Where("id=?", uid).Update(&models.User{Password: utils.DefaultPassword()})
	if err != nil {
		log4go.Error("ResetPassword error:", err)
	}
	return err
}

func UserRoleModify(uid int64, roleList string) error {
	session := x.NewSession()
	defer session.Close()
	_, err := session.Where("uid=?", uid).Delete(new(models.UserRole))
	if err != nil {
		session.Rollback()
		log4go.Error("Delete error:", err)
		return err
	}
	roles := strings.Split(roleList, "#")
	for i := 0; i < len(roles)-1; i++ {
		rid, err := strconv.ParseInt(roles[i], 10, 16)
		if err != nil {
			session.Rollback()
			log4go.Error("ParseInt error:", err)

			return err
		}
		_, err = session.InsertOne(&models.UserRole{
			Uid: uid,
			Rid: rid,
		})
		if err != nil {
			log4go.Error("InsertOne error:", err)

			session.Rollback()
			return err
		}
	}
	session.Commit()
	return nil
}
