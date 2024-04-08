package client

import "github.com/go-resty/resty/v2"

type restyClient struct {
	cl *resty.Client
}

func New(baseURL string) Client {
	cl := resty.
		New().
		SetBaseURL(baseURL)
	return NewForClient(cl)
}

func NewForClient(cl *resty.Client) Client {
	return &restyClient{
		cl: cl,
	}
}
