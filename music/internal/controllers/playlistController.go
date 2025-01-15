package controllers

import (
	"log"
	"net/http"
	"strconv"
	"todo/internal/services"
	"todo/models"
)

func PlaylistController(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	playlistId, err := strconv.Atoi(path[len("/playlist/"):])
	if err != nil {
		log.Fatal(err)
	}
	
	if r.Method == http.MethodGet {

		songs := services.GetSongsFromPlaylist(playlistId)

		data := models.PlaylistSongs{
			PlaylistId: playlistId,
			Songs:      songs,
		}
		err = tmpl.ExecuteTemplate(w, "playlist", data)
		if err != nil {
			log.Fatal(err)
		}

	} else if r.Method == http.MethodPost {
		title := r.FormValue("title")
		artistName := r.FormValue("artist")

		services.AddNewSong(title, artistName, playlistId)

		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)

	}

}
