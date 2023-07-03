package iface

import (
	"server/core/request"
	"server/core/response"
	"server/core/xerror"
)

type IWrapperHandler interface {
	PreHandler(*request.Request)
	Handler(*request.Request) (*response.Response, xerror.Error)
	PostHandler(*request.Request)
}
