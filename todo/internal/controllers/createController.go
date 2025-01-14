package controllers

import (
	"log"
	"net/http"
	"todo/internal/services"
	"todo/models"
	"todo/repositories"
	"todo/session"
)

func CreateController(w http.ResponseWriter, r *http.Request) {
	session, _ := session.Store.Get(r, "session-name")

	if email, ok := session.Values["email"].(string); ok {

		title := r.FormValue("title")
		
		userId, err := repositories.GetUserIdFromDb(email)
		if err != nil {
			log.Fatal(err)
		}

		file, header, err := r.FormFile("image")

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		filePath, err := services.UploadImage(file, header)
		if err != nil {
			log.Fatal(err)
		}

		playlist := models.Playlist{
			Title:     title,
			UserId:    userId,
			ImagePath: filePath,
		}

		err = repositories.InsertPlaylistInBd(playlist)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}
