package httpclient

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"net/url"
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
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true, //不校验服务端证书
				},
				//ForceAttemptHTTP2: true, //强制使用http2(默认值:true)
				//MaxIdleConns:      100, //最大空闲数量(默认值:100)
				DisableKeepAlives: true,

				//TLSHandshakeTimeout:   5 * time.Second,  //TLS握手超时(默认值:10)
				//ResponseHeaderTimeout: 1 * time.Second,  //限制读取response header的时间
				//ExpectContinueTimeout: 1 * time.Second,  //限制client在发送包含Expect:100-continue的header到收到继续发送body的response之间的时间等待。
				//IdleConnTimeout:       60 * time.Second, //连接空闲超时
				DisableCompression: true, //禁止压缩
			},
			Timeout: config.TimeOut, //从连接(Dial)到读完response body
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
	defer func() {
		c.httpRequest.Request.Body.Close()
		resp.Body.Close()
	}()

	//response
	//retData := bytes.Buffer{}
	//if _, err := io.Copy(&retData, resp.Body); err != nil {
	//	return nil, err
	//}
	retData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//return
	return &HttpResponse{
		Response: resp,
		Data:     retData,
		Close:    true,
	}, nil
}
