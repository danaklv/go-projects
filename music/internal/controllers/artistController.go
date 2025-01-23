package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"todo/internal/services"
	"todo/repositories"
)


//Artist Page Handler
func ArtistController(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if !strings.Contains(path, "/uploads") {

		artistId, err := strconv.Atoi(path[len("/artist/"):])
		if err != nil {
			log.Fatal(err)
		}

		if r.Method == http.MethodGet {
			artist := services.GetArtistById(artistId)

			tmpl.ExecuteTemplate(w, "artist", artist)

		}
	}

}

//Upliad Artist Image
func UploadArtistImageController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		artistId, _ := strconv.Atoi(r.FormValue("artistId"))

		file, header, err := r.FormFile("image")

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		filePath, err := services.UploadImage(file, header)

		
	

		if err != nil {
			log.Fatal(err)
		}

		err = repositories.UpdateArtistImage(filePath, artistId)

		if err != nil {
			log.Fatal(err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}
