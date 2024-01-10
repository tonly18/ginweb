package request

import (
	"context"
	"github.com/gin-gonic/gin"
	"server/library/command"
	"time"
)

// Request 请求
type Request struct {
	ctx      *gin.Context   //context
	userID   int            //user id
	clientIP string         //client ip
	traceId  string         //trace id
	data     map[string]any //data
}

// NewRequest
func NewRequest(ctx *gin.Context, userId int, clientIp, traceId string) *Request {
	return &Request{
		ctx:      ctx,
		userID:   userId,
		clientIP: clientIp,
		traceId:  traceId,
		data:     make(map[string]any, 10),
	}
}

// GetCtx
func (r *Request) GetCtx() context.Context {
	return r.ctx
}

// GetUserID
func (r *Request) GetUserID() int {
	return r.userID
}

// GetClientIP
func (r *Request) GetClientIP() string {
	return r.clientIP
}

// GetTraceID
func (r *Request) GetTraceID() string {
	return r.traceId
}

// GetData
func (r *Request) GetData(key string) any {
	return r.data[key]
}
func (r *Request) SetData(key string, value any) {
	r.data[key] = value
}

// Deadline
func (r *Request) Deadline() (deadline time.Time, ok bool) {
	return r.ctx.Deadline()
}

// Done
func (r *Request) Done() <-chan struct{} {
	return r.ctx.Done()
}

// Err
func (r *Request) Err() error {
	return r.ctx.Err()
}

// Value
func (r *Request) Value(key any) any {
	if r.ctx == nil {
		return nil
	}

	value := r.ctx.Value(key)
	if command.IsValueNil(value) {
		if k, ok := key.(string); ok {
			value = r.GetData(k)
		}
	}
	return value
}
