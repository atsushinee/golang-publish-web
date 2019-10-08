package db

import (
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/jeanphorn/log4go"
)

func GetUserRolesByUid(uid int64) []*models.UserRole {
	var roles []*models.UserRole

	rows, err := x.Where("uid=?", uid).Rows(new(models.UserRole))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		r := new(models.UserRole)
		err = rows.Scan(r)
		if err != nil {
			panic(err)
		}
		roles = append(roles, r)
	}
	return roles
}

func CreateRole(name string) bool {
	role := &models.Role{
		Name: name,
	}
	_, err := x.InsertOne(role)
	if err != nil {
		log4go.Error("CreateRole error:", err)
		return false
	}
	return true
}

func CreateRolePermission(rid int64, allowAccessPath, operation string) bool {
	rolePermission := &models.RolePermission{
		Rid:             rid,
		AllowAccessPath: allowAccessPath,
		Operation:       operation,
	}
	_, err := x.InsertOne(rolePermission)
	if err != nil {
		log4go.Error("CreateRolePermission error:", err)
		return false
	}
	return true
}

func GetRolePermissionsByRid(rid int64) []*models.RolePermission {
	var rolePermissions []*models.RolePermission
	rows, err := x.Where("rid=?", rid).Rows(new(models.RolePermission))
	if err != nil {
		log4go.Error("GetRolePermissionsByRid error:", err)
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rp := &models.RolePermission{}
		err := rows.Scan(rp)
		if err != nil {
			panic(err)
		}
		rolePermissions = append(rolePermissions, rp)
	}
	return rolePermissions
}

func CreateRoleMenu(rid, mid int64) bool {
	menu := &models.RoleMenu{
		Rid: rid,
		Mid: mid,
	}
	_, err := x.InsertOne(menu)
	if err != nil {
		log4go.Error("CreateRoleMenu error:", err)
		return false
	}
	return true
}

func GetAllRoles() ([]*models.Role, error) {
	roles := []*models.Role{}
	rows, err := x.Rows(new(models.Role))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		i := new(models.Role)
		err = rows.Scan(i)
		if err != nil {
			return nil, err
		}
		roles = append(roles, i)
	}
	return roles, nil
}
