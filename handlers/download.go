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
	"path/filepath"
	"strconv"
)

func DownloadListHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer func() {
		err := recover()
		if err != nil {
			log4go.Error(fmt.Sprintf("DownloadListHandler err: %s", err))
			w.WriteHeader(500)
			w.Write([]byte("internal error"))
			return
		}
	}()
	data := make(map[interface{}]interface{})
	service.ShowMenu(w, r, data)

	data["Projects"] = db.GetAllProjectWithProducts()

	utils.Render(w, "download.html", data)
}

//
//func RateHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	result := &models.Result{}
//	defer func(r *models.Result) {
//		bytes, _ := json.Marshal(result)
//		w.Write(bytes)
//	}(result)
//	cid := ps.ByName("cid")
//	cidInt64, err := strconv.ParseInt(cid, 10, 64)
//	if err != nil {
//		result.Code = 0
//		result.Message = "无效资源id"
//		return
//	}
//	err = db.RateItemById(cidInt64)
//	if err != nil {
//		result.Code = 0
//		result.Message = "点赞失败"
//		return
//	}
//	item := db.GetItemById(cidInt64)
//	if item == nil {
//		result.Code = 0
//		result.Message = "资源未找到"
//		return
//	}
//	result.Code = 1
//	result.Message = strconv.FormatInt(item.RateCount, 10)
//}
//
func DownloadFileToRefreshHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	count, err := db.CountApplicationDownloadLog(cidInt64)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		return
	}
	result.Code = 1
	result.Message = strconv.FormatInt(count, 10)
}

func DownloadDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer func() {
		err := recover()
		if err != nil {
			log4go.Error(fmt.Sprintf("DownloadDetailHandler err: %s", err))
			w.WriteHeader(500)
			w.Write([]byte("internal error"))
			return
		}
	}()
	pid := ps.ByName("pid")
	pidInt64, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("project not found"))
		return
	}

	data := make(map[interface{}]interface{})
	service.ShowMenu(w, r, data)

	projects := db.GetAllProjectWithProducts()
	var products []*models.Product
	for _, project := range projects {
		if project.Id == pidInt64 {
			products = project.Products
			break
		}
	}

	for _, product := range products {
		product.Applications = db.GetApplicationsByProductId(product.Id)
		for _, application := range product.Applications {
			count, err := db.CountApplicationDownloadLog(application.Id)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte("internal error"))
				return
			}
			application.DownloadCount = count
		}
	}
	data["Projects"] = projects

	data["Products"] = products
	data["Pid"] = pidInt64
	utils.IsTheRole(data)
	utils.Render(w, "download-detail.html", data)
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	cid := ps.ByName("cid")
	cidInt64, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("application not found"))
		return
	}

	application := db.GetApplicationById(cidInt64)
	if application == nil {
		w.WriteHeader(404)
		w.Write([]byte("application not found"))
		return
	}
	s, _, ok := service.AuthSession(w, r)
	if !ok {
		w.WriteHeader(500)
		w.Write([]byte("session timeout"))
		return
	}

	if !service.CheckPermission(w, r, s.Uid, "get") {
		return
	}

	file, err := os.Open(models.UploadDir + application.Url)
	_, fileName := filepath.Split(application.Url)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("not found"))
		return
	}

	defer file.Close()

	go db.RecordApplicationDownloadLog(s.Uid, cidInt64)

	reader := bufio.NewReader(file)
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-disposition", "attachment; filename="+fileName)
	io.Copy(w, reader)
}
