package controllers

import (
	"fmt"
	"html/template"
	"log"
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

		if repositories.ChechUserInDb(email, password) {
			session.CreateNewSession(w, r, email)
			http.Redirect(w, r, "/home", http.StatusSeeOther)
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
		id := 0
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		user := models.User{id, name, email, password}

		if utils.UserValidation(user) {
			repositories.InsertUserIntoDb(&user)
		} else {
			tmpl, err := template.ParseFiles("front/templates/index.html")
			if err != nil {
				log.Fatal("Error with parsing files ", err)
				return
			}

			tmpl.Execute(w, map[string]string{
				"Error": "User with this email arleady exist",
			})

			if err != nil {
				log.Fatal("Error with parsing files ", err)
				return
			}
			fmt.Println("User with this email arleady exist")
		}

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Incorecct method")
	}

}
