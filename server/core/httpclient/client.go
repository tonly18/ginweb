package httpclient

import (
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"net/url"
	"server/core/pool"
	"strings"
	"time"
)

type HttpClient struct {
	httpClient  *http.Client
	httpRequest *Request
}

func NewHttpClient(config *Config) *HttpClient {
	if config.TimeOut == 0 {
		config.TimeOut = time.Second * 5
	}

	return &HttpClient{
		httpClient: &http.Client{
			Transport: transport,
			Timeout:   config.TimeOut, //从连接(Dial)到读完response body
		},
		httpRequest: nil,
	}
}

func (c *HttpClient) Get(rawurl string, params map[string]any) *HttpClient {
	if len(params) > 0 {
		urlValue := url.Values{}
		for k, v := range params {
			urlValue.Set(k, cast.ToString(v))
		}
		rawurl = fmt.Sprintf(`%v?%v`, rawurl, urlValue.Encode())
	}
	req, err := NewRequest(http.MethodGet, rawurl, nil)
	if err != nil {
		return nil
	}
	c.httpRequest = req

	return c
}

func (c *HttpClient) Post(rawurl string, params []byte) *HttpClient {
	req, err := NewRequest(http.MethodPost, rawurl, params)
	if err != nil {
		return nil
	}
	c.httpRequest = req

	return c
}

func (c *HttpClient) NewRequest(method, url string, body []byte) *HttpClient {
	if req, err := NewRequest(strings.ToUpper(method), url, body); err != nil {
		return nil
	} else {
		c.httpRequest = req
	}

	return c
}

func (c *HttpClient) SetHeader(params map[string]any) *HttpClient {
	c.httpRequest.SetHeader(params)

	//return
	return c
}

func (c *HttpClient) Do() (*HttpResponse, error) {
	resp, err := c.httpClient.Do(c.httpRequest.Request)
	if err != nil {
		return nil, err
	}

	//response
	retData := pool.Buffer4096Pool.Get().(*bytes.Buffer)
	retData.Reset()
	defer func() {
		pool.Buffer4096Pool.Put(retData)
		c.httpRequest.Request.Body.Close()
		resp.Body.Close()
	}()
	if _, err := io.Copy(retData, resp.Body); err != nil {
		return nil, err
	}

	//return
	return &HttpResponse{
		Response: resp,
		Data:     retData.Bytes(),
		Close:    true,
	}, nil
}
