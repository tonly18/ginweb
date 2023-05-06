package httpclient

import (
	"net/http"
	"time"
)

type Config struct {
	Transport http.RoundTripper
	TimeOut   time.Duration
}
