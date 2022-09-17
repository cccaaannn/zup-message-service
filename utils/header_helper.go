package utils

import (
	"errors"
	"net/http"
	"strings"
)

func GetTokenFromHeader(r *http.Request) (string, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return "", errors.New("Invalid token format")
	}

	reqToken = strings.TrimSpace(splitToken[1])

	return reqToken, nil
}
