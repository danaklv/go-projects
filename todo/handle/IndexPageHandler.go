package handle

import (
	"fmt"
	"html/template"
	"net/http"
)

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {

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
