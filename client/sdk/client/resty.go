package client

import "github.com/go-resty/resty/v2"

type restyClient struct {
	cl *resty.Client
}

type ClientOpt func(r *restyClient)

func WithBasicAuth(username, password string) ClientOpt {
	return func(r *restyClient) {
		r.cl.SetBasicAuth(username, password)
	}
}

func New(baseURL string, opts ...ClientOpt) Client {
	cl := resty.
		New().
		SetBaseURL(baseURL)
	return NewForClient(cl, opts...)

}

func NewForClient(rst *resty.Client, opts ...ClientOpt) Client {
	cl := &restyClient{
		cl: rst,
	}
	for _, opt := range opts {
		opt(cl)
	}
	return cl
}
