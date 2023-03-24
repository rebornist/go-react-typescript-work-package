package controller

import "github.com/labstack/echo"

type ControllerService interface {
	Hello(echo.Context) error
}
