package handle

import (
	"fmt"
	"html/template"
	"net/http"
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
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		fmt.Println(name, email, password)
	}

}
