package controller

import (
	"errors"
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
}

//Handler 业务处理
func (c *TestHandler) Handler(req *request.Request) (*response.Response, xerror.Error) {
	testService := service.NewTestService(req)
	//data, err := testService.Query(1, 4)
	//data, err := testService.QueryMap(1, 2)
	data, err := testService.Get(1, 2)
	//data, err := testService.GetMap(1, 2)
	if err != nil {
		if err.Is(model.ErrorNoRows) {
			return nil, xerror.Wrap(req, err, &xerror.TempError{
				Code:    500000010,
				Err:     errors.New("test handler bag.query"),
				Message: "test bag.query",
				Type:    1,
			})
		}
		return nil, xerror.Wrap(req, err, &xerror.TempError{
			Code:    500000011,
			Err:     err.GetErr(),
			Message: "test handler bag.query",
			Type:    1,
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
