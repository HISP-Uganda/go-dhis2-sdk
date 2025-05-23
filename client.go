package dhis2

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	BaseURL  string
	Username string
	Password string
	Resty    *resty.Client
}

func NewClient(url, user, pass string) *Client {
	client := resty.New().
		SetBaseURL(url).
		SetBasicAuth(user, pass).
		SetHeader("Content-Type", "application/json")

	return &Client{
		BaseURL:  url,
		Username: user,
		Password: pass,
		Resty:    client,
	}
}

func (c *Client) Post(path string, payload any, queryParams ...map[string]string) (*resty.Response, error) {
	r := c.Resty.R().SetBody(payload)
	if len(queryParams) > 0 {
		r.SetQueryParams(queryParams[0])
	}
	return r.Post(path)
}

func (c *Client) Get(path string, result any, queryParams ...map[string]string) (*resty.Response, error) {
	r := c.Resty.R().SetResult(result)
	if len(queryParams) > 0 {
		r.SetQueryParams(queryParams[0])
	}
	return r.Get(path)
}

func (c *Client) Put(path string, payload any, queryParams ...map[string]string) (*resty.Response, error) {
	r := c.Resty.R().SetBody(payload)
	if len(queryParams) > 0 {
		r.SetQueryParams(queryParams[0])
	}
	return r.Put(path)
}

func (c *Client) Delete(path string, queryParams ...map[string]string) (*resty.Response, error) {
	r := c.Resty.R()
	if len(queryParams) > 0 {
		r.SetQueryParams(queryParams[0])
	}
	return r.Delete(path)
}
