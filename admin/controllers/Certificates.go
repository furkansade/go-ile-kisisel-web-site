package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"html/template"
	"io"
	"net/http"
	"os"
)

type Certificates struct{}

func (certificates Certificates) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("certificates/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Certificates"] = models.Certificate{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

func (certificates Certificates) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles(helpers.Include("certificates/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	view.ExecuteTemplate(w, "index", nil)
}

func (certificates Certificates) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	title := r.FormValue("cert-title")
	description := r.FormValue("cert-desc")
	filter := r.FormValue("cert-filter")
	// UPLOAD
	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("cert-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	// file : aldigim dosya
	// f : file'dan aldigim icerigi yeni olusan ayni isimli dosyaya aktarma islemi
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	// file -> f'e kopyalamak -> io.copy(f, file)
	_, err = io.Copy(f, file)
	// UPLOAD END
	if err != nil {
		fmt.Println(err)
		return
	}

	models.Certificate{
		Title:       title,
		Description: description,
		Picture_url: "uploads/" + header.Filename,
		Filter:      filter,
	}.Add()

	http.Redirect(w, r, "/admin/certificates", http.StatusSeeOther)
}

func (certificates Certificates) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	certificate := models.Certificate{}.Get(params.ByName("id"))
	certificate.Delete()
	http.Redirect(w, r, "/admin/certificates", http.StatusSeeOther)
}
