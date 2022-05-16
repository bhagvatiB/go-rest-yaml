package validtaors_test

import (
	"net/http/httptest"
	"payloadrest/src/logger"
	"payloadrest/src/validators"
	"payloadrest/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetUp() {
	logger.InitLogger()
}
func TestValidateRequestValid(t *testing.T) {

	SetUp()
	appMetadata := testdata.GetAppMetadataReq("1")
	w := httptest.NewRecorder()
	_, err := validators.ValidateRequest(w, appMetadata)
	assert.Nil(t, err)
}

func TestValidateRequestInValid(t *testing.T) {

	appMetadata := testdata.GetEmptyAppMetadataReq()
	w := httptest.NewRecorder()
	_, err := validators.ValidateRequest(w, appMetadata)
	assert.NotNil(t, err)
}

func TestValidateRequestInValidEmail(t *testing.T) {

	appMetadata := testdata.GetInvalidEmailAppMetadataReq()
	w := httptest.NewRecorder()
	_, err := validators.ValidateRequest(w, appMetadata)
	assert.NotNil(t, err)
}

func TestValidateRequestInValidUrl(t *testing.T) {

	appMetadata := testdata.GetInvalidEmailAppMetadataReq()
	w := httptest.NewRecorder()
	_, err := validators.ValidateRequest(w, appMetadata)
	assert.NotNil(t, err)
}
