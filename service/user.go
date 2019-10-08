package service

import (
	"github.com/atsushinee/golang-publish-web/db"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/utils"
	"net/http"
	"strings"
)

func AuthUser(username, pwd string) (*models.User, bool) {
	if len(username) == 0 || len(pwd) == 0 {
		return nil, false
	}
	user := db.GetUser(username)
	if user == nil || user.Password != utils.MD5(pwd) {
		return nil, false
	}
	return user, true
}

func GetUsernameFromCookie(r *http.Request) (string, string) {
	cookie, err := r.Cookie("username")
	if err != nil {
		return "", ""
	}
	values := strings.Split(cookie.Value, "$")
	return values[0], utils.B64UrlDecode(values[1])
}
