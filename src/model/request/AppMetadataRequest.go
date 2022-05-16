package request

import (
	"payloadrest/src/model"
)

/**
* request object structure for AppMetadata
**/
type AppMetadataRequest struct {
	Title       string             `yaml:"title" validate:"required"`
	Version     string             `yaml:"version" validate:"required"`
	Maintainers []model.Maintainer `yaml:"maintainers" validate:"required"`
	Company     string             `yaml:"company" validate:"required"`
	Website     string             `yaml:"website" validate:"required,url"`
	Source      string             `yaml:"source" validate:"required,url"`
	License     string             `yaml:"license" validate:"required"`
	Description string             `yaml:"description" validate:"required"`
}
