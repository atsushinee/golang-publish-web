package main

import (
	"fmt"
	"github.com/atsushinee/golang-publish-web/db"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/utils"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test(t *testing.T) {
	fmt.Println(time.Now())
	fmt.Println(utils.NowTimeString())
	fmt.Println(utils.S2T(utils.NowTimeString()))
	uid := db.GetUserRolesByUid(1)
	for _, r := range uid {
		fmt.Printf("%#v", r)
	}

}

func TestGetUserRolesByUid(t *testing.T) {
	uid := db.GetUserRolesByUid(1)
	for _, r := range uid {
		fmt.Printf("%#v", r)
	}
}

func TestGetRolePermissionsByRid(t *testing.T) {
	roles := db.GetUserRolesByUid(1)
	for _, r := range roles {
		rid := db.GetRolePermissionsByRid(r.Rid)
		for _, p := range rid {
			fmt.Printf("%+v\n", p)
		}
	}
}
func TestGetUserPermission(t *testing.T) {
	//db.CreateUserRole(5, 2)

	roles := db.GetUserPermission(2)
	for k, v := range roles {
		fmt.Println(k, v)
	}
}
func TestAlready(t *testing.T) {
	filename := "123.123.12354"

	split := strings.Split(filename, ".")

	fmt.Println(strings.Join(split[:len(split)-1], "."), strings.Join(split[len(split)-1:], ""))

}

func TestMenu(t *testing.T) {
	db.CreateMenu(&models.Menu{Id: 1, Name: "应用发布", Path: "/publish", Xh: 50})
	db.CreateMenu(&models.Menu{Id: 2, Name: "文档", Path: "/doc", Xh: 90})
	db.CreateMenu(&models.Menu{Id: 3, Name: "下载", Path: "/download", Xh: 80})
	db.CreateMenu(&models.Menu{Id: 4, Name: "系统管理", Path: "/admin", Xh: 1})
}
func TestGetUserMenu(t *testing.T) {
	mid := db.GetUserRoleMenu(1)
	fmt.Printf("%v\n", mid)
}
func TestRoleMenu(t *testing.T) {
	db.CreateRoleMenu(1, 4)
	db.CreateRoleMenu(1, 3)
	db.CreateRoleMenu(2, 2)
	db.CreateRoleMenu(2, 1)
	db.CreateRoleMenu(3, 2)
	db.CreateRoleMenu(3, 3)
	db.CreateRoleMenu(4, 2)
	db.CreateRoleMenu(4, 3)
}

func TestCount(t *testing.T) {
	roles := db.GetUserRolesByUid(1)
	roleList := []string{}
	for _, r := range roles {
		roleList = append(roleList, strconv.FormatInt(r.Rid, 10))
		fmt.Printf("%+v\n", r)
	}
	fmt.Println(roleList)
}

func TestDoc(t *testing.T) {
	docs, err := db.GetAllDocsByPid(1)
	if err != nil {
		panic(err)
	}
	for _, doc := range docs {
		fmt.Printf("%+v\n", doc)
	}

}

func TestInitCommon(t *testing.T) {

	db.CreateUser("admin", "管理员", utils.DefaultPassword())
	db.CreateRole("管理员")
	db.CreateRole("开发人员")
	db.CreateRole("部门内部人员")
	db.CreateRole("普通用户")
	db.CreateRolePermission(1, "/doc", "view;modify")
	db.CreateRolePermission(1, "/publish", "view;modify")
	db.CreateRolePermission(1, "/download", "view;modify")
	db.CreateRolePermission(1, "/downloads", "view;modify")

	db.CreateRolePermission(2, "/doc", "view;modify")
	db.CreateRolePermission(2, "/publish", "view;modify")
	db.CreateRolePermission(2, "/download", "view;modify")
	db.CreateRolePermission(2, "/downloads", "view;modify")

	db.CreateRolePermission(3, "/doc", "view")
	db.CreateRolePermission(3, "/download", "view")
	db.CreateRolePermission(3, "/downloads", "view")

	db.CreateRolePermission(4, "/doc", "view")
	db.CreateRolePermission(4, "/download", "view")

}
func TestInitData(t *testing.T) {

	//db.CreateUser("lichun", "李淳", "111111")

	//db.CreateUserRole(1, 1)
	//db.CreateUserRole(1, 2)

	db.CreateUserRole(2, 1)
	db.CreateUserRole(2, 2)
	db.CreateUserRole(2, 3)

	//db.CreateUserRole(3, 2)
	//db.CreateUserRole(3, 3)
	//
	//db.CreateUserRole(4, 2)
	//db.CreateUserRole(4, 3)
	//
	//db.CreateUserRole(5, 2)
	//db.CreateUserRole(5, 3)
	//
	//db.CreateUserRole(6, 2)
	//db.CreateUserRole(6, 3)
	//
	//db.CreateUserRole(7, 3)
	//db.CreateUserRole(8, 3)
	//db.CreateUserRole(9, 3)
	//db.CreateUserRole(10, 3)
	//db.CreateUserRole(11, 3)
}
