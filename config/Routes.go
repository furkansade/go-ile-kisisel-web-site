package config

import (
	"github.com/julienschmidt/httprouter"
	admin "goblog/admin/controllers"
	site "goblog/site/controllers"
	"net/http"
)

func Routes() *httprouter.Router {
	r := httprouter.New()

	// Admin
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/new-repository", admin.Dashboard{}.NewItem)
	r.POST("/admin/add", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)

	r.GET("/admin/certificates", admin.Certificates{}.Index)
	r.GET("/admin/new-certificate", admin.Certificates{}.NewItem)
	r.POST("/admin/certificate/add", admin.Certificates{}.Add)
	r.GET("/admin/certificate/delete/:id", admin.Certificates{}.Delete)

	// Site
	r.GET("/", site.Homepage{}.Index)

	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	r.ServeFiles("/site/assets/*filepath", http.Dir("site/assets"))
	return r
}
