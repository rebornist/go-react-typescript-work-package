package web

import (
	"github.com/labstack/echo"
	customlog "workPackage/logger"
)

type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

// 정상 응답
func SuccessResponse(c echo.Context, r Response) error {

	// 로그 값 받아오기
	logger := c.Request().Context().Value("LOG").(map[string]interface{})

	// 로그 패키지 불러오기
	log := customlog.NewLogger()

	// 로그 출력
	if err := log.PrintLogger(logger, r.Code, ""); err != nil {
		log.RaiseError(err)
	}

	// 로그 파일에 저장
	if err := log.WriteLogger(logger, r.Code, ""); err != nil {
		log.RaiseError(err)
	}

	// 응답 반환
	return c.JSON(r.Code, r)
}

// 에러 응답
func ErrorResponse(c echo.Context, r Response) error {

	// 로그 값 받아오기
	logger := c.Request().Context().Value("LOG").(map[string]interface{})

	// 로그 패키지 불러오기
	var log = customlog.NewLogger()

	// 로그 출력
	if err := log.PrintLogger(logger, r.Code, r.Message); err != nil {
		log.RaiseError(err)
	}

	// 로그 파일에 저장
	if err := log.WriteLogger(logger, r.Code, r.Message); err != nil {
		log.RaiseError(err)
	}

	// 에러 메시지 변환
	r.Message = "에러가 발생했습니다. 관리자에게 문의해주세요."

	// 응답 반환
	return c.JSON(r.Code, r.Message)
}
