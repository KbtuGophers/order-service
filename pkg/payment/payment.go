package payment

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (s *Client) CreateBilling(url string, req Request) (*Response, error) {
	aByte, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	respByte, code, err := s.handler(http.MethodPost, url, aByte)
	if err != nil {
		return nil, err
	}

	// check response code
	switch code {
	case 200:
		// unmarshal response data
		resp := new(HandlerResponse)
		err = json.Unmarshal(respByte, &resp)
		if err != nil {
			return nil, err
		}

		return &resp.Data, nil
	default:
		return nil, errors.New(string(respByte))
	}

}

func (s *Client) Pay(url string) error {
	_, code, err := s.handler(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New("unexpected status code")
	}

	return nil

}

func (s *Client) Cancel(url string) error {
	_, code, err := s.handler(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	if code != 200 {
		return errors.New("unexpected status code")
	}

	return nil
}
