package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var HttpClient = &http.Client{
	Timeout: 15 * time.Second,
}

func QueryGet(url string, params map[string]string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()
	for x := range headers {
		req.Header.Set(x, headers[x])
	}
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func QueryPost(url string, params map[string]string, headers map[string]string, queryType string, body io.Reader) ([]byte, error) {

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()
	for x := range headers {
		req.Header.Set(x, headers[x])
	}
	switch queryType {
	case "json":
		req.Header.Set("Content-Type", "application/json")
	case "x-www-form-urlencoded":
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
