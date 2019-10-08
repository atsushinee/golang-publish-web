package service

import (
	"github.com/atsushinee/golang-publish-web/db"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/utils"
	"net/http"
	"strings"
)

const sep = "/"

func ShowMenu(w http.ResponseWriter, r *http.Request, data map[interface{}]interface{}) *models.Session {
	s, _, ok := AuthSession(w, r)
	data["IsLogin"] = ok
	data["FullPath"] = r.URL.Path
	urlSplits := strings.Split(r.URL.Path, sep)
	if len(urlSplits) > 2 {
		data["UrlPath"] = sep + urlSplits[1]
	} else {
		data["UrlPath"] = r.URL.Path
	}
	if ok {
		data["Name"], data["Username"] = GetUsernameFromCookie(r)
		data["Menus"] = db.GetUserRoleMenu(s.Uid)
	} else {
		data["Menus"] = db.GetDefaultMenu()
	}
	return s
}

func CheckPermission(w http.ResponseWriter, r *http.Request, uid int64, op string) bool {
	userPermissionMap := db.GetUserPermission(uid)

	permission := userPermissionMap[utils.GetBasePath(r.URL.Path)]
	_, ok := permission[op]
	if !ok {
		http.Redirect(w, r, "/permission", 302)
		return false
	}
	return true
}
