package middleware

import (
	"github.com/atsushinee/golang-publish-web/db"
	"github.com/atsushinee/golang-publish-web/service"
	"github.com/atsushinee/golang-publish-web/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"regexp"
	"strings"
)

var extraPrefix = []string{
	"/get",
}

var allowPrefix = []string{
	"/static",
	"/static",
	"/login",
	"/password",
	"/permission",
	"/error",
	"/logout",
}
var allowUrl = []string{
	"/",
}

type middleware struct {
	router *httprouter.Router
	p      *PermissionHandler
}

func NewMiddlewareHandler(r *httprouter.Router) *middleware {
	return &middleware{
		p: &PermissionHandler{
			router: r,
		},
		router: r,
	}
}

func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, prefix := range extraPrefix {
		if strings.HasPrefix(r.URL.Path, prefix) {
			m.router.ServeHTTP(w, r)
			return
		}
	}
	for _, prefix := range allowPrefix {
		if strings.HasPrefix(r.URL.Path, prefix) {
			m.router.ServeHTTP(w, r)
			return
		}
	}
	for _, url := range allowUrl {
		if r.URL.Path == url {
			m.router.ServeHTTP(w, r)
			return
		}
	}
	_, msg, ok := service.AuthSession(w, r)
	if !ok {
		if len(msg) > 0 {
			http.Redirect(w, r, "/login?msg="+utils.B64UrlEncode(msg), 302)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
		return
	} else {
		m.p.ServeHTTP(w, r)
	}
}

type PermissionHandler struct {
	router *httprouter.Router
}

func (p *PermissionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s, _, _ := service.AuthSession(w, r)
	user := db.GetUserById(s.Uid)
	if user == nil {
		http.Redirect(w, r, "/error", 302)
		return
	}
	permissionMap := db.GetUserPermission(user.Id)
	var ok bool
	for k, v := range permissionMap {
		if r.URL.Path == k || regexp.MustCompile(`^` + k + `/\w+`).MatchString(r.URL.Path) {
			ok = v["view"]
			break
		}
	}
	if !ok {
		http.Redirect(w, r, "/permission", 302)
		return
	}
	p.router.ServeHTTP(w, r)
}
