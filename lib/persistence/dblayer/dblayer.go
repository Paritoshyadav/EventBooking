package dblayer

import (
	"github.com/eventbooking/lib/persistence"
	"github.com/eventbooking/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB  DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongolayer.InitMongoLayer(connection)
	}
	return nil, nil

}
