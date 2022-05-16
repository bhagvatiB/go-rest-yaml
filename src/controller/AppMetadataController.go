package controller

import (
	"io/ioutil"
	"net/http"

	"payloadrest/src/bizlogic"
	"payloadrest/src/logger"
	"payloadrest/src/model/request"
	"payloadrest/src/model/response"
	"payloadrest/src/util"
	"payloadrest/src/validators"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type AppMetadataController struct {
	getAppMetaDataBizLogic bizlogic.AppMetadataBizLogic
}

/**
* initialize AppMetadataController
**/
func NewAppMetadataController(appMetadataLogic bizlogic.AppMetadataBizLogic) AppMetadataController {
	return AppMetadataController{getAppMetaDataBizLogic: appMetadataLogic}
}

/**
* searchAppMetadata returns the related records matching url query parameters
* if url query params are empty then returns all records
* Input : -- Http Request- search parameters
* return : list of appmetadata response in yaml format
**/
func (getData *AppMetadataController) searchAppMetadata(w http.ResponseWriter, r *http.Request) {

	// read the request parameters from http request
	values := r.URL.Query()

	// get the list of AppMetadata for search query
	AppMetadataList, err := getData.getAppMetaDataBizLogic.SearchAppMetadataLogic(values)

	if err != nil {
		var errorAppMetadataResponse = response.ErrorAppMetadataResponse{Status: "Failed", Err: err.Error()}
		w.Header().Add("Content-Type", "application/yaml")
		w.WriteHeader(http.StatusInternalServerError)

		yaml.NewEncoder(w).Encode(errorAppMetadataResponse)
		return
	}

	//convert AppMetaData to response object
	var AppMetadataResList []response.AppMetadataResponse
	for _, p := range AppMetadataList {
		AppMetadataRes := util.ConvertDaoToAppMetadataRes(p)
		AppMetadataResList = append(AppMetadataResList, AppMetadataRes)
	}

	// return AppMetadata Response in yaml format
	yaml.NewEncoder(w).Encode(AppMetadataResList)
}

/**
* function to create and store AppMetadata in the application
* Input : -- appMetadata req from endpoint in yaml format
* returns : -- AppMetadataResponse in yaml format
**/
func (getData *AppMetadataController) createAppMetadata(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)

	var appMetadataReq request.AppMetadataRequest
	var errorAppMetadataResponse = response.ErrorAppMetadataResponse{Status: "Success"}

	// unmarshal request body
	yaml.Unmarshal(reqBody, &appMetadataReq)

	// validating AppMetadata request
	_, err := validators.ValidateRequest(w, appMetadataReq)

	// if request is invalid write appropriate message in the response object
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		// mark response header with bad request
		w.Header().Add("Content-Type", "application/yaml")
		w.WriteHeader(http.StatusBadRequest)

		responseBody := map[string]string{"Invalid Request": validationErrors.Error()}
		logger.SugarLogger.Errorf("Invalid Request for creating AppMetadata", responseBody, err)

		if err := yaml.NewEncoder(w).Encode(responseBody); err != nil {
			logger.SugarLogger.Errorf("Error creating yaml response.", err)
			panic("Error creating yaml response after validation failure.")
		}

		return
	}

	// if valid, store in memory db
	err = getData.getAppMetaDataBizLogic.CreateAppMetadataLogic(appMetadataReq)

	if err != nil {
		w.Header().Add("Content-Type", "application/yaml")
		w.WriteHeader(http.StatusInternalServerError)
		errorAppMetadataResponse.Status = "Failed"
		errorAppMetadataResponse.Err = err.Error()
	}

	// return AppMetadata Response in yaml format
	yaml.NewEncoder(w).Encode(errorAppMetadataResponse)
}

/**
* Define endpoints that would be listening to incoming request to this service
**/
func (getData *AppMetadataController) HandleRequest() {
	logger.SugarLogger.Infof("Set up Listeners")

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	// create endpoint for search AppMetadata
	myRouter.HandleFunc("/api/v1/searchappmetadata", getData.searchAppMetadata)

	// create endpoit for CreateAppMetadata
	myRouter.HandleFunc("/api/v1/createappmetadata", getData.createAppMetadata).Methods("POST")

	http.Handle("/", myRouter)

	logger.SugarLogger.Infof("Endpoints are created")
}
