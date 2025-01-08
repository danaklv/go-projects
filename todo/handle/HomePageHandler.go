package handle

import (
	"html/template"
	"log"
	"net/http"
	"todo/models"
	"todo/session"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	email, ok := session.CheckSession(r)

	if !ok {
		log.Fatal("User not authenticated")
		return
	} else {
		tmpl, err := template.ParseFiles("front/templates/main.html")
		if err != nil {
			log.Fatal("Error with parsing files ", err)
		}

		data := models.User {
			Name: email,
		}

		err = tmpl.ExecuteTemplate(w, "main",data)

		if err != nil {
			log.Fatal("Error with parsing files ", err)
		}

	}

}
