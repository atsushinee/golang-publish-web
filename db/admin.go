package db

import "github.com/atsushinee/golang-publish-web/models"

func GetAllAdminMenus() ([]*models.AdminMenu, error) {
	menus := []*models.AdminMenu{}
	rows, err := x.Rows(new(models.AdminMenu))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		i := new(models.AdminMenu)
		err := rows.Scan(i)
		if err != nil {
			return nil, err
		}
		menus = append(menus, i)
	}
	return menus, nil
}
