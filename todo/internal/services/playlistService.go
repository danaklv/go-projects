package services

import (
	"log"
	"todo/models"
	"todo/repositories"
)

func GetSongsFromPlaylist(playlistId int) []models.Song {

	songs, err := repositories.GetSongsFromDb(playlistId)
	if err != nil {
		log.Fatal("error with get songs from db : ", err)
	}

	songs, err = repositories.UpdateSongsWithArtistNames(songs)
	if err != nil {
		log.Fatal("error with get artist name by id : ", err)
	}

	return songs

}

func AddNewSong(title, artistName string, playlistId int) {

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

}
