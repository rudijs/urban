package lib

import (
	"io/ioutil"
	"net/http"
)

// GetURL get URL
func GetURL(url string) (string, error) {

	resp, err1 := http.Get(url)

	if err1 != nil {
		return "", err1
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		return "", err2
	}

	json := string(body[:len(body)])

	return json, nil

}
