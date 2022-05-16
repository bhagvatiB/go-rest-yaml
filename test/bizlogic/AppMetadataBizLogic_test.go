package bizlogic__test

import (
	"payloadrest/src/bizlogic"
	"payloadrest/src/client"
	"payloadrest/src/dao"
	"payloadrest/src/logger"
	"payloadrest/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

var appMetadataBizLogic bizlogic.AppMetadataBizLogic

func SetUp() {
	logger.InitLogger()
	inMemDBClient := client.NewInMemDBClient()
	inMemAppMetadataDao := dao.NewInMemAppMetadataDao(inMemDBClient)
	appMetadataBizLogic = bizlogic.NewAppMetadataBizLogic(inMemAppMetadataDao)
}

func TestCreateAppMetadataValidReq(t *testing.T) {
	SetUp()
	expectedAppMetadata := testdata.GetAppMetadataReq("1")

	err := appMetadataBizLogic.CreateAppMetadataLogic(expectedAppMetadata)
	assert.Nil(t, err)
}
