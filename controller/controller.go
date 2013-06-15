package controller

import (
	"github.com/cybersiddhu/gobioweb"
	"net/http"
)

type TemplateData struct {
	Message    string
	FormFields map[string]string
	FormError bool
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

func NewUser(w http.ResponseWriter, r *http.Request, a *gobioweb.App) *gobioweb.AppError {
	f := map[string]string{"fname": "First Name", "lname": "Last Name"}
	err := a.Template.ExecuteTemplate(w, "signup.tmpl", &TemplateData{FormFields: f})

	if err != nil {
		return &gobioweb.AppError{Error: err, Code: 500, Message: "Cannot display template signup.tmpl"}
	}
	return nil
}

func CreateUser(w http.ResponseWriter, r *http.Request, a *gobioweb.App) *gobioweb.AppError {
	 err := r.ParseForm()
	 if err != nil {
		return &gobioweb.AppError{Error: err, Code: 500, Message: "Could not parse form"}
	 }

	 //form data processing
	 email := r.FormValue("email")
	 if email == {
		err := a.Template.ExecuteTemplate(w, "signup.tmpl", &TemplateData{FormError: true,
		Message: "Email missing"})
	if err != nil {
		return &gobioweb.AppError{Error: err, Code: 500, Message: "Cannot display template signup.tmpl"}
	}
	 }

	 pass := r.FormValue("pass")
	 if !pass {
		err := a.Template.ExecuteTemplate(w, "signup.tmpl", &TemplateData{FormError: true,
		Message: "Password missing"})
	if err != nil {
		return &gobioweb.AppError{Error: err, Code: 500, Message: "Cannot display template signup.tmpl"}
	}
	 }
	return nil
}

func Login(w http.ResponseWriter, r *http.Request, a *gobioweb.App) *gobioweb.AppError {
	return nil
}

func CreateSession(w http.ResponseWriter, r *http.Request, a *gobioweb.App) *gobioweb.AppError {
	return nil
}

func DeleteSession(w http.ResponseWriter, r *http.Request, a *gobioweb.App) *gobioweb.AppError {
	return nil
}
