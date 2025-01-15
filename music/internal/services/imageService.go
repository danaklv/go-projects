package services

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadImage(file multipart.File, header *multipart.FileHeader) (string, error) {
	
	title := header.Filename
	filePath := filepath.Join("uploads", title) 
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return filePath, err
}
