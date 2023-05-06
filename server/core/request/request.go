package request

import (
	"github.com/gin-gonic/gin"
	"time"
)

//Request 请求
type Request struct {
	//gin context
	ginCtx *gin.Context
	//user id
	userID int
	//client ip
	clientIP string
	//trace id
	traceId string
	//data
	data map[string]any
}

//NewRequest
func NewRequest(ctx *gin.Context, userId int, clientIp, traceId string) *Request {
	return &Request{
		ginCtx:   ctx,
		userID:   userId,
		clientIP: clientIp,
		traceId:  traceId,
		data:     make(map[string]any, 10),
	}
}

//Get Gin Context
func (r *Request) GetGinCtx() *gin.Context {
	return r.ginCtx
}

//func (r *Request) SetGinCtx(ctx *gin.Context) {
//	r.ginCtx = ctx
//}

//GetUserID
func (r *Request) GetUserID() int {
	return r.userID
}

//func (r *Request) SetUserID(uid int) {
//	r.userID = uid
//}

//client ip
func (r *Request) GetClientIP() string {
	return r.clientIP
}

//func (r *Request) SetClientIP(ip string) {
//	r.clientIP = ip
//}

//trace id
func (r *Request) GetTraceID() string {
	return r.traceId
}

//func (r *Request) SetTraceID(id string) {
//	r.traceId = id
//}

//Data
func (r *Request) GetData(key string) any {
	return r.data[key]
}
func (r *Request) SetData(key string, value any) {
	r.data[key] = value
}

//Deadline
func (r *Request) Deadline() (deadline time.Time, ok bool) {
	return r.ginCtx.Deadline()
}

//Done
func (r *Request) Done() <-chan struct{} {
	return r.ginCtx.Done()
}

//Err
func (r *Request) Err() error {
	return r.ginCtx.Err()
}

//Value
func (r *Request) Value(key any) any {
	return r.ginCtx.Value(key)
}
