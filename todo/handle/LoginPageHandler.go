package handle

import (
	"fmt"
	"net/http"
	"todo/database"
	"todo/session"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		email := r.FormValue("email")
		password := r.FormValue("password")

		if database.ChechUserInDb(email, password) {
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
