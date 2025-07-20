package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	TOKEN_STRING := strings.Trim(strings.TrimPrefix(headers.Get("Authorization"), "ApiKey "), " ")
	if len(TOKEN_STRING) <= 0 {
		return "", fmt.Errorf("malformed authorization header: ApiKey")
	}
	return TOKEN_STRING, nil
}
