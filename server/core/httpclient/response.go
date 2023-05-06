package httpclient

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type HttpResponse struct {
	Response *http.Response
}

func NewHttpResponse(response *http.Response) *HttpResponse {
	return &HttpResponse{
		Response: response,
	}
}

func (c *HttpResponse) GetHeaderCode() int {
	return c.Response.StatusCode
}

func (c *HttpResponse) GetHeadByKey(key string) (any, error) {
	if val := c.Response.Header.Get(key); val != "" {
		return val, nil
	}

	return nil, errors.New(key + " is empty!")
}

func (c *HttpResponse) GetData() ([]byte, error) {
	defer c.Response.Body.Close()

	if data, err := ioutil.ReadAll(c.Response.Body); err != nil {
		return nil, err
	} else {
		return data, err
	}
}
