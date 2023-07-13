package warehouse

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (s *Client) GetProduct(productID string) (*Response, error) {
	aByte, err := json.Marshal(productID)
	if err != nil {
		return nil, err
	}

	url := s.Endpoint + "/inventories/" + productID

	respByte, code, err := s.handler(http.MethodGet, url, aByte)
	if err != nil {
		return nil, err
	}

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
