package pkg

import (
	"errors"
	"regexp"
	"strings"
)

func Raw2Url(rawUrl string) (url string, err error) {
	if !strings.Contains(rawUrl, "http") {
		return "", errors.New("no http url")
	}
	urlRegex := regexp.MustCompile(`(https?://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(?:/[a-zA-Z0-9-._~:/?#\[\]@!$&'()*+,;=%]*)?)`)
	url = urlRegex.FindString(rawUrl)
	return
}
