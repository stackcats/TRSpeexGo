package util

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// Err ...
type Err struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// Download ...
func Download(url string) (string, error) {
	out, err := ioutil.TempFile("./uploads", "tmp")
	if err != nil {
		return "", err
	}

	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	e := &Err{}
	err = json.Unmarshal(b, e)
	if err == nil {
		return "", errors.New(e.ErrMsg)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return out.Name(), nil
}
