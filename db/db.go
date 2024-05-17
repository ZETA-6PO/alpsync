package db

import (
	"alpsync-api/models"
	"github.com/kamva/mgm/v3"
	"time"
)

// Add file-entry to the database and return the id of the file. Return null,err in case of an error.
func AddFileEntry(name string, expireAt string) (string, error) {
	//creer une entree dans la base de donne
	dbentry := models.NewASFile(name, time.Now().Format(time.UnixDate))
	coll := mgm.CollectionByName("files")
	err := coll.Create(dbentry)
	if err != nil {
		return "", err
	}

	return dbentry.ID.Hex(), nil
}
