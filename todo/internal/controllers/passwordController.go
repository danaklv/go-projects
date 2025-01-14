package controllers

import (
	"net/http"
)

func ResetPasswordController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl.ExecuteTemplate(w, "resetPassword", nil)
	} else if r.Method == http.MethodPost {
		// email := r.FormValue("email")
		// err := services.RequestPasswordReset(email)
		// if err != nil {

		// }
	}

}
