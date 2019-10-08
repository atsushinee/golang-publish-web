package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/atsushinee/golang-publish-web/db"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/service"
	"github.com/atsushinee/golang-publish-web/utils"
	"github.com/jeanphorn/log4go"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
)

func PublishHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := make(map[interface{}]interface{})
	service.ShowMenu(w, r, data)

	data["Projects"] = db.GetAllProjects()
	utils.Render(w, "publish.html", data)
}

func PublishApplicationPostHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	result := &models.Result{}
	defer func(r *models.Result) {
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}(result)
	projectId := r.FormValue("projectId")
	_, err := strconv.ParseInt(projectId, 10, 64)
	if err != nil {
		result.Code = 0
		result.Message = "所选项目错误"
		log4go.Error("ParseInt error:", err)
		return
	}
	productId := r.FormValue("productId")
	productIdInt64, err := strconv.ParseInt(productId, 10, 64)
	if err != nil {
		result.Code = 0
		result.Message = "所选应用分类错误"
		log4go.Error("ParseInt error:", err)
		return
	}
	versionCode := r.FormValue("versionCode")
	versionCodeInt64, err := strconv.ParseInt(versionCode, 10, 64)

	if err != nil {
		result.Code = 0
		result.Message = "所输入VersionCode不满足要求"
		log4go.Error("ParseInt error:", err)
		return
	}

	filename := r.FormValue("filename")
	versionName := r.FormValue("versionName")
	productName := r.FormValue("productName")
	projectName := r.FormValue("projectName")
	desc := r.FormValue("desc")
	file, _, err := r.FormFile("file")
	if err != nil {
		result.Code = 0
		result.Message = "获取文件错误"
		log4go.Error("FormFile error:", err)
		return
	}
	defer file.Close()

	createTime := utils.NowTimeString()
	url := "/" + projectName + "/" + productName + "/" + versionName + "_" + versionCode + "/" + filename
	itemPath := models.UploadDir + url
	if utils.IsExist(itemPath) {
		result.Code = 0
		result.Message = "应用文件已存在"
		log4go.Error("IsExist error:", err)
		return
	}
	err = os.MkdirAll(path.Dir(itemPath), os.ModePerm)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		log4go.Error("MkdirAll error:", err)
		return
	}

	newFile, err := os.Create(itemPath)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		log4go.Error("Create error:", err)
		return
	}
	defer newFile.Close()
	reader := bufio.NewReader(file)
	_, err = io.Copy(newFile, reader)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		log4go.Error("io.Copy error:", err)
		return
	}
	session, _, _ := service.AuthSession(w, r)
	ok := db.CreateApplication(productIdInt64, session.Uid, productName, url, versionName, versionCodeInt64, desc, createTime)
	if !ok {
		result.Code = 0
		result.Message = "系统内部错误"
		log4go.Error("CreateItem error:", err)
		err := os.RemoveAll(itemPath)
		if err != nil {
			log4go.Error(fmt.Sprintf("RemoveAll error: %s", err))
		}
		return
	}

	result.Code = 1
}

func PublishDocPostHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	result := &models.Result{}
	defer func(r *models.Result) {
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}(result)
	projectId := r.FormValue("projectId")
	projectIdInt64, err := strconv.ParseInt(projectId, 10, 64)
	if err != nil {
		result.Code = 0
		result.Message = "所选项目错误"
		log4go.Error("ParseInt error:", err)
		return
	}
	projectName := r.FormValue("projectName")
	fileFullName := r.FormValue("fileName")
	name, ext := utils.FileNameSplit(fileFullName)
	file, _, err := r.FormFile("file")
	if err != nil {
		result.Code = 0
		result.Message = "获取文件错误"
		log4go.Error("FormFile error:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		log4go.Error("ReadAll error:", err)
		return
	}
	createTime := utils.NowTimeString()
	url := "/" + projectName + "/" + fileFullName
	itemPath := models.DocDir + url
	if utils.IsExist(itemPath) {
		result.Code = 0
		result.Message = "应用文件已存在"
		log4go.Error("IsExist error:", err)
		return
	}
	err = os.MkdirAll(path.Dir(itemPath), os.ModePerm)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		log4go.Error("MkdirAll error:", err)
		return
	}
	newFile, err := os.Create(itemPath)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		log4go.Error("Create error:", err)
		return
	}
	defer newFile.Close()
	_, err = io.Copy(newFile, reader)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		log4go.Error("Copy error:", err)
		return
	}
	session, _, _ := service.AuthSession(w, r)
	err = db.CreateDoc(projectIdInt64, session.Uid, name, ext, url, createTime)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		err := os.RemoveAll(itemPath)
		if err != nil {
			log4go.Error(fmt.Sprintf("RemoveAll error: %s", err))
		}
		return
	}

	result.Code = 1
}

func PublishGetApplicationHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	result := &models.Result{}
	defer func(r *models.Result) {
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}(result)
	pid := ps.ByName("pid")
	pidInt64, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		result.Code = 0
		result.Message = "所选项目错误"
		log4go.Error("ParseInt error:", err)
		return
	}
	products := db.GetProductionsByProjectId(pidInt64)
	result.Code = 1
	result.Data = products
}
