package middlewares

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"strings"
)

// LogrusLogger 미들웨어
// request_id, body, connect_ip, request_url, user_agent를 logrus에 저장
func LogrusLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			data := make(map[string]interface{})

			// request_id를 가져와 logEntry에 세팅
			id := c.Request().Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = c.Response().Header().Get(echo.HeaderXRequestID)
			}

			// form data를 가져와 logEntry에 세팅
			var getBodyData []string
			values := c.QueryParams()
			for k, v := range values {
				value := fmt.Sprintf("%s: %s", k, strings.Join(v, "&"))
				getBodyData = append(getBodyData, value)
			}

			// data에 저장
			data["request_id"] = id
			data["body"] = strings.Join(getBodyData, "&")
			data["connect_ip"] = c.RealIP()
			data["request_url"] = c.Request().URL.RequestURI()
			data["user_agent"] = c.Request().UserAgent()

			// data를 Context에 저장
			req := c.Request()
			c.SetRequest(req.WithContext(
				context.WithValue(
					req.Context(),
					"LOG",
					data,
				),
			))

			return next(c)
		}
	}
}
