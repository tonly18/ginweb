package iface

import (
	"github.com/tonly18/xerror"
	"server/core/request"
	"server/core/response"
)

type IWrapperHandler interface {
	PreHandler(*request.Request)
	Handler(*request.Request) (*response.Response, xerror.Error)
	PostHandler(*request.Request)
}
