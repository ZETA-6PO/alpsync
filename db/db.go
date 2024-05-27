package db

import (
	"alpsync-api/models"
	"errors"
	"time"

	"github.com/kamva/mgm/v3"
)

// Add file-entry to the database and return the id of the file. Return null,err in case of an error.
// ExpiresIn 0: 1d; 1: 14d
func AddFileEntry(name string, expiresIn int) (string, error) {
	expiryDate := time.Now().AddDate(0, 0, expiresIn)
	//creer une entree dans la base de donne
	dbentry := models.NewASFile(name, expiryDate.Format(time.UnixDate))
	coll := mgm.CollectionByName("files")
	err := coll.Create(dbentry)
	if err != nil {
		return "", err
	}

	return dbentry.ID.Hex(), nil
}

// Retrieve the file name of the provided file id.
func GetFileEntry(id string) (string, error) {
	coll := mgm.CollectionByName("files")
	dbentry := &models.ASFile{}
	err := coll.FindByID(id, dbentry)
	if err != nil {
		return "", err
	}
	expiresAt, err := time.Parse(time.UnixDate, dbentry.ExpiresAt)
	if err != nil {
		return "", err
	}
	if time.Now().After(expiresAt) {
		return "", errors.New("This link has expired.")
	}
	return dbentry.Name, nil

}
