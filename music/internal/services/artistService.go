package services

import (
	"log"
	"todo/models"
	"todo/repositories"
)

func GetArtistById(artistId int) models.Artist {
	artist, err := repositories.GetArtistFromDb(artistId)

	if err != nil {
		log.Fatal(err)
	}

	artistSongs, err := repositories.GetSongsByArtistId(artistId)	
	if err != nil {
		log.Fatal(err)
	}

	artist.Songs = artistSongs
	return artist
}