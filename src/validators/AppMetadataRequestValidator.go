package validators

import (
	"net/http"
	"payloadrest/src/logger"
	"payloadrest/src/model"
	"payloadrest/src/model/request"

	"github.com/go-playground/validator/v10"
)

/**
* validate AppMetadata request
**/
func ValidateRequest(w http.ResponseWriter, AppMetadataReq request.AppMetadataRequest) (http.ResponseWriter, error) {
	logger.SugarLogger.Infof("Validating AppMetadata request ", AppMetadataReq)
	validate := validator.New()

	err := validate.Struct(AppMetadataReq)
	if err == nil {
		//for
		err = validateMaintainers(validate, AppMetadataReq.Maintainers)
	}

	logger.SugarLogger.Infof("Error in validation ", err)
	return w, err
}

func validateMaintainers(validate *validator.Validate, maintainers []model.Maintainer) error {

	for _, maintainer := range maintainers {
		err := validate.Struct(maintainer)
		if err != nil {
			return err
		}
	}
	return nil
}
