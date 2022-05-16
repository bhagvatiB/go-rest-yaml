package model

/**
* dao structure for AppMetadata
**/
type AppMetadata struct {
	Id          string
	Title       string
	Version     string
	Maintainers []Maintainer
	Company     string
	Website     string
	Source      string
	License     string
	Description string
}

// generate unique key for AppMetadata table
// using title & version of AppMetadata
func GenerateKey(title string, version string) string {
	return title + "_" + version
}
