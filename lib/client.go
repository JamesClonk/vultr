package lib

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	version         = "0.0.1"
	mediaType       = "application/json"
	defaultEndpoint = "https://api.vultr.com/"
)

type Client struct {
	// HTTP client for communication with the Vultr API
	client *http.Client

	// User agent for HTTP client
	UserAgent string

	// Endpoint URL for API requests
	Endpoint *url.URL

	// API key for accessing the Vultr API
	ApiKey string
}

type Options struct {
	// HTTP client for communication with the Vultr API
	HttpClient *http.Client

	// User agent for HTTP client
	UserAgent string

	// Endpoint URL for API requests
	Endpoint string
}

func NewClient(apiKey string, options *Options) (*Client, error) {
	userAgent := "vultr-go/" + version
	client := http.DefaultClient
	endpoint, err := url.Parse(defaultEndpoint)
	if err != nil {
		return nil, err
	}

	if options != nil {
		if options.HttpClient != nil {
			client = options.HttpClient
		}
		if options.UserAgent != "" {
			userAgent = options.UserAgent
		}
		if options.Endpoint != "" {
			endpoint, err = url.Parse(options.Endpoint)
			if err != nil {
				return nil, err
			}
		}
	}

	return &Client{
		UserAgent: userAgent,
		client:    client,
		Endpoint:  endpoint,
		ApiKey:    apiKey,
	}, nil
}

func (c *Client) get(path string, data interface{}) error {
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return err
	}

	return c.do(req, data)
}

func (c *Client) post(path string, values url.Values, data interface{}) error {
	req, err := c.newRequest("POST", path, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	return c.do(req, data)
}

func (c *Client) do(req *http.Request, data interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		return err
	}

	if data != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// avoid unmarshalling problem because Vultr API returns empty array instead of empty map when it shouldn't!
		if string(body) == `[]` {
			data = nil
		} else {
			if err := json.Unmarshal(body, data); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	relPath, err := url.Parse(path + "?api_key=" + c.ApiKey)
	if err != nil {
		return nil, err
	}

	url := c.Endpoint.ResolveReference(relPath)

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Accept", mediaType)

	if req.Method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

func checkResponse(resp *http.Response) error {
	// 200 is OK
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return errors.New(string(data))
}
