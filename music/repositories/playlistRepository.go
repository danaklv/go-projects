package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"todo/models"
)

func InsertPlaylistInBd(playlist models.Playlist) error {

	stmt := `INSERT INTO playlists (title, user_id, image_path) VALUES ($1, $2, $3)`
	_, err := DB.Exec(stmt, playlist.Title, playlist.UserId, playlist.ImagePath)

	return err
}

func SelectPlaylistsFromBd(userId int) ([]models.Playlist, error) {
	query := `SELECT id, title, image_path FROM playlists WHERE user_id = $1`

	rows, err := DB.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error querying playlists: %v", err)
	}

	var playlists []models.Playlist

	for rows.Next() {
		var playlist models.Playlist

		err := rows.Scan(&playlist.Id, &playlist.Title, &playlist.ImagePath)

		if err != nil {
			return nil, fmt.Errorf("error scan: %v", err)
		}

		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func GetSongsFromDb(playlistId int) ([]models.Song, error) {
	songId := 0
	query := `SELECT song_id FROM playlist_songs WHERE playlist_id = $1`
	rows, err := DB.Query(query, playlistId)
	if err != nil {
		return nil, fmt.Errorf("error querying playlist: %v", err)
	}
	var songs []models.Song

	for rows.Next() {
		err = rows.Scan(&songId)
		if err != nil {
			return nil, fmt.Errorf("error scan1: %v", err)
		}
		var song models.Song
		query = `SELECT title, create_date, artist_id FROM songs WHERE id = $1`
		err := DB.QueryRow(query, songId).Scan(&song.Title, &song.CreateDate, &song.Artist.Id)
		if err != nil {
			return nil, fmt.Errorf("error querying playlist: %v", err)
		}
		songs = append(songs, song)

	}

	return songs, nil
}


func GetSongsByArtistId(artistId int) ([]models.Song, error) {

	query := `SELECT title FROM songs WHERE artist_id = $1`
	rows, err := DB.Query(query, artistId)
	if err != nil {
		return nil, fmt.Errorf("error querying songs: %v", err)
	}
	var songs []models.Song

	for rows.Next() {
		var song models.Song
		rows.Scan(&song.Title)

		songs = append(songs, song)

	}

	return songs, nil
}


func GetArtistId(name string) int {
	id := 0
	query := `SELECT id FROM artists WHERE name = $1`

	err := DB.QueryRow(query, name).Scan(&id)

	if err == sql.ErrNoRows {
		stmt := `INSERT INTO artists (name) VALUES ($1)`
		_, err = DB.Exec(stmt, name)
		if err != nil {
			log.Fatal(err)
		}
		query := `SELECT id FROM artists WHERE name = $1`

		err = DB.QueryRow(query, name).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

	} else if err != nil {
		log.Fatal(err)
	}

	return id

}

func UpdateSongsWithArtistNames(songs []models.Song) ([]models.Song, error) {
	for i := range songs { 
		query := `SELECT name FROM artists WHERE id = $1`
		var artistName string

		
		err := DB.QueryRow(query, songs[i].Artist.Id).Scan(&artistName)
		if err != nil {
			return nil, fmt.Errorf("error querying artist name for artist_id %d: %v", songs[i].Artist.Id, err)
		}

	
		songs[i].Artist.Name = artistName
	}

	return songs, nil
}


func InsertSongInBd(song models.Song) (int, error) {
	songId := 0
	check := `SELECT id FROM songs WHERE title = $1`
	err := DB.QueryRow(check, song.Title).Scan(&songId)
	if err == sql.ErrNoRows {
		query := `INSERT INTO songs (artist_id, title) VALUES ($1, $2) RETURNING id`
		err = DB.QueryRow(query, song.Artist.Id, song.Title).Scan(&songId)
	}

	return songId, err
}

func InsertSongInPlaylist(songId, playlistId int) error {

	stmt := `INSERT INTO playlist_songs (playlist_id, song_id) VALUES ($1, $2)`

	_, err := DB.Exec(stmt, playlistId, songId)
	return err
}
