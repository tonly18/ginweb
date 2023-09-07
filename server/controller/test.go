package controller

import (
	"server/core/controller"
	"server/core/request"
	"server/core/response"
	"server/core/xerror"
	"server/service"
	"server/service/model"
)

//TestHandler Test测试接口
type TestHandler struct {
	controller.BaseHandle
}

//PreHandler 在Handler之前执行
func (c *TestHandler) PreHandler(req *request.Request) {
	//fmt.Println("test.PerHandler - 1111111111111111")

	//raw := map[int]int{1: 11, 4: 44, 3: 33, 2: 22, 5: 55, 6: 66}
	//raw := []string{"1", "2", "3", "4", "5", "6", "abcdf", "abc"}
	//fmt.Printf("raw-0::: %T %v %p\n", raw, raw, raw)
	//
	//ret := command.SliceContains(raw, "abc")
	//
	//fmt.Printf("raw-1::: %T %v %p\n", raw, raw, &ret)
	//fmt.Printf("ret::::: %T %v %v\n", ret, ret, ret)

	//fmt.Println("StringGenRandom::::", string(command.StringGenRandom(6, []byte("asdfasdfasdf")...)))
}

//Handler 业务处理
func (c *TestHandler) Handler(req *request.Request) (*response.Response, xerror.Error) {
	testService := service.NewTestService(req)
	data, err := testService.Query(1, 4)
	//data, err := testService.QueryMap(1, 2)
	if err != nil {
		if err.Is(model.ErrorNoRows) {
			return nil, xerror.Wrap(err, &xerror.NewError{
				Code:    500000010,
				Err:     err.GetErr(),
				Message: "test bag.query",
			})
		}
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

//PostHandler 在Handler之后执行
func (c *TestHandler) PostHandler(req *request.Request) {
	//fmt.Println("test.PostHandler - 333333333333333")
}
