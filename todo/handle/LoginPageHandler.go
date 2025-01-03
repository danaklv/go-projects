package handle

import (
	"fmt"
	"html/template"
	"net/http"
	"todo/database"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {

	 if r.Method == http.MethodPost {

		email := r.FormValue("email")
		password := r.FormValue("password")

		if database.ChechUserInDb(email, password) {

			tmpl, err := template.ParseFiles("front/templates/index.html")
			if err != nil {
				fmt.Println(err)
				return
			}

			err = tmpl.ExecuteTemplate(w, "index", "")
			if err != nil {
				fmt.Println(err)
				return
			}

		} else {
			fmt.Println("Incorrect email or password")

		}

	} else {
		w.WriteHeader(http.StatusLoopDetected)
		fmt.Println("Incorecct method")

	}

}
