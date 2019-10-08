package utils

import (
	"errors"
	"github.com/atsushinee/golang-publish-web/models"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const viewDir = "./views/"
const viewCustomDir = viewDir + "custom"

func unescaped(x string) interface{} { return template.HTML(x) }
func Render(w http.ResponseWriter, tplName string, data interface{}) {
	var views []string
	views = append(views, viewDir+tplName)

	filepath.Walk(viewCustomDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			views = append(views, path)
		}
		return err
	})
	t := template.New(tplName)
	t = t.Funcs(template.FuncMap{
		"unescaped": unescaped,
	})
	t = template.Must(t.ParseFiles(views...))

	t.Execute(w, data)
}

func GetRequestSession(r *http.Request) (*models.Session, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil, err
	}

	value, ok := models.SessionMap.Load(cookie.Value)
	if !ok {
		return nil, errors.New("No cookie found ")
	}
	return value.(*models.Session), nil
}

func GetRequestIp(r *http.Request) string {
	return strings.Split(r.RemoteAddr, ":")[0]
}

func IsTheRole(data map[interface{}]interface{}) {
	for _, menu := range data["Menus"].([]*models.Menu) {
		name := strings.Replace(menu.Path, "/", "", -1)
		data["Is"+strings.ToUpper(name[:1])+name[1:]] = true
	}
}

func SetUserNameCookie(w http.ResponseWriter, val string, age bool) {
	cookie := &http.Cookie{
		Name:  "username",
		Value: val,
	}
	if age {
		cookie.MaxAge = 1 << 32
	}
	http.SetCookie(w, cookie)
}

func SetSessionCookie(w http.ResponseWriter, val string, age bool) {
	cookie := &http.Cookie{
		Name:  "session_id",
		Value: val,
	}
	if age {
		cookie.MaxAge = 1 << 32
	}
	http.SetCookie(w, cookie)
}
func DefaultPassword() string {
	return MD5("111111")
}
