package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	TOKEN_STRING := strings.Trim(strings.TrimPrefix(headers.Get("Authorization"), "Bearer "), " ")
	if len(TOKEN_STRING) <= 0 {
		return "", fmt.Errorf("malformed authorization header: Bearer")
	}
	return TOKEN_STRING, nil
}
