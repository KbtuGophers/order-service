package payment

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (s *Client) CreateBilling(req Request) (*Response, error) {
	aByte, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := s.Endpoint + "/payment"
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

func (s *Client) Pay(link string) (string, error) {
	res, code, err := s.handler(http.MethodGet, link, nil)
	if err != nil {
		return "", err
	}

	if code != 200 {
		return "", errors.New("unexpected status code")
	}

	return string(res), nil

}

func (s *Client) Cancel(billingID string) error {
	url := s.Endpoint + "/payment/" + billingID + "/cancel"
	_, code, err := s.handler(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	if code != 200 {
		fmt.Println(code)
		return errors.New("unexpected status code")
	}

	return nil
}
