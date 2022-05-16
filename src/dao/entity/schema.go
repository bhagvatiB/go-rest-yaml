package entity

import (
	"github.com/hashicorp/go-memdb"
)

/**
* Method to create database table schema
* return schema
**/
func GetAppMetadataSchema() *memdb.DBSchema {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"APPDATA": &memdb.TableSchema{
				Name: "APPDATA",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
					"title": &memdb.IndexSchema{
						Name:    "title",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Title"},
					},
					"version": &memdb.IndexSchema{
						Name:    "version",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Version"},
					},
					"company": &memdb.IndexSchema{
						Name:    "company",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Company"},
					},
					"website": &memdb.IndexSchema{
						Name:    "website",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Website"},
					},
					"source": &memdb.IndexSchema{
						Name:    "source",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Source"},
					},
					"license": &memdb.IndexSchema{
						Name:    "license",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "License"},
					},
					"description": &memdb.IndexSchema{
						Name:    "description",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Description"},
					},
					"maintainers": &memdb.IndexSchema{
						Name:    "maintainers",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Maintainers"},
					},
				},
			},
			/* "MAINTAINER": &memdb.TableSchema{
				Name: "MAINTAINER",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringSliceFieldIndex{Field: "Name"},
					},
					"email": &memdb.IndexSchema{
						Name:    "email",
						Unique:  false,
						Indexer: &memdb.StringSliceFieldIndex{Field: "Email"},
					},
				},
			}, */
		},
	}
	return schema
}
