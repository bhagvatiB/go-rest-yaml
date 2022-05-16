//go:build wireinject
// +build wireinject

package main

import (
	"payloadrest/src/bizlogic"
	"payloadrest/src/client"
	"payloadrest/src/controller"
	"payloadrest/src/dao"

	"github.com/google/wire"
)

/**
* using wire dependency injections to initialize the server
* return appMetadataController
**/
func InitializeServer() controller.AppMetadataController {
	wire.Build(controller.NewAppMetadataController,
		bizlogic.NewAppMetadataBizLogic,
		dao.NewInMemAppMetadataDao,
		client.NewInMemDBClient)
	return controller.AppMetadataController{}
}
