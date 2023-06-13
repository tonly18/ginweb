package httpclient

import (
	"io/ioutil"
	"net/http"
)

type HttpResponse struct {
	Response *http.Response
	Data     []byte
	Close    bool
}

func NewHttpResponse(response *http.Response) *HttpResponse {
	return &HttpResponse{
		Response: response,
	}
}

func (c *HttpResponse) GetHeaderCode() int {
	return c.Response.StatusCode
}

func (c *HttpResponse) GetDataFromHeader(key string) string {
	return c.Response.Header.Get(key)
}

func (c *HttpResponse) GetData() ([]byte, error) {
	defer c.Response.Body.Close()

	if c.Close {
		return c.Data, nil
	}
	c.Close = true
	if data, err := ioutil.ReadAll(c.Response.Body); err != nil {
		return nil, err
	} else {
		c.Data = data
		return c.Data, nil
	}
}
