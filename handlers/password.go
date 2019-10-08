package handlers

import (
	"encoding/json"
	"github.com/atsushinee/golang-publish-web/db"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/service"
	"github.com/atsushinee/golang-publish-web/utils"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type PasswordModify struct {
	Username    string `json:"username"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

func ModifyPasswordHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := make(map[interface{}]interface{})
	service.ShowMenu(w, r, data)

	m := r.FormValue("m")
	if len(m) > 0 {
		data["Message"] = utils.B64UrlDecode(m)
	}
	utils.Render(w, "password-modify.html", data)
}

func ModifyPasswordVerifyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var msg string
	defer func() {
		err := recover()
		result := &models.Result{}
		if err != nil || len(msg) > 0 {
			result.Message = msg
			result.Code = 0
		} else {
			result.Message = utils.B64UrlEncode("密码修改成功，请重新登录")
			result.Code = 1
		}
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg = "参数异常"
		panic(err)
	}
	passwordModify := new(PasswordModify)
	err = json.Unmarshal(bytes, passwordModify)
	if err != nil {
		msg = "参数异常"
		panic(err)
	}
	user := db.GetUser(passwordModify.Username)
	if user == nil {
		msg = "该用户不存在"
		return
	}
	if user.Password != utils.MD5(passwordModify.OldPassword) {
		msg = "原密码错误"
		return
	}
	err = db.UpdateUserPassword(user.Username, utils.MD5(passwordModify.NewPassword))
	if err != nil {
		msg = "修改密码失败"
		panic(err)
	}
}
