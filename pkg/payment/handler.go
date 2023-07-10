package payment

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	client *http.Client
	mutex  *sync.Mutex
}

func NewClient() *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	client.Timeout = 30 * time.Second

	return &Client{
		client: client,
		mutex:  &sync.Mutex{},
	}
}

func (s *Client) handler(method string, url string, body []byte) ([]byte, int, error) {
	// setup request
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	// setup request header
	req.Header.Add("Content-Type", "application/json")

	res, err := s.client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	// read response body
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

	return resBody, res.StatusCode, nil
}
