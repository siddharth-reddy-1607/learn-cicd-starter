package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := make(http.Header)
	headers.Add("Accept", "text/html")
	headers.Add("Content", "en-us")
	apiKey, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded && apiKey != "" {
		t.Fatalf("With no authorization header: Expected err: %v and empty API Key but got err:%v and API Key:%v", ErrNoAuthHeaderIncluded, err, apiKey)
	}
	headers.Set("Authorization", "ApiKey:myAPIKey")
	apiKey, err = GetAPIKey(headers)
	if err.Error() != "malformed authorization header" && apiKey != "" {
		t.Fatalf("With malformed authorization header: Expected err: malformed authorization header and empty API Key but got err:%v and API Key:%v", err, apiKey)
	}
	headers.Set("Authorization", "ApiKey myAPIKey")
	apiKey, err = GetAPIKey(headers)
	if err != nil && apiKey == "" {
		t.Fatalf("With correct authorization header: Expected err: nil and non-empty API Key but got err:%v and API Key:%v", err, apiKey)
	}
}
