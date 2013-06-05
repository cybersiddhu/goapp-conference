package controller

import (
	"html/template"
	"net/http"
	"github.com/cybersiddhu/gobioweb"
)


type TemplateData struct {
	Message string
}

type Action struct {
	 gobioweb.Controller
}

func Welcome(w http.ResponseWriter, r *http.Request, t *template.Template) *gobioweb.AppError {
	d := &TemplateData{Message: "Welcome to GoSidd framework"}
	err := t.ExecuteTemplate(w, "welcome.tmpl", d)
	if err != nil {
		return &gobioweb.AppError{Error: err, Code: 500, Message: "Cannot display template"}
	}
	return nil
}

func ListAbstract(w http.ResponseWriter, r *http.Request, t *template.Template) *gobioweb.AppError {
	return nil
}

func CreateAbstract(w http.ResponseWriter, r *http.Request, t *template.Template) *gobioweb.AppError {
	return nil
}

func NewAbstract(w http.ResponseWriter, r *http.Request, t *template.Template) *gobioweb.AppError {
	return nil
}

func ShowAbstract(w http.ResponseWriter, r *http.Request, t *template.Template) *gobioweb.AppError {
	return nil
}
