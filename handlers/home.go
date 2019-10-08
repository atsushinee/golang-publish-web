package handlers

import (
	"github.com/atsushinee/golang-publish-web/service"
	"github.com/atsushinee/golang-publish-web/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := make(map[interface{}]interface{})
	service.ShowMenu(w, r, data)
	data["IsHome"] = true
	utils.Render(w, "home.html", data)
}
