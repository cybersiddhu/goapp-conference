package controller

import (
	"github.com/cybersiddhu/gobioweb"
	"net/http"
)

func Welcome(c *gobioweb.Controller) *gobioweb.AppError {
	t := c.App.Template
	c.Stash("message", "Welcome to GoSidd framework")
	err := t.Process("welcome").Execute(c.Response, c)
	if err != nil {
		return &gobioweb.AppError{Error: err, Code: 500, Message: err.Error()}
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

	err := c.App.Template.Process("signup").Execute(c.Response, c)
	if err != nil {
		return &gobioweb.AppError{
			Error:   err,
			Code:    500,
			Message: "Cannot display template signup.tmpl",
		}
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
				Error:   err,
				Code:    500,
				Message: err.Error(),
			}
		}
		http.Redirect(w, r, "/users/new", http.StatusFound)
	}

	//now the password
	pass := r.FormValue("pass")
	if len(pass) == 0 {
		err := c.SetFormErrors("Password not provided")
		if err != nil {
			return &gobioweb.AppError{
				Error:   err,
				Code:    500,
				Message: err.Error(),
			}
		}
		http.Redirect(w, r, "/users/new", http.StatusFound)
	}

	u := &gobioweb.User{
		Email:     email,
		Password:  pass,
		FirstName: r.FormValue("fname"),
		LastName:  r.FormValue("lname"),
	}

	if err := u.Create(c.App.Database); err != nil {
		err := c.SetFormErrors(err.Error())
		if err != nil {
			return &gobioweb.AppError{
				Error:   err,
				Code:    500,
				Message: err.Error(),
			}
		}
		http.Redirect(w, r, "/users/new", http.StatusFound)
	}

	if err := c.SaveUserSession(u); err != nil {
		return &gobioweb.AppError{
			Error:   err,
			Code:    500,
			Message: err.Error(),
		}
	}
	http.Redirect(w, r, "/", http.StatusFound)
	return nil
}

func Login(c *gobioweb.Controller) *gobioweb.AppError {
	err := c.App.Template.Process("login").Execute(c.Response, c)
	if err != nil {
		return &gobioweb.AppError{
			Error:   err,
			Code:    500,
			Message: err.Error(),
		}
	}
	return nil
}

func CreateSession(c *gobioweb.Controller) *gobioweb.AppError {
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
				Error:   err,
				Code:    500,
				Message: err.Error(),
			}
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	//now the password
	pass := r.FormValue("pass")
	if len(pass) == 0 {
		err := c.SetFormErrors("Password not provided")
		if err != nil {
			return &gobioweb.AppError{
				Error:   err,
				Code:    500,
				Message: err.Error(),
			}
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	u := &gobioweb.User{Email: email, Password: pass}
	if err := u.Validate(c.App.Database); err != nil {
		 c.SetFormErrors("Could not login: Wrong email or password")
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		if err := c.SaveUserSession(u); err != nil {
			return &gobioweb.AppError{
				Error:   err,
				Code:    500,
				Message: err.Error(),
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
	return nil
}

func DeleteSession(c *gobioweb.Controller) *gobioweb.AppError {
	if err := c.DeleteUserSession(); err != nil {
		c.SetFormErrors("Could not remove user sessions")
	}
	http.Redirect(c.Response, c.Request, "/", http.StatusFound)
	return nil
}
