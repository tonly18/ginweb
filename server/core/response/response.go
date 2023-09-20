package response

// Response
type Response struct {
	Code uint32 `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (r *Response) GetCode() uint32 {
	return r.Code
}
func (r *Response) SetCode(code uint32) {
	r.Code = code
}

func (r *Response) GetData() any {
	return r.Data
}
func (r *Response) SetData(data any) {
	r.Data = data
}

func (r *Response) GetMsg() string {
	return r.Msg
}
func (r *Response) SetMsg(msg string) {
	r.Msg = msg
}
