package utils

import (
	"io"
	"net/http"
	"regexp"
	"strings"
)

func ApiCall(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	responseBytes, _ := io.ReadAll(response.Body)
	return responseBytes, nil
}

func NormalizeCityName(city string) string {
    re := regexp.MustCompile(`[^a-zA-Z0-9]+`) // Matches non-alphanumeric characters
    return strings.ToLower(re.ReplaceAllString(city, " "))
}