package dao

import (
	"fmt"
	"net/url"
	"payloadrest/src/client"
	"payloadrest/src/logger"
	"payloadrest/src/model"
	"payloadrest/src/util"
)

type void struct{}

var member void

type InMemAppMetadataDao struct {
	dbClient client.InMemDBClient
}

/**
* initialize InMemAppMetadataDao
**/
func NewInMemAppMetadataDao(dbClient client.InMemDBClient) InMemAppMetadataDao {
	return InMemAppMetadataDao{dbClient: dbClient}
}

/**
* function to search AppMetadata from database based on search query parameters
* Input: -- Query parameters
* return: -- AppMetadata dao object
**/
func (query *InMemAppMetadataDao) SearchAppMetadataDao(values url.Values) ([]model.AppMetadata, error) {
	logger.SugarLogger.Infof("Dao: SearchAppMetadataDao")

	db := query.dbClient.Db

	var PayloadList []model.AppMetadata

	// if query parameters are empty, return all AppMetadata list
	if len(values) == 0 {
		return query.GetAllAppMetadataDao()
	}

	// create set to store unique AppMetadata list
	set := make(map[string]void)

	// iterate query parametres, read all records of all app data
	for k, v := range values {

		// check if all key parameters are valid
		// if key is not valid do not query in memory database
		if util.CheckKeysPresent(k) {
			txn := db.Txn(false)
			it, err := txn.Get("APPDATA", k+"_prefix", v[0])

			if err != nil {
				logger.SugarLogger.Infof("Error getting data from in memory database %w", err)
				return []model.AppMetadata{}, fmt.Errorf("failed to get data from in memory database")
			}

			for obj := it.Next(); obj != nil; obj = it.Next() {
				p := obj.(model.AppMetadata)

				_, exists := set[p.Id]

				if !exists {
					set[p.Id] = member
					PayloadList = append(PayloadList, p)
				}
			}
		}
	}

	return PayloadList, nil
}

/**
* function to insert AppMetadata in to in memory database
* Input: --AppMetadata dao object
**/
func (query *InMemAppMetadataDao) CreateAppMetadataDao(Payload model.AppMetadata) error {
	logger.SugarLogger.Infof("Dao: CreateNewMetaDataDao")
	// create a write transaction
	db := query.dbClient.Db
	txn := db.Txn(true)

	// insert AppMetadata into in memory database
	if err := txn.Insert("APPDATA", Payload); err != nil {
		logger.SugarLogger.Infof("Error inserting into database %w", err)
		return fmt.Errorf("failed to insert data in to in-memory database")
	}

	// Commit the transaction
	txn.Commit()

	return nil
}

/**
* function to get all AppMetadata list from database
* return: -- List of all AppMetadata Dao object
**/
func (query *InMemAppMetadataDao) GetAllAppMetadataDao() ([]model.AppMetadata, error) {
	logger.SugarLogger.Infof("Dao: GetAllAppMetadataDao")

	var AppMetaDataList []model.AppMetadata

	// create a read transaction
	db := query.dbClient.Db
	txn := db.Txn(false)

	// get list all AppMetadata from database
	it, err := txn.Get("APPDATA", "id")
	if err != nil {
		return []model.AppMetadata{}, fmt.Errorf("failed to get data from in memory database")
	}

	// iterate all list and put it in to AppMetadataList variable
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(model.AppMetadata)
		AppMetaDataList = append(AppMetaDataList, p)
	}

	return AppMetaDataList, nil
}
