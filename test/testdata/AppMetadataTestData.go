package testdata

import (
	"payloadrest/src/bizlogic"
	"payloadrest/src/client"
	"payloadrest/src/controller"
	"payloadrest/src/dao"
	"payloadrest/src/model"
	"payloadrest/src/model/request"
	"strconv"
	"sync"
	"testing"

	"github.com/go-playground/assert/v2"
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

func TestCreateAppMetadataDaoMultiThreadingSameData(t *testing.T) {
	var wg sync.WaitGroup
	expectedAppMetadata := testdata.GetAppMetadata("3")

	for i := 1; i <= 20000; i++ {
		wg.Add(1)
		go func(expectedAppMetadata model.AppMetadata) {

			defer wg.Done()

			inMemAppMetadataDao.CreateAppMetadataDao(expectedAppMetadata)
		}(expectedAppMetadata)
	}

	wg.Wait()

	actualAppMetadataList, _ := inMemAppMetadataDao.SearchAppMetadataDao(nil)

	assert.Equal(t, len(actualAppMetadataList), 3) // two are created in above test
}

func TestCreateAppMetadataDaoMultiThreadingDifferentData(t *testing.T) {

	var wg sync.WaitGroup

	for i := 1; i <= 20000; i++ {
		expectedAppMetadata := testdata.GetAppMetadata(strconv.Itoa(i))

		wg.Add(1)

		go func(expectedAppMetadata model.AppMetadata) {

			defer wg.Done()

			inMemAppMetadataDao.CreateAppMetadataDao(expectedAppMetadata)
		}(expectedAppMetadata)
	}

	wg.Wait()

	actualAppMetadataList, _ := inMemAppMetadataDao.SearchAppMetadataDao(nil)

	assert.Equal(t, len(actualAppMetadataList), 20000) // 3 are created in above tests
}
