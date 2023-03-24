package rest

import (
	"net/http"
	"workPackage/controller"
	"workPackage/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/**
 * RunServer runs the REST server
 */
func RunServer(address string) error {

	// 컨트롤러 초기화
	c := controller.NewController()

	// 서버 핸들러 실행
	return RunServerHandler(address, c)

}

func RunServerHandler(address string, c controller.ControllerService) error {

	// Create a new echo instance
	e := echo.New()

	// 미들웨어 세팅
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middlewares.LogrusLogger())

	// Cors 설정
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
	}))

	// 라우터 세팅
	api := e.Group("/api/v1")
	api.GET("/hello", c.Hello)

	// 서버 실행
	return e.Start(address)
}
