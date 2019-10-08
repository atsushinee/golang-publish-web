package handlers

import (
	"github.com/atsushinee/golang-publish-web/service"
	"github.com/atsushinee/golang-publish-web/utils"
	"github.com/jeanphorn/log4go"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	defer func() {
		recover()
		http.Redirect(w, r, "/", http.StatusFound)
	}()
	session, err := utils.GetRequestSession(r)
	if err != nil {
		log4go.Error("GetRequestSession error:", err)
		panic(err)
	}
	service.Logout(w, session)
}
