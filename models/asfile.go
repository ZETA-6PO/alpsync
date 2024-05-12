package models

import (
	"github.com/kamva/mgm/v3"
)

type ASFile struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	ExpireAt         string `json:"expireAt" bson:"expireAt"`
}

func NewASFile(name string, expireAt string) *ASFile {
	return &ASFile{
		Name:     name,
		ExpireAt: expireAt,
	}
}
