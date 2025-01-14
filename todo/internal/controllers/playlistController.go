package controllers

import (
	"log"
	"net/http"
	"strconv"
	"todo/models"
	"todo/repositories"
)

func PlaylistController(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	playlistId, err := strconv.Atoi(path[len("/playlist/"):])
	if err != nil {
		log.Fatal(err)
	}

	if playlistId == 0 {
		log.Fatal("playlistId nil")
		return
	}

	if r.Method == http.MethodGet {

		songs, err := repositories.GetSongsFromDb(playlistId)

		if err != nil {
			log.Fatal("error with get songs from db : ", err)
		}

		songs, err = repositories.UpdateSongsWithArtistNames(songs)
		if err != nil {
			log.Fatal("error with get artist name by id : ", err)
		}

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

		if err != nil {

			log.Fatal(err)

		}
		artistId := repositories.GetArtistId(artistName)
		artist := models.Artist{
			Id:   artistId,
			Name: artistName,
		}

		song := models.Song{
			Title:  title,
			Artist: artist,
		}

		songId, err := repositories.InsertSongInBd(song)
		if err != nil {
			log.Fatal(err)
		}

		err = repositories.InsertSongInPlaylist(songId, playlistId)
		if err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, r.URL.Path, http.StatusSeeOther)

	}

}
