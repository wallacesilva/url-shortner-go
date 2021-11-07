package core

import (
	"os"
	"testing"
)

func TestGetBaseUrlWithoutEnvVar(t *testing.T) {
	os.Setenv("BASE_URL", "")
	uri := "/tarantino"
	url := GetBaseUrl(uri)

	if "http://0.0.0.0:8080"+uri != url {
		t.Errorf("GetBaseUrl need to return http://0.0.0.0:8080 + uri. Returned %s", url)
	}
}

func TestGetBaseUrlWithEnvVar(t *testing.T) {
	base_url := "https://in9.us"
	os.Setenv("BASE_URL", base_url)
	uri := "/tarantino"
	url := GetBaseUrl(uri)

	if base_url+uri != url {
		t.Errorf("GetBaseUrl need to return %s + uri. Returned %s", base_url, url)
	}
}

func TestGetUrlByCodeWithExistingCode(t *testing.T) {
	// TODO: Improve this test
	code := "4242"
	url_with_code := GetUrlByCode(code)
	if url_with_code != GetBaseUrl("/"+code) {
		t.Errorf("GetUrlByCode need to return BASE_URL + code. Returned %s", url_with_code)
	}
}
