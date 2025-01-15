package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an api_key from headers of
// an http request
// Example:
// Authorization: ApiKey {insert api_key here}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing authorization header")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first header parameter")
	}

	key := vals[1]
	return key, nil
}
