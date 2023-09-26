package httpclient

import (
	"crypto/tls"
	"net/http"
	"time"
)

type Config struct {
	CheckRedirect func(req *http.Request, via []*http.Request) error
	Jar           http.CookieJar
	TimeOut       time.Duration
}

/*
http.Transport内都会维护一个自己的空闲连接池,如果每个client都创建一个新的http.Transport,就会导致底层的TCP连接无法复用.
如果网络请求过大,上面这种情况会导致协程数量变得非常多,导致服务不稳定.
*/
var transport = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true, //不校验服务端证书
	},
	DisableKeepAlives:     false,
	MaxConnsPerHost:       2000,
	MaxIdleConns:          100,
	MaxIdleConnsPerHost:   1000,
	IdleConnTimeout:       90 * time.Second,
	DisableCompression:    true,
	ExpectContinueTimeout: 1 * time.Second,
}
