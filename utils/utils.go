package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func CreateFile(file multipart.File, file_path string) error {
	// Crée un fichier local pour stocker le fichier téléchargé
	f, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	// Copie le contenu du fichier téléchargé dans le fichier local
	_, err = io.Copy(f, file)
	if err != nil {
		return err
	}

	return nil

}

func ReadFile(path string) (os.FileInfo, *os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	fileStat, err := file.Stat()
	if err != nil {
		return nil, nil, err
	}

	return fileStat, file, nil
}
