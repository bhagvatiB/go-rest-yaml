package controller_test

import (
	"net/http"
	"net/http/httptest"
	"payloadrest/src/bizlogic"
	"payloadrest/src/client"
	"payloadrest/src/controller"
	"payloadrest/src/dao"
	"payloadrest/src/logger"
	"strings"
	"testing"
)

var appMetadataController controller.AppMetadataController

func Setup() {
	logger.InitLogger()
	inMemDBClient := client.NewInMemDBClient()
	inMemAppMetadataDao := dao.NewInMemAppMetadataDao(inMemDBClient)
	appMetadataBizLogic := bizlogic.NewAppMetadataBizLogic(inMemAppMetadataDao)
	appMetadataController = controller.NewAppMetadataController(appMetadataBizLogic)
}

func TestCreateAppMetadataValidReq(t *testing.T) {

	Setup()
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/createAppMetadata", strings.NewReader("title: Valid App 1"+
		"version: 0.0.1"+
		"maintainers:"+
		"- name: firstmaintainer app1"+
		"email: firstmaintainer@hotmail.com"+
		"- name: firstmaintainer app2"+
		"email: firstmaintaine2r@hotmail.com"+
		"company: Random Inc."+
		"website: https://website.com"+
		"source: https://github.com/random/repo"+
		"license: Apache-2.0"+
		"description: Interesting Title Some application content, and description"))

	r.Header.Set("Content-Type", "text")

	controller.ExportCreateAppMetadata(&appMetadataController, w, r)
}
