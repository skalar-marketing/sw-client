package sw

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	storeAPI        = "store-api"
	contextEndpoint = "/context"
)

// Client is a client for the Shopware storefront API.
type Client struct {
	baseURL     string
	swAccessKey string
}

// NewClient creates a new Client for given host and access key.
func NewClient(host, accessKey string) (*Client, error) {
	u, err := url.Parse(host)

	if err != nil {
		return nil, err
	}

	u.Path = storeAPI
	return &Client{
		baseURL:     u.String(),
		swAccessKey: accessKey,
	}, nil
}

// Context gives some general information about the store and the user.
func (client *Client) Context() (*Context, error) {
	result := new(Context)

	if err := client.performGet(contextEndpoint, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) performGet(endpoint string, result any) error {
	req, err := http.NewRequest("GET", client.baseURL+endpoint, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("sw-access-key", client.swAccessKey)
	c := http.Client{}
	resp, err := c.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("error requesting data (%d): %s", resp.StatusCode, string(body)))
	}

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(result); err != nil {
		return err
	}

	return nil
}
