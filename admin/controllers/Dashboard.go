package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"html/template"
	"net/http"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Repos"] = models.Repo{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "index", nil)
}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	title := r.FormValue("repo-title")
	description := r.FormValue("repo-desc")
	repoUrl := r.FormValue("repo-url")
	icons := r.FormValue("repo-icons")

	models.Repo{
		Title:       title,
		Description: description,
		Repo_Url:    repoUrl,
		Icon:        icons,
	}.Add()

	http.Redirect(w, r, "/admin", http.StatusSeeOther) // response + reques + nereye yonlendirecegi + hangi port ile
	// bu olmazsa boş bi sayfa donecektir

}

func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	repo := models.Repo{}.Get(params.ByName("id"))
	repo.Delete()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Repo"] = models.Repo{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	repo := models.Repo{}.Get(params.ByName("id"))
	title := r.FormValue("repo-title")
	description := r.FormValue("repo-desc")
	repoUrl := r.FormValue("repo-url")
	icons := r.FormValue("repo-icons")

	repo.Updates(models.Repo{
		Title:       title,
		Description: description,
		Repo_Url:    repoUrl,
		Icon:        icons,
	})
	http.Redirect(w, r, "/admin/edit/"+params.ByName("id"), http.StatusSeeOther)
}
