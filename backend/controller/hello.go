package controller

import "github.com/labstack/echo"

func (c *Controller) Hello(ctx echo.Context) error {
	c.Response.Code = 200
	c.Response.Message = "Success"
	c.Response.Data = map[string]interface{}{"name": "Hello, Golang Echo Server!"}

	return ctx.JSON(c.Response.Code, c.Response)
}
