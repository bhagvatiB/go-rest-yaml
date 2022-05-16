package util

import (
	"strings"
)

/**
* function to check if query parameter from search query is valid
* return: -- boolean value
**/
func CheckKeysPresent(key string) bool {

	switch queryParam := strings.TrimSpace(key); queryParam {
	case "title":
		return true
	case "version":
		return true
	case "maintainer.name":
		return true
	case "maintainer.email":
		return true
	case "company":
		return true
	case "website":
		return true
	case "source":
		return true
	case "license":
		return true
	case "description":
		return true
	}

	return false
}
