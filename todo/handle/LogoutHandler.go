package handle

import (
	"fmt"
	"html/template"
	"net/http"
	"todo/session"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session.DeleteSession(w, r)

	t, err := template.ParseFiles("front/templates/index.html", "front/styles/style.css")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.ExecuteTemplate(w, "index", "")
	if err != nil {
		fmt.Println(err)
		return
	}

}
