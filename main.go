package main

import (
	admin "goblog/admin/models"
	"goblog/config"
	site "goblog/site/models"
	"net/http"
)

func main() {
	admin.Repo{}.Migrate()
	admin.Certificate{}.Migrate()
	site.Repo{}.Migrate()

	/* site.Repo{
		Title:       "Go ile WEB Programlama",
		Description: "Admin ve site şablonu ile blog sitesi.",
		Repo_Url:    "https://github.com/furkansade/web-development-with-go",
		Icon:        "bi bi-globe",
	}.Add() */

	http.ListenAndServe(":9090", config.Routes())
}
