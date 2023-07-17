package warehouse

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (s *Client) GetProduct(storeID, productID string, quantity uint) (*Response, error) {
	aByte, err := json.Marshal(quantity)
	if err != nil {
		return nil, err
	}

	url := s.Endpoint + "/stores/" + storeID + "/product/" + productID
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
		fmt.Println(resp.Data)
		return &resp.Data, nil
	default:
		return nil, errors.New(string(respByte))
	}
}
