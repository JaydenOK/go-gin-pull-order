package service

import (
	"io"
	"net/http"
)

type HttpClient interface {
	Get(url string) (*http.Response, error)
	Post(url string, bodyType string, body io.Reader) (*http.Response, error)
	Put(url string, body io.Reader) (*http.Response, error)
	Delete(url string, body io.Reader) (*http.Response, error)
}

type MyHttpClient struct {
}

func (httpClient MyHttpClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

func (httpClient MyHttpClient) Post(url string, bodyType string, body io.Reader) (*http.Response, error) {
	return http.Post(url, bodyType, body)
}

func (httpClient MyHttpClient) Put(url string, body io.Reader) (*http.Response, error) {
	return httpClient.executeHTTPRequest(http.MethodPut, url, body)
}

func (httpClient MyHttpClient) Delete(url string, body io.Reader) (*http.Response, error) {
	return httpClient.executeHTTPRequest(http.MethodDelete, url, body)
}

func (httpClient MyHttpClient) executeHTTPRequest(method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		//if debugEnable {
		//	log.Printf("Error when creating %s request %s.", method, err.Error())
		//}
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		//if debugEnable {
		//	log.Printf("Error while calling url: %s\n Error: %s", url, err.Error())
		//}
		return nil, err
	}

	return resp, nil
}
