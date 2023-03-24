package controller

import "workPackage/web"

type Controller struct {
	web.Response
}

// REST API 컨트롤러 초기화
func NewController() *Controller {
	return &Controller{web.Response{}}
}
