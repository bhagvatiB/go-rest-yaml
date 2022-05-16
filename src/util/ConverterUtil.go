package util

import (
	"payloadrest/src/model"
	"payloadrest/src/model/request"
	"payloadrest/src/model/response"
)

/**
* function to convert AppMetadata request object to AppMetadata dao object
**/
func ConvertAppMetadataReqToDao(AppMetadataReq request.AppMetadataRequest) model.AppMetadata {

	return model.AppMetadata{
		Id:          model.GenerateKey(AppMetadataReq.Title, AppMetadataReq.Version),
		Title:       AppMetadataReq.Title,
		Version:     AppMetadataReq.Version,
		Maintainers: AppMetadataReq.Maintainers,
		Company:     AppMetadataReq.Company,
		Website:     AppMetadataReq.Website,
		Source:      AppMetadataReq.Source,
		License:     AppMetadataReq.License,
		Description: AppMetadataReq.Description,
	}
}

/**
* function to convert AppMetadata dao object to AppMetadata response object
**/
func ConvertDaoToAppMetadataRes(AppMetadata model.AppMetadata) response.AppMetadataResponse {

	return response.AppMetadataResponse{
		Title:       AppMetadata.Title,
		Version:     AppMetadata.Version,
		Maintainers: AppMetadata.Maintainers,
		Company:     AppMetadata.Company,
		Website:     AppMetadata.Website,
		Source:      AppMetadata.Source,
		License:     AppMetadata.License,
		Description: AppMetadata.Description,
	}
}
