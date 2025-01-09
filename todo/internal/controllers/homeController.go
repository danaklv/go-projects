package controllers

import (
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl.ExecuteTemplate(w, "base", nil)
		return
	}

}
