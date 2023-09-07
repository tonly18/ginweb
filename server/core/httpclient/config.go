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
	DisableKeepAlives:  true,
	DisableCompression: true, //禁止压缩

	//ForceAttemptHTTP2: true, //强制使用http2(默认值:true)
	//MaxIdleConns:      100, //最大空闲数量(默认值:100)
	//IdleConnTimeout:       60 * time.Second, //连接空闲超时
	//TLSHandshakeTimeout:   5 * time.Second,  //TLS握手超时(默认值:10)
	//ResponseHeaderTimeout: 1 * time.Second,  //限制读取response header的时间
	//ExpectContinueTimeout: 1 * time.Second,  //限制client在发送包含Expect:100-continue的header到收到继续发送body的response之间的时间等待。
}
