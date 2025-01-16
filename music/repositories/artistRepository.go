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

	 fmt.Println("artist img = ", artist.ImagePath)

	 artist.Id = artistId
	

	


	return artist, nil
}


func UpdateArtistImage(filePath string, artistId int) error {

	stmt :=`UPDATE artists SET image_path = $1 WHERE id = $2`
	_, err := DB.Exec(stmt, filePath, artistId)
	return err
}