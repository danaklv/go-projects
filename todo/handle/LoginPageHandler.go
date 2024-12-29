package handle

import (
	"fmt"
	"html/template"
	"net/http"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("front/templates/login.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		err = tmpl.ExecuteTemplate(w,"login", "")
		if err != nil {
			fmt.Println(err)
			return
		}

	} 
	

}