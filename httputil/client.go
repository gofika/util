package httputil

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/publicsuffix"
)

type Client struct {
	client *http.Client
	ctx    context.Context
	opts   *ClientOptions
}

func NewClient(ctx context.Context, opts ...ClientOption) *Client {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	c := &Client{
		client: &http.Client{
			Timeout: time.Minute,
			Jar:     jar,
		},
		ctx: ctx,
		opts: &ClientOptions{
			timeout:               time.Minute,
			dialTimeout:           30 * time.Second,
			keepAliveTimeout:      30 * time.Second,
			maxIdleConns:          100,
			idleConnTimeout:       90 * time.Second,
			tlsHandshakeTimeout:   10 * time.Second,
			expectContinueTimeout: 1 * time.Second,
		},
	}
	for _, opt := range opts {
		opt(c.opts)
	}
	if c.opts.proxy != "" {
		c.client.Transport = &http.Transport{
			Proxy: func(_ *http.Request) (*url.URL, error) {
				return url.Parse(c.opts.proxy)
			},
			DialContext: (&net.Dialer{
				Timeout:   c.opts.dialTimeout,
				KeepAlive: c.opts.keepAliveTimeout,
			}).DialContext,
			MaxIdleConns:          c.opts.maxIdleConns,
			IdleConnTimeout:       c.opts.idleConnTimeout,
			TLSHandshakeTimeout:   c.opts.tlsHandshakeTimeout,
			ExpectContinueTimeout: c.opts.expectContinueTimeout,
		}
	}
	return c
}

func (c *Client) Close() {
	c.client.CloseIdleConnections()
}

func (c *Client) fillHeader(header http.Header, opts *RequestOptions) {
	if c.opts.userAgent != "" {
		header.Set("User-Agent", c.opts.userAgent)
	}
	header.Set("Accept", "*/*")
	header.Set("Accept-Language", "*")
	header.Set("Upgrade-Insecure-Requests", "1")
	header.Set("Connection", "keep-alive")
	if opts.referer != "" {
		header.Set("Referer", opts.referer)
	}
	if opts.contentType != "" {
		header.Set("Content-Type", opts.contentType)
	}
	for key, values := range opts.headers {
		for _, value := range values {
			if header.Get(key) != "" {
				continue
			}
			header.Add(key, value)
		}
	}
}

func (c *Client) Get(url string, opts ...RequestOption) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}
	reqOpts := &RequestOptions{}
	for _, opt := range opts {
		opt(reqOpts)
	}
	c.fillHeader(req.Header, reqOpts)
	resp, err = c.client.Do(req)
	return
}

func (c *Client) Post(url string, contentType string, body io.Reader, opts ...RequestOption) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, url, body)
	if err != nil {
		return
	}
	reqOpts := &RequestOptions{}
	WithContentType(contentType)(reqOpts)
	for _, opt := range opts {
		opt(reqOpts)
	}
	c.fillHeader(req.Header, reqOpts)
	resp, err = c.client.Do(req)
	return
}

func (c *Client) PostForm(url string, data url.Values, opts ...RequestOption) (resp *http.Response, err error) {
	return c.Post(url, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()), opts...)
}

func (c *Client) PostJSON(url string, body any, opts ...RequestOption) (resp *http.Response, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return
	}
	return c.Post(url, "application/json", bytes.NewBuffer(payload), opts...)
}
