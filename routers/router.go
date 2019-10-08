package routers

import (
	"github.com/atsushinee/golang-publish-web/handlers"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func RegisterRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handlers.HomeHandler)

	router.GET("/publish", handlers.PublishHandler)
	router.POST("/publish/application", handlers.PublishApplicationPostHandler)
	router.POST("/publish/doc", handlers.PublishDocPostHandler)
	router.GET("/publish/project/:pid", handlers.PublishGetApplicationHandler)
	router.GET("/login", handlers.LoginHandler)
	router.POST("/login", handlers.LoginHandler)
	router.GET("/logout", handlers.LogoutHandler)
	router.GET("/download", handlers.DownloadListHandler)
	router.GET("/download/project/:pid", handlers.DownloadDetailHandler)
	//router.POST("/download/rate/:cid", handlers.RateHandler)
	router.GET("/download/application/:cid", handlers.DownloadFileHandler)
	router.GET("/download/application-download-count/:cid", handlers.DownloadFileToRefreshHandler)
	router.GET("/password/modify", handlers.ModifyPasswordHandler)
	router.POST("/password/modify", handlers.ModifyPasswordVerifyHandler)
	router.GET("/permission", handlers.PermissionHandler)
	router.GET("/error", handlers.ErrorHandler)
	router.GET("/doc", handlers.DocListHandler)
	router.GET("/doc/project/:pid", handlers.DocDetailHandler)
	router.GET("/doc/view/:id/:name", handlers.DocViewHandler)
	router.GET("/doc/doc-view-count/:cid", handlers.ViewDocToRefreshHandler)
	router.GET("/admin", handlers.AdminListHandler)
	router.GET("/admin/application-download-log/:cid", handlers.AdminDownloadApplicationLogHandler)
	router.GET("/admin/doc-view-log/:cid", handlers.AdminDocViewLogHandler)
	router.GET("/admin/user", handlers.AdminUserListHandler)
	router.GET("/admin/user-role/:uid", handlers.AdminGetUserRoleHandler)
	router.GET("/admin/role", handlers.AdminRoleListHandler)
	router.GET("/admin/product", handlers.AdminProductListHandler)
	router.POST("/admin/user/add", handlers.AdminAddUserHandler)
	router.POST("/admin/user-role-modify", handlers.AdminUserRoleModifyHandler)
	router.POST("/admin/product/add", handlers.AdminAddProductHandler)
	router.POST("/admin/project/add", handlers.AdminAddProjectHandler)
	router.POST("/admin/user/reset-pwd/:uid", handlers.AdminResetPasswordHandler)

	router.NotFound = NotFoundHandler{}
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	return router
}

type NotFoundHandler struct{}

func (NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "404 page not found")
}
