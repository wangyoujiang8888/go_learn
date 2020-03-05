package network

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var httpStatusError = errors.New("Http Request Error")
var readBodyError = errors.New("Read Body Error")

func Get(url string, params map[string]string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("error")
	}
	query := req.URL.Query()
	if params != nil {
		for k, v := range params {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil,httpStatusError
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, readBodyError
	}
	return body, err
}

func Post(url string, body interface{}, params map[string]string, headers map[string]string) ([]byte, error) {
	var JsonBody []byte
	if body != nil {
		var err error
		JsonBody, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(JsonBody))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	query := req.URL.Query()
	if params != nil {
		for k, v := range params {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}
	req.Header.Set("Content-type", "application/json")
	if headers != nil {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil,httpStatusError
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, readBodyError
	}
	return respBody, err
}
