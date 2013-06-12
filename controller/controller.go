package controller

import (
	"github.com/cybersiddhu/gobioweb"
	"net/http"
)

type TemplateData struct {
	Message string
}

func Welcome(w http.ResponseWriter, r *http.Request, a *gobioweb.App) *gobioweb.AppError {
	t := a.Template
	d := &TemplateData{Message: "Welcome to GoSidd framework"}
	err := t.ExecuteTemplate(w, "welcome.tmpl", d)
	if err != nil {
		return &gobioweb.AppError{Error: err, Code: 500, Message: "Cannot display template"}
	}
	return nil
}

func ListAbstract(w http.ResponseWriter, r *http.Request, a *gobioweb.App) *gobioweb.AppError {
	return nil
}

func CreateAbstract(w http.ResponseWriter, r *http.Request, t *gobioweb.App) *gobioweb.AppError {
	return nil
}

func NewAbstract(w http.ResponseWriter, r *http.Request, t *gobioweb.App) *gobioweb.AppError {
	return nil
}

func ShowAbstract(w http.ResponseWriter, r *http.Request, t *gobioweb.App) *gobioweb.AppError {
	return nil
}
