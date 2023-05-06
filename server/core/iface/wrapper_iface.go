package iface

import (
	"server/core/request"
	"server/core/response"
)

type IWrapperHandler interface {
	PreHandler(*request.Request)
	Handler(*request.Request) *response.Response
	PostHandler(*request.Request)
}
