package cepv1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {}

func NewClient() Client {
    return Client{}
}

func (c Client) handler(input string, output interface{}) error {
	url := fmt.Sprint("https://brasilapi.com.br/api", "/cep/v1/", input)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return  err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return  err
	}
	if res.StatusCode != http.StatusOK {
        return fmt.Errorf("error to call %s. status code: %d", "/cep/v1/", res.StatusCode)
	}
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return  err
	}
	if err := json.Unmarshal(bytes, &output); err != nil {
		return  err
	}
    return nil
}
