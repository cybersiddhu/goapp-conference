package controller

import (
	"github.com/cybersiddhu/gobioweb"
	"net/http"
)

func Welcome(c *gobioweb.Controller) *gobioweb.AppError {
	t := c.App.Template
	c.Stash("message", "Welcome to GoSidd framework")
	err := t.ExecuteTemplate(c.Response, "welcome.tmpl", c)
	if err != nil {
		return &gobioweb.AppError{Error: err, Code: 500, Message: "Cannot display template"}
	}
	return nil
}

func ListAbstract(c *gobioweb.Controller) *gobioweb.AppError {
	return nil
}

func CreateAbstract(c *gobioweb.Controller) *gobioweb.AppError {
	return nil
}

func NewAbstract(c *gobioweb.Controller) *gobioweb.AppError {
	return nil
}

func ShowAbstract(c *gobioweb.Controller) *gobioweb.AppError {
	return nil
}

func NewUser(c *gobioweb.Controller) *gobioweb.AppError {
	f := map[string]string{"fname": "First Name", "lname": "Last Name"}
	c.Stash("names", f)
	err := c.App.Template.ExecuteTemplate(c.Response, "signup.tmpl", c)
	if err != nil {
		return &gobioweb.AppError{Error: err, Code: 500, Message: "Cannot display template signup.tmpl"}
	}
	return nil
}

func CreateUser(c *gobioweb.Controller) *gobioweb.AppError {
	r := c.Request
	w := c.Response
	err := r.ParseForm()
	if err != nil {
		return &gobioweb.AppError{Error: err, Code: 500, Message: err.Error()}
	}

	//form data processing
	//email input
	email := r.FormValue("email")
	if len(email) == 0 {
		err := c.SetFormErrors("Email not provided")
		if err != nil {
			 return &gobioweb.AppError{
					Error: err,
					Code: 500,
					Message: err.Error(),
			 }
		}
		http.Redirect(w, r, "/users/new", http.StatusSeeOther)
	}

	//now the password
	pass := r.FormValue("pass")
	if len(pass) == 0 {
		err := c.SetFormErrors("Password not provided")
		if err != nil {
			 return &gobioweb.AppError{
					Error: err,
					Code: 500,
					Message: err.Error(),
			 }
		}
		http.Redirect(w, r, "/users/new", http.StatusFound)
	}
	return nil
}

func Login(c *gobioweb.Controller) *gobioweb.AppError {
	return nil
}

func CreateSession(c *gobioweb.Controller) *gobioweb.AppError {
	return nil
}

func DeleteSession(c *gobioweb.Controller) *gobioweb.AppError {
	return nil
}
