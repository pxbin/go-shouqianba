package shouqianba

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

const (
	// 收钱吧接入域名 api_domain
	baseURL = "https://vsi-api.shouqianba.com"

	// apiGatewayBaseURL WAP支付收银台网关地址
	apiGatewayBaseURL = "https://qr.shouqianba.com/gateway"

	headerAuthorizationKey = "Authorization"
	headerContentTypeKey   = "Content-Type"
	jsonContentType        = "application/json"
)

type Config struct {
	AppID       string
	Code        string
	DeviceID    string
	VendorSN    string
	VendorKey   string
	TerminalSN  string
	TerminalKey string
}

type ClientOptionFunc func(config *Config)

// Client mrepresents a Shouqianba REST API Client
type Client struct {
	// sync.Mutex
	mu sync.Mutex
	// HTTP client used to communicate with the API.
	client *http.Client

	config *Config

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Shouqianba API.
	Gateway  *GatewayService
	UPay     *UPayService
	Terminal *TerminalService
}

// Service represents a service on the Shouqianba API
type service struct {
	client *Client
}

// NewClient returns a new API client.
func NewClient(config *Config, options ...ClientOptionFunc) *Client {
	for _, option := range options {
		option(config)
	}

	c := &Client{
		client: &http.Client{},
		config: config,
	}
	c.initialize()
	return c
}

func (c *Client) initialize() {
	if c.client == nil {
		c.client = &http.Client{}
	}

	c.common.client = c
	c.Gateway = (*GatewayService)(&c.common)
	c.UPay = (*UPayService)(&c.common)
	c.Terminal = (*TerminalService)(&c.common)
}

// RequestOption represents an option that can modify an http.Request.
type RequestOption func(req *http.Request)

// WithAuthentication sets the request's Authorization header to the provided value.
func WithAuthentication(auth string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set(headerAuthorizationKey, auth)
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, url string, body interface{}, opts ...RequestOption) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		bs, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(bs)
	}

	req, err := http.NewRequest(method, url, buf)

	if err != nil {
		return nil, err
	}

	req.Header.Set(headerContentTypeKey, jsonContentType)

	for _, opt := range opts {
		opt(req)
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer interface,
// the raw response body will be written to v, without attempting to first
// decode it. If v is nil, and no error happens, the response is returned as is.
// If rate limit is exceeded and reset time is in the future, Do returns
// *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it
// is canceled or times out, ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

func sign(req interface{}, sn, key string) (string, error) {
	bs, err := json.Marshal(req)
	if err != nil {
		return "", err
	}
	signstr := string(bs) + key
	sum := md5.Sum([]byte(signstr))
	return sn + " " + hex.EncodeToString(sum[:]), nil
}
