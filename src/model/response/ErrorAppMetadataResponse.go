package response

/**
* response object structure for CreateAppMetadata
**/
type ErrorAppMetadataResponse struct {
	Status string `yaml:"status"`
	Err    string `yaml:"err"`
}
