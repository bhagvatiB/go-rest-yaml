package client

import (
	"payloadrest/src/dao/entity"

	"github.com/hashicorp/go-memdb"
)

type InMemDBClient struct {
	Db *memdb.MemDB
}

func NewInMemDBClient() InMemDBClient {
	return InMemDBClient{Db: getInMemDBClient()}
}

func getInMemDBClient() *memdb.MemDB {
	// Create a new data base
	db, err := memdb.NewMemDB(entity.GetAppMetadataSchema())
	if err != nil {
		panic(err)
	}
	return db
}
