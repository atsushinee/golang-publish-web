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
	"strconv"
)

func AdminListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	data := make(map[interface{}]interface{})
	service.ShowMenu(w, r, data)
	menus, err := db.GetAllAdminMenus()
	if err != nil {
		panic(err)
	}
	data["AdminMenus"] = menus
	utils.Render(w, "admin.html", data)
}

func AdminUserListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := make(map[interface{}]interface{})
	service.ShowMenu(w, r, data)
	menus, err := db.GetAllAdminMenus()
	if err != nil {
		panic(err)
	}
	data["AdminMenus"] = menus
	users, err := db.GetAllUsers()
	if err != nil {
		panic(err)
	}
	roles, err := db.GetAllRoles()
	if err != nil {
		panic(err)
	}
	data["Users"] = users
	data["Roles"] = roles
	utils.Render(w, "user-list.html", data)
}

func AdminRoleListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := make(map[interface{}]interface{})
	service.ShowMenu(w, r, data)
	menus, err := db.GetAllAdminMenus()
	if err != nil {
		panic(err)
	}
	data["AdminMenus"] = menus
	users, err := db.GetAllUsers()
	if err != nil {
		panic(err)
	}
	roles, err := db.GetAllRoles()
	if err != nil {
		panic(err)
	}
	data["Users"] = users
	data["Roles"] = roles
	utils.Render(w, "role-list.html", data)
}

func AdminProductListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := make(map[interface{}]interface{})
	service.ShowMenu(w, r, data)
	menus, err := db.GetAllAdminMenus()
	if err != nil {
		panic(err)
	}
	data["AdminMenus"] = menus

	projects := db.GetAllProjectWithProducts()

	data["Projects"] = projects

	utils.Render(w, "product-list.html", data)
}

func AdminDownloadApplicationLogHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	result := &models.Result{}
	defer func(r *models.Result) {
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}(result)
	cid := ps.ByName("cid")
	cidInt64, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		result.Code = 0
		result.Message = "无效资源id"
		return
	}
	logs, err := db.GetDownloadLogByApplicationId(cidInt64)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		return
	}
	result.Code = 1
	result.Data = logs
}

func AdminDocViewLogHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	result := &models.Result{}
	defer func(r *models.Result) {
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}(result)
	cid := ps.ByName("cid")
	cidInt64, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		result.Code = 0
		result.Message = "无效资源id"
		return
	}
	logs, err := db.GetDocViewLogByDocId(cidInt64)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		return
	}
	result.Code = 1
	result.Data = logs
}

func AdminGetUserRoleHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	result := &models.Result{}
	defer func(r *models.Result) {
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}(result)
	uid := ps.ByName("uid")
	uidInt64, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		result.Code = 0
		result.Message = "无效用户id"
		return
	}
	roles := db.GetUserRolesByUid(uidInt64)
	roleList := []string{}
	for _, r := range roles {
		roleList = append(roleList, strconv.FormatInt(r.Rid, 10))
	}
	result.Code = 1
	result.Data = roleList
}

func AdminResetPasswordHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	result := &models.Result{}
	defer func(r *models.Result) {
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}(result)
	uid := ps.ByName("uid")
	uidInt64, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		result.Code = 0
		result.Message = "无效用户id"
		return
	}
	user := db.GetUserById(uidInt64)
	if user == nil {
		result.Code = 0
		result.Message = "用户id不存在"
		return
	}
	err = db.ResetPassword(uidInt64)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		return
	}
	result.Code = 1
}

type NewUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Roles    string `json:"roles"`
}

func AdminAddUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result := &models.Result{}
	defer func(r *models.Result) {
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}(result)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		result.Code = 0
		result.Message = "数据读取异常"
		return
	}

	user := new(NewUser)
	err = json.Unmarshal(data, user)
	if err != nil {
		result.Code = 0
		result.Message = "数据格式异常"
		return
	}

	ok := db.CreateNewUser(user.Username, user.Name, user.Roles)
	if !ok {
		result.Code = 0
		result.Message = "系统内部错误"
		return
	}
	result.Code = 1
}

type UserRoleModify struct {
	Uid   int64  `json:"uid"`
	Roles string `json:"roles"`
}

func AdminUserRoleModifyHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result := &models.Result{}
	defer func(r *models.Result) {
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}(result)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		result.Code = 0
		result.Message = "数据读取异常"
		return
	}

	user := new(UserRoleModify)
	err = json.Unmarshal(data, user)
	if err != nil {
		result.Code = 0
		result.Message = "数据格式异常"
		return
	}

	err = db.UserRoleModify(user.Uid, user.Roles)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		return
	}
	result.Code = 1
}

type NewProduct struct {
	ProjectId int64  `json:"project_id"`
	Name      string `json:"name"`
}

func AdminAddProductHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result := &models.Result{}
	defer func(r *models.Result) {
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}(result)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		result.Code = 0
		result.Message = "数据读取异常"
		return
	}

	product := new(NewProduct)
	err = json.Unmarshal(data, product)
	if err != nil {
		result.Code = 0
		result.Message = "数据格式异常"
		return
	}

	err = db.CreateProduct(product.ProjectId, product.Name)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		return
	}
	result.Code = 1
}

type NewProject struct {
	Name string `json:"name"`
}

func AdminAddProjectHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result := &models.Result{}
	defer func(r *models.Result) {
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}(result)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		result.Code = 0
		result.Message = "数据读取异常"
		return
	}

	project := new(NewProject)
	err = json.Unmarshal(data, project)
	if err != nil {
		result.Code = 0
		result.Message = "数据格式异常"
		return
	}

	err = db.CreateProject(project.Name)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		return
	}
	result.Code = 1
}
