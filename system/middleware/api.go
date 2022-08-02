package middleware

import (
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo/util/function"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"strings"
	"time"
)

func ApiHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 计算请求超时

		headerDate := c.Request().Header.Get("Content-Date")
		timeNow := time.Now()
		date := time.Unix(cast.ToInt64(headerDate), 0)
		diff := timeNow.Sub(date).Seconds()
		if diff > 10 {
			return echo.ErrRequestTimeout
		}
		// 签名验证
		md5 := c.Request().Header.Get("Content-MD5")
		AccessKey := c.Request().Header.Get("AccessKey")
		if md5 == "" || AccessKey == "" {
			return echo.ErrUnauthorized
		}
		apiInfo := service.ServiceApi[AccessKey]
		if apiInfo == nil {
			return echo.ErrUnauthorized
		}
		url := c.Request().Host + c.Request().RequestURI
		signStr := `url=` + url + `&timestamp=` + headerDate + `&key=` + apiInfo.SecretKey

		signStr = strings.ToUpper(function.Md5(signStr))
		if signStr != md5 {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
