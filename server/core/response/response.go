package response

//Response
type Response struct {
	Code int  `json:"code"`
	Data any  `json:"data"`
	Msg  any  `json:"msg"`
	Type int8 `json:"-"` //保存到日志文件: 0不保存 | 1保存
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
