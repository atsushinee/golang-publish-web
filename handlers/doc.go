package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/atsushinee/golang-publish-web/db"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/atsushinee/golang-publish-web/service"
	"github.com/atsushinee/golang-publish-web/utils"
	"github.com/jeanphorn/log4go"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"strconv"
	"time"
)

func DocListHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer func() {
		err := recover()
		if err != nil {
			log4go.Error(fmt.Sprintf("DocListHandler err: %s", err))
			w.WriteHeader(500)
			w.Write([]byte("internal error"))
			return
		}
	}()
	data := make(map[interface{}]interface{})
	service.ShowMenu(w, r, data)

	data["IsDoc"] = true
	data["Projects"] = db.GetAllProjectWithDocs()

	utils.Render(w, "doc.html", data)
}

func DocDetailHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer func() {
		err := recover()
		if err != nil {
			log4go.Error(fmt.Sprintf("DocDetailHandler err: %s", err))
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

	projects := db.GetAllProjectWithDocs()
	var docs []*models.Doc
	for _, project := range projects {
		if pidInt64 == project.Id {
			docs = project.Docs
			data["Pid"] = pidInt64
			data["Project"] = project
			break
		}
	}
	for _, doc := range docs {
		user := db.GetUserById(doc.Uid)
		if user != nil {
			doc.Author = user.Name
		}
		i, err := db.CountDocViewLog(doc.Id)
		if err != nil {
			panic(err)
		}
		doc.ViewCount = i
	}

	data["Projects"] = projects
	data["Docs"] = docs

	utils.IsTheRole(data)
	utils.Render(w, "doc-detail.html", data)
}

func ViewDocToRefreshHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	count, err := db.CountDocViewLog(cidInt64)
	if err != nil {
		result.Code = 0
		result.Message = "系统内部错误"
		return
	}
	result.Code = 1
	result.Message = strconv.FormatInt(count, 10)
}

func DocViewHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer func() {
		err := recover()
		if err != nil {
			log4go.Error(fmt.Sprintf("DocPdfHandler err: %s", err))
			w.WriteHeader(500)
			w.Write([]byte("internal error"))
			return
		}
	}()
	id := ps.ByName("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("doc not found"))
		return
	}
	s, _, _ := service.AuthSession(w, r)
	if !service.CheckPermission(w, r, s.Uid, "get") {
		return
	}

	doc := db.GetDocById(idInt64)
	if doc == nil {
		w.WriteHeader(404)
		w.Write([]byte("doc not found"))
		return
	}

	file, err := os.Open(models.DocDir + doc.Url)

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("doc not found"))
		return
	}
	defer file.Close()

	if doc.Type == "" {
		w.WriteHeader(404)
		w.Write([]byte("type not found"))
		return
	}

	go func() {
		db.RecordDocViewLog(s.Uid, doc.Id)
	}()
	http.ServeContent(w, r, doc.Name, time.Now(), file)
}
