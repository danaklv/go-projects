package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"todo/models"
	"todo/repositories"
	"todo/session"
	"todo/utils"
)

var tmpl = template.Must(template.ParseGlob("front/templates/*.html"))

func LoginController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl.ExecuteTemplate(w, "login", nil)
		return
	}

	if r.Method == http.MethodPost {

		email := r.FormValue("email")
		password := r.FormValue("password")

		if repositories.CheckUserInDb(email, password) {
			session.CreateNewSession(w, r, email)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return

		} else {
			fmt.Println("Incorrect email or password")

		}

	} else {
		fmt.Println("Incorecct method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}

}

func RegisterController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl.ExecuteTemplate(w, "register", nil)
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		user := models.User{
			Name:     name,
			Email:    email,
			Password: password,
		}

		if utils.UserValidation(user) {
			repositories.InsertUserIntoDb(&user)
			http.Redirect(w,r,"/login", http.StatusSeeOther)
		} else {
			fmt.Println("User with this email arleady exist")
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Incorecct method")
	}

}
