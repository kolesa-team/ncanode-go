package ncanode

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const _timeout = 60 * time.Second

type Option func(c *Client) error

func WithTimeout(t time.Duration) Option {
	return func(c *Client) error {
		c.client.Timeout = t
		return nil
	}
}

func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) error {
		c.client = hc
		return nil
	}
}

type Client struct {
	host   string
	client *http.Client
}

func NewClient(addr string, opts ...Option) (*Client, error) {
	if addr == "" {
		return nil, errors.New("ncanode: address invalid or empty")
	}

	client := &Client{
		host:   addr,
		client: &http.Client{Timeout: _timeout},
	}

	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

type modifier func(data []byte) ([]byte, error)

func (c *Client) call(body, reply interface{}, mods ...modifier) error {
	buf, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("encode payload: %w", err)
	}

	for _, mod := range mods {
		data, err := mod(buf)
		if err != nil {
			return fmt.Errorf("apply modifier: %w", err)
		}
		buf = data
	}

	req, err := http.NewRequest("POST", c.host, bytes.NewReader(buf))
	if err != nil {
		return fmt.Errorf("create request %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read request: %w", err)
	}
	defer resp.Body.Close()

	var apiResp apiResponse
	if err := json.Unmarshal(data, &apiResp); err != nil {
		return fmt.Errorf("read api response: %w", err)
	}
	if apiResp.Status != 0 || apiResp.Message != "" {
		return apiResp
	}

	if err := json.Unmarshal(data, &reply); err != nil {
		return fmt.Errorf("decode payload: %w", err)
	}

	return nil
}
