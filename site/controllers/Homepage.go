package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goblog/site/helpers"
	"goblog/site/models"
	"html/template"
	"net/http"
)

type Homepage struct{}

func (homepage Homepage) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Repos"] = models.Repo{}.GetAll()
	data["Certificates"] = models.Certificate{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}
