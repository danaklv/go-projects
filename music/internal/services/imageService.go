package services

import (
	"io"
	"mime/multipart"
	"os"
)

func UploadImage(file multipart.File, header *multipart.FileHeader) (string, error) {

	title := header.Filename

	outFile, err := os.Create("uploads/" + title)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return title, err
}
