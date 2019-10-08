package main

import (
	"github.com/atsushinee/golang-publish-web/middleware"
	"github.com/atsushinee/golang-publish-web/routers"
	"github.com/atsushinee/golang-publish-web/service"
	"log"
	"net/http"
)

func main() {

	service.LoadSessions()
	handler := middleware.NewMiddlewareHandler(routers.RegisterRouter())
	log.Fatal(http.ListenAndServe(":1912", handler))
}
