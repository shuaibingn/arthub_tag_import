package internal

import (
	"io"
	"net/http"
)

var client = new(http.Client)

func HttpDo(method, url string, reqBody io.Reader, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Get(url string, header map[string]string) ([]byte, error) {
	return HttpDo("GET", url, nil, header)
}

func Post(url string, body io.Reader, header map[string]string) ([]byte, error) {
	return HttpDo("POST", url, body, header)
}
