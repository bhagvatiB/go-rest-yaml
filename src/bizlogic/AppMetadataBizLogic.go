package bizlogic

import (
	"net/url"
	"payloadrest/src/dao"
	"payloadrest/src/model"
	"payloadrest/src/model/request"
	"payloadrest/src/util"
)

type AppMetadataBizLogic struct {
	getMetadataDao dao.InMemAppMetadataDao
}

/**
* initialize AppMetadataBizLogic
**/
func NewAppMetadataBizLogic(appMetadataDao dao.InMemAppMetadataDao) AppMetadataBizLogic {
	return AppMetadataBizLogic{getMetadataDao: appMetadataDao}
}

/**
* logic for create AppMetadata
* Input: -- AppMetadata request Object
* convert request to AppMetadata dao
**/
func (createNewMetadataLogic *AppMetadataBizLogic) CreateAppMetadataLogic(AppMetadataReq request.AppMetadataRequest) error {

	// convert AppMetaData Request to AppMetadata Dao
	AppMetadata := util.ConvertAppMetadataReqToDao(AppMetadataReq)
	return createNewMetadataLogic.getMetadataDao.CreateAppMetadataDao(AppMetadata)
}

/**
* logic for search AppMetadata
* Input: -- AppMetadata request parameters
* return: -- List of all AppMetaData matched with query parameters
**/
func (createNewMetadataLogic *AppMetadataBizLogic) SearchAppMetadataLogic(values url.Values) ([]model.AppMetadata, error) {
	return createNewMetadataLogic.getMetadataDao.SearchAppMetadataDao(values)
}
