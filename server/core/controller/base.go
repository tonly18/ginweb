package controller

import (
	"server/core/request"
	"server/core/response"
)

type BaseHandle struct{}

func (c *BaseHandle) PreHandler(req *request.Request) {}
func (c *BaseHandle) Handler(req *request.Request) *response.Response {
	return &response.Response{}
}
func (c *BaseHandle) PostHandler(req *request.Request) {}
