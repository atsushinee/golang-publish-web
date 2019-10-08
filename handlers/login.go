package handlers

import (
	"github.com/atsushinee/golang-publish-web/db"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/service"
	"github.com/atsushinee/golang-publish-web/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer func() {
		if recover() != nil {
			msg := utils.B64UrlEncode("系统内部错误")
			http.Redirect(w, r, "/login?msg="+msg, 302)
		}
	}()
	data := make(map[interface{}]interface{})
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		pwd := r.FormValue("pwd")
		autoLogin := r.FormValue("autoLogin")
		user, ok := service.AuthUser(username, pwd)
		if ok {
			if user.Password == utils.DefaultPassword() {
				utils.SetUserNameCookie(w, username+"$"+utils.B64UrlEncode(user.Name), true)
				m := utils.B64UrlEncode("请先修改初始密码")
				http.Redirect(w, r, "/password/modify?m="+m, 302)
				return
			}
			db.UpdateUser(user, utils.GetRequestIp(r))
			session := db.InsertSession(user.Id)
			models.SessionMap.Store(session.SessionId, session)
			utils.SetSessionCookie(w, session.SessionId, len(autoLogin) > 0)
			utils.SetUserNameCookie(w, username+"$"+utils.B64UrlEncode(user.Name), true)
			http.Redirect(w, r, "/", 302)
			return
		} else {
			utils.SetUserNameCookie(w, username+"$", true)
			msg := utils.B64UrlEncode("用户名或密码不正确")
			http.Redirect(w, r, "/login?msg="+msg, 302)
			return
		}
	}
	msg := r.FormValue("msg")
	if len(msg) > 0 {
		data["Message"] = utils.B64UrlDecode(msg)
	}
	data["Menus"] = db.GetDefaultMenu()
	data["UrlPath"] = r.URL.Path
	utils.Render(w, "login.html", data)
}
