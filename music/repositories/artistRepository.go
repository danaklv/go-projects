package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"todo/models"
)

func GetArtistFromDb(artistId int) (models.Artist, error) {
	var artist models.Artist

	query := `SELECT name, image_path FROM artists WHERE id = $1`

	err := DB.QueryRow(query, artistId).Scan(&artist.Name, &artist.ImagePath)
	if err == sql.ErrNoRows {
		fmt.Println(err)
		return artist, errors.New("no such artist")
	}

	artist.Id = artistId

	return artist, nil
}

func UpdateArtistImage(filePath string, artistId int) error {

	stmt := `UPDATE artists SET image_path = $1 WHERE id = $2`
	_, err := DB.Exec(stmt, filePath, artistId)
	return err
}

func GetArtistsBySongs(songs []models.Song) ([]models.Artist, error) {

	var artists []models.Artist
	for i := range songs {
		artist, err := GetArtistFromDb(songs[i].Artist.Id)
		if err != nil {
			return nil, err
		}
		if !Contains(artists, artist) {
			artists = append(artists, artist)
		}
	}

	return artists, nil

}

func Contains(artists []models.Artist, artist models.Artist) bool {
	for i := range artists {
		if artists[i].Id == artist.Id {
			return true
		}
	}

	return false
}
