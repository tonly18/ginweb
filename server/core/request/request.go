package request

import (
	"github.com/gin-gonic/gin"
	"server/core"
	"server/library/command"
	"time"
)

// Request 请求
type Request struct {
	ctx     *gin.Context   //context
	traceId string         //trace id
	data    map[string]any //data
}

// NewRequest
func NewRequest(ctx *gin.Context) *Request {
	return &Request{
		ctx:     ctx,
		traceId: ctx.GetString(core.TraceID), //链路追踪ID
		data:    make(map[string]any, 10),
	}
}

// GetCtx
func (r *Request) GetCtx() *gin.Context {
	return r.ctx
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

	if k, ok := key.(string); ok {
		value := r.GetData(k)
		if command.IsNil(value) {
			value = r.ctx.Value(key)
		}
		return value
	}

	return nil
}
