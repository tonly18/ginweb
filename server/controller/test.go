package controller

import (
	"fmt"
	"server/core/controller"
	"server/core/request"
	"server/core/response"
	"server/core/xerror"
	"server/library/command"
	"server/service"
)

// TestHandler Test测试接口
type TestHandler struct {
	controller.BaseHandle
}

// PreHandler 在Handler之前执行
func (c *TestHandler) PreHandler(req *request.Request) {
	//fmt.Println("test.PerHandler - 1111111111111111")

	raw := map[int]int{1: 11, 4: 44, 3: 33, 2: 22, 5: 55, 6: 66}
	//raw := []string{"1", "2", "3", "4", "5", "6", "abcdf", "abc"}
	//raw := []int{1, 2, 3, 4, 4, 5, 5, 6}
	fmt.Printf("raw:::::: %T %p %v\n", raw, raw, raw)

	ret := command.MapKeys[int, int](raw, 2)
	fmt.Printf("ret:::::: %T %p %v\n", ret, ret, ret)

	//fmt.Println("StringGenRandom::::", string(command.StringGenRandom(6, []byte("asdfasdfasdf")...)))
}

// Handler 业务处理
func (c *TestHandler) Handler(req *request.Request) (*response.Response, xerror.Error) {
	testService := service.NewTestService(req)
	data, err := testService.Query(1, 4)
	//data, err := testService.QueryMap(1, 2)
	//id, err := testService.Insert(1, 8, map[string]any{
	//	"uid":    8,
	//	"item":   "item-8",
	//	"expire": 1694086514,
	//})
	//id, err := testService.Modify(1, 6, map[string]any{
	//	"item":   "item-666",
	//	"expire": 1694086514,
	//})
	//fmt.Println("id::::::::::", id)
	//fmt.Println("err:::::::::", err)
	if err != nil {
		return nil, xerror.Wrap(err, &xerror.NewError{
			Code:    500000011,
			Err:     err.GetErr(),
			Message: "test handler bag.query",
		})
	}
	//for k, v := range data {
	//	fmt.Println("k-v::::::", k, v.Uid, v.Item, v.Expire, v.Itime)
	//	fmt.Println("k-v::::::", k, v)
	//}

	return &response.Response{
		Data: data,
	}, nil
}

// PostHandler 在Handler之后执行
func (c *TestHandler) PostHandler(req *request.Request) {
	//fmt.Println("test.PostHandler - 333333333333333")
}
