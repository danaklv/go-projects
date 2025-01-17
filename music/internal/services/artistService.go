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

func GetArtistsByPlaylists(playlists []models.Playlist) []models.Artist {

	var res []models.Artist
	for i := range playlists {
		songs, err := repositories.GetSongsFromDb(playlists[i].Id)
		if err != nil {
			log.Fatal(err)
		}
		artists, err := repositories.GetArtistsBySongs(songs)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, artists...)
	}

	return res

}
