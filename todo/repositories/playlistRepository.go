package repositories

import (
	"fmt"
	"todo/models"
)

func InsertPlaylistInBd(playlist models.Playlist) error {

	stmt := `INSERT INTO playlists (title, user_id, image_path) VALUES ($1, $2, $3)`
	_, err := DB.Exec(stmt, playlist.Title, playlist.UserId, playlist.ImagePath)

	return err
}

func SelectPlaylistsFromBd(userId int) ([]models.Playlist, error) {
	query := `SELECT title, image_path FROM playlists WHERE user_id = $1`

	rows, err := DB.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error querying playlists: %v", err)
	}

	var playlists []models.Playlist

	for rows.Next() {
		var playlist models.Playlist

		err := rows.Scan(&playlist.Title, &playlist.ImagePath)

		if err != nil {
			return nil, fmt.Errorf("error scan: %v",err)
		}

		playlists = append(playlists, playlist)
	}

	return playlists, nil
}
