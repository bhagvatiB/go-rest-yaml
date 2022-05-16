package response

import (
	"payloadrest/src/model"
)

/**
* response object structure for AppMetadata
**/
type AppMetadataResponse struct {
	Title       string             `yaml:"title"`
	Version     string             `yaml:"version"`
	Maintainers []model.Maintainer `yaml:"maintainers"`
	Company     string             `yaml:"company"`
	Website     string             `yaml:"website"`
	Source      string             `yaml:"source"`
	License     string             `yaml:"license"`
	Description string             `yaml:"description"`
}
