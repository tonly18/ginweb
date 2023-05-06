package httpclient

import (
	"bytes"
	"net/http"
	"strings"
)

type Request struct {
	Request *http.Request
}

func NewRequest(method, url string, data []byte, headtype ...string) (*Request, error) {
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	//二进制流类型
	if len(headtype) > 0 {
		req.Header.Set("Content-Type", headtype[0])
	} else {
		req.Header.Set("Content-Type", "application/octet-stream")
	}

	return &Request{
		Request: req,
	}, nil
}
