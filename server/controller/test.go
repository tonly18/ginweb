package controller

import (
	"fmt"
	"server/core/controller"
	"server/core/request"
	"server/core/response"
	"server/library/command"
	"server/model"
)

//TestHandler Test测试接口
type TestHandler struct {
	controller.BaseHandle
}

//PreHandler 在Handler之前执行
func (c *TestHandler) PreHandler(req *request.Request) {
	fmt.Println("test.PerHandler----------------")
}

//Handler 业务处理
func (c *TestHandler) Handler(req *request.Request) *response.Response {
	bag := model.NewBag(req)
	err := bag.Get(22)
	//panic("dfasdfasdf")

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	num := command.SliceIntInsert(s, 22, 0)
	fmt.Printf("%T, %v::::::\n", num, num)

	if err.GetError() == model.ErrorNoRows {
		return &response.Response{
			Code: 1000,
			Type: 1,
		}
	} else if err != nil {
		return &response.Response{
			Code: 1006,
			Type: 1,
		}
	}

	return &response.Response{
		Data: bag.Data,
	}
}

//PostHandler 在Handler之后执行
func (c *TestHandler) PostHandler(req *request.Request) {
	fmt.Println("test.PostHandler+++++++++++++++")
}
