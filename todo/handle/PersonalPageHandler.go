package handle

import (
	"log"
	"net/http"
	"text/template"
)

func PersonalPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("front/templates/main")

		if err != nil {
			log.Fatal("Error parse front/template/main - ", err)
			return
		}

		err = tmpl.ExecuteTemplate(w, "main", "")

		if err != nil {
			log.Fatal("Error execute template main - ", err)
			return
		}
	}

}
