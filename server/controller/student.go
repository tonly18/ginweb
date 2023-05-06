package controller

import (
	"fmt"
	"server/core/controller"
	"server/core/request"
	"server/core/response"
)

//StudentHandler
type StudentHandler struct {
	controller.BaseHandle
}

//func (c *StudentHandler) PreHandler(req *request.Request) {
//	fmt.Println("PerHandler----------------")
//}

func (c *StudentHandler) Handler(req *request.Request) *response.Response {

	fmt.Println("00000000000000000000000000")

	//panic("panic......")

	//ip := req.GinCtx.ClientIP()
	//ip := req.ClientIP
	ip := req.GetClientIP()
	fmt.Println("ip::::", ip)

	//go func() {
	//	fmt.Println("00000000000000")
	//	panic("goroutine panic......")
	//}()

	//time.Sleep(5 * time.Second)
	fmt.Println("11111111111111111111111")

	//return &response.Response{
	//	Code: 100000,
	//	Err:  fmt.Sprintf(`error: %v`, "error is xxx"),
	//}
	//fmt.Println("222222222222222222222")

	//return
	return &response.Response{
		Data: "this is ok!",
	}
}

func (c *StudentHandler) PostHandler(req *request.Request) {
	fmt.Println("PostHandler+++++++++++++++++")
}
