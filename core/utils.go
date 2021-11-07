package core

import "os"

// TODO: fix me please
// var baseUrl string = os.Getenv("BASE_URL")

func GetBaseUrl(uri string) string {

	var url string
	var baseUrl string

	baseUrl = os.Getenv("BASE_URL")

	if baseUrl != "" {
		url = baseUrl + uri
	} else {
		url = "http://0.0.0.0:8080" + uri
	}
	return url
}

// TODO: Add tests
func GetUrlByCode(code string) string {
	return GetBaseUrl("/" + code)
}
