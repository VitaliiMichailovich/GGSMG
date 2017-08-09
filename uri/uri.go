package uri

import (
	"errors"
	"net/url"
)

func URI(uri string) (string, error) {
	url, err := url.Parse("http://" + uri)
	if err != nil {
		return "", errors.New("Incorrect URI. Check your input.")
	}
	return url.Scheme+"://"+url.Host+"/", nil

}
