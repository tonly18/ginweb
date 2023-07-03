package response

//Response
type Response struct {
	Code int `json:"code"`
	Data any `json:"data"`
	Msg  any `json:"msg"`
}

func (r *Response) GetCode() int {
	return r.Code
}
func (r *Response) SetCode(code int) {
	r.Code = code
}

func (r *Response) GetData() any {
	return r.Data
}
func (r *Response) SetData(data any) {
	r.Data = data
}

func (r *Response) GetMsg() any {
	return r.Msg
}
func (r *Response) SetMsg(msg any) {
	r.Msg = msg
}
