package httpclient

import (
	"crypto/tls"
	"github.com/spf13/cast"
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
	httpClient  *http.Client
	httpRequest *Request
}

func NewHttpClient(config *Config) *HttpClient {
	if config.TimeOut == 0 {
		config.TimeOut = time.Second * 10
	}

	return &HttpClient{
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, //不校验服务端证书
				},
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          10,
				IdleConnTimeout:       60 * time.Second, //连接空闲超时
				TLSHandshakeTimeout:   5 * time.Second,  //tls握手超时
				ExpectContinueTimeout: 1 * time.Second,
				DisableCompression:    true,
			},
			Timeout: config.TimeOut,
		},
		httpRequest: nil,
	}
}

func (c *HttpClient) Get(rawurl string, params map[string]any, headtype ...string) (*HttpResponse, error) {
	urlValue := url.Values{}
	for k, v := range params {
		urlValue.Set(k, cast.ToString(v))
	}

	req, err := NewRequest("GET", rawurl+"?"+urlValue.Encode(), nil, headtype...)
	if err != nil {
		return nil, err
	}
	defer req.Request.Body.Close()

	resp, err := c.httpClient.Do(req.Request)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		Response: resp,
	}, err
}

func (c *HttpClient) Post(rawurl string, data []byte, headtype ...string) (*HttpResponse, error) {
	req, err := NewRequest("POST", rawurl, data, headtype...)
	if err != nil {
		return nil, err
	}
	defer req.Request.Body.Close()

	c.httpRequest = req
	resp, err := c.httpClient.Do(c.httpRequest.Request)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		resp,
	}, err
}

func (c *HttpClient) NewRequest(method, url string, body []byte, headtype ...string) *HttpClient {
	if req, err := NewRequest(method, url, body, headtype...); err != nil {
		return nil
	} else {
		c.httpRequest = req
	}

	return c
}

func (c *HttpClient) SetHeader(params map[string]any) *HttpClient {
	for key, value := range params {
		c.httpRequest.Request.Header.Set(key, cast.ToString(value))
	}

	return c
}

func (c *HttpClient) Do() (*HttpResponse, error) {
	resp, err := c.httpClient.Do(c.httpRequest.Request)
	if err != nil {
		return nil, err
	}
	defer c.httpRequest.Request.Body.Close()

	return &HttpResponse{
		Response: resp,
	}, nil
}
