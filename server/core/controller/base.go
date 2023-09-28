package controller

import (
	"github.com/tonly18/xerror"
	"server/core/request"
	"server/core/response"
)

type BaseHandle struct{}

func (c *BaseHandle) PreHandler(req *request.Request) {}
func (c *BaseHandle) Handler(req *request.Request) (*response.Response, xerror.Error) {
	return &response.Response{}, nil
}
func (c *BaseHandle) PostHandler(req *request.Request) {}
