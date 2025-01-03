package handle

import (
	"fmt"
	"html/template"
	"net/http"
	"todo/check"
	"todo/database"
	"todo/models"
)

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("front/templates/register.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = tmpl.ExecuteTemplate(w, "register", "")
		if err != nil {
			fmt.Println(err)
			return
		}

	} else if r.Method == http.MethodPost {
		id := 0
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		user := models.User {id, name, email, password}

		if check.UserValidation(user) {
			database.InsertUserIntoDb(&user)
		} else {
			fmt.Println("User with this email arleady exist")
		}

		
		
	}

}
