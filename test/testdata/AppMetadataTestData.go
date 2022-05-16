package testdata

import (
	"payloadrest/src/bizlogic"
	"payloadrest/src/client"
	"payloadrest/src/controller"
	"payloadrest/src/dao"
	"payloadrest/src/model"
	"payloadrest/src/model/request"
)

func GetAppMetadata(i string) model.AppMetadata {
	Maintainer := model.Maintainer{
		Name:  "Bhagvati" + i,
		Email: i + "b.b@gmail.com",
	}

	var MaintainerList []model.Maintainer
	MaintainerList = append(MaintainerList, Maintainer)

	return model.AppMetadata{
		Id:          "Valid App " + i + "_0.0." + i,
		Title:       "Valid App " + i,
		Version:     "0.0." + i,
		Maintainers: MaintainerList,
		Company:     "microsoft+i",
		Website:     "http://" + i + "microsoft.com",
		Source:      "https://github.com/" + i + "upbound/repo",
		License:     i + "Apache-2.0",
		Description: i + "How are you",
	}
}

func GetAppMetadataReq(i string) request.AppMetadataRequest {
	Maintainer := model.Maintainer{
		Name:  "Bhagvati" + i,
		Email: i + "b.b@gmail.com",
	}

	var MaintainerList []model.Maintainer
	MaintainerList = append(MaintainerList, Maintainer)

	return request.AppMetadataRequest{
		Title:       "Valid App " + i,
		Version:     "0.0." + i,
		Maintainers: MaintainerList,
		Company:     "microsoft+i",
		Website:     "http://" + i + "microsoft.com",
		Source:      "https://github.com/" + i + "upbound/repo",
		License:     i + "Apache-2.0",
		Description: i + "How are you",
	}
}

func GetEmptyAppMetadataReq() request.AppMetadataRequest {
	return request.AppMetadataRequest{}
}

func GetInvalidEmailAppMetadataReq() request.AppMetadataRequest {
	Maintainer := model.Maintainer{
		Name:  "Bhagvati",
		Email: "b.b@.com",
	}

	var MaintainerList []model.Maintainer
	MaintainerList = append(MaintainerList, Maintainer)

	return request.AppMetadataRequest{
		Title:       "Valid App 1",
		Version:     "0.0.1",
		Maintainers: MaintainerList,
		Company:     "microsoft",
		Website:     "http://microsoft.com",
		Source:      "https://github.com/upbound/repo",
		License:     "Apache-2.0",
		Description: "How are you",
	}
}

func GetInvalidUrlAppMetadataReq() request.AppMetadataRequest {
	Maintainer := model.Maintainer{
		Name:  "Bhagvati",
		Email: "b.b@gmail.com",
	}

	var MaintainerList []model.Maintainer
	MaintainerList = append(MaintainerList, Maintainer)

	return request.AppMetadataRequest{
		Title:       "Valid App 1",
		Version:     "0.0.1",
		Maintainers: MaintainerList,
		Company:     "microsoft",
		Website:     "123",
		Source:      "https://github.com/upbound/repo",
		License:     "Apache-2.0",
		Description: "How are you",
	}
}

func InitializeServer() controller.AppMetadataController {
	inMemDBClient := client.NewInMemDBClient()
	inMemAppMetadataDao := dao.NewInMemAppMetadataDao(inMemDBClient)
	appMetadataBizLogic := bizlogic.NewAppMetadataBizLogic(inMemAppMetadataDao)
	appMetadataController := controller.NewAppMetadataController(appMetadataBizLogic)
	return appMetadataController
}
