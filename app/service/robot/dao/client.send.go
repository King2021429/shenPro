package dao

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Send(str []byte, url string) {
	reqBodyReader := strings.NewReader(string(str))
	req, err := http.NewRequest(http.MethodPost, url, reqBodyReader)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{Timeout: 5 * time.Second, Transport: &http.Transport{IdleConnTimeout: time.Second * 5, MaxIdleConns: 300, MaxIdleConnsPerHost: 300, MaxConnsPerHost: 2000}}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
