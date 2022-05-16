package model

/**
* dao structure for Maintainer
**/
type Maintainer struct {
	Name  string `yaml:"name" validate:"required"`
	Email string `yaml:"email" validate:"required,email"`
}
