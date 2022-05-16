package dao_test

import (
	"fmt"
	"payloadrest/src/client"
	"payloadrest/src/dao"
	"payloadrest/src/logger"
	"payloadrest/src/model"
	"payloadrest/test/testdata"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var inMemAppMetadataDao dao.InMemAppMetadataDao

func SetUp() {
	logger.InitLogger()
	inMemDBClient := client.NewInMemDBClient()
	inMemAppMetadataDao = dao.NewInMemAppMetadataDao(inMemDBClient)
}

func TestCreateAppMetadataDao(t *testing.T) {
	SetUp()
	expectedAppMetadata1 := testdata.GetAppMetadata("1")
	inMemAppMetadataDao.CreateAppMetadataDao(expectedAppMetadata1)

	expectedAppMetadata2 := testdata.GetAppMetadata("2")
	inMemAppMetadataDao.CreateAppMetadataDao(expectedAppMetadata2)
}

func TestSearchAppMetadataDao(t *testing.T) {

	var expectedAppMetadataList []model.AppMetadata
	expectedAppMetadata := testdata.GetAppMetadata("1")
	expectedAppMetadataList = append(expectedAppMetadataList, expectedAppMetadata)
	expectedAppMetadata2 := testdata.GetAppMetadata("2")
	expectedAppMetadataList = append(expectedAppMetadataList, expectedAppMetadata2)

	actualAppMetadataList, _ := inMemAppMetadataDao.SearchAppMetadataDao(nil)
	assert.True(t, reflect.DeepEqual(expectedAppMetadataList, actualAppMetadataList))
	assert.Equal(t, len(actualAppMetadataList), 2)
}

func TestSearchAppMetadataDaoWithQueryParameters(t *testing.T) {

	//Create query params
	queryParams := make(map[string][]string)
	queryParams["title"] = []string{"Valid App 1"}

	var expectedAppMetadataList []model.AppMetadata
	expectedAppMetadata := testdata.GetAppMetadata("1")
	expectedAppMetadataList = append(expectedAppMetadataList, expectedAppMetadata)

	actualAppMetadataList, _ := inMemAppMetadataDao.SearchAppMetadataDao(queryParams)
	fmt.Println("expect 2: ", expectedAppMetadataList)
	fmt.Println("actual 2: ", actualAppMetadataList)
	assert.True(t, reflect.DeepEqual(expectedAppMetadataList, actualAppMetadataList))

	queryParams["title"] = []string{"Valid App 3"}
	actualAppMetadataList, _ = inMemAppMetadataDao.SearchAppMetadataDao(queryParams)
	fmt.Println("expect 3: ", expectedAppMetadataList)
	fmt.Println("actual 3: ", actualAppMetadataList)
	assert.Nil(t, actualAppMetadataList)

	queryParams["version"] = []string{"0.0.1"}
	actualAppMetadataList, _ = inMemAppMetadataDao.SearchAppMetadataDao(queryParams)
	fmt.Println("expect 4: ", expectedAppMetadataList)
	fmt.Println("actual 4: ", actualAppMetadataList)
	assert.True(t, reflect.DeepEqual(expectedAppMetadataList, actualAppMetadataList))
}
