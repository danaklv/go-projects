package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"todo/repositories"
	"todo/session"
)

func HomeController(w http.ResponseWriter, r *http.Request) {

	session, _ := session.Store.Get(r, "session-name")

	if email, ok := session.Values["email"].(string); ok {
		userId, err := repositories.GetUserIdFromDb(email)
		if err != nil {
			log.Fatal(err)
		}

		playlists, err := repositories.SelectPlaylistsFromBd(userId)

		for i := range playlists {
			playlists[i].ImagePath = strings.ReplaceAll(playlists[i].ImagePath, "\\", "/")
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(playlists[1].ImagePath)

		tmpl.ExecuteTemplate(w, "base", playlists)
		return

	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return

	}

}
