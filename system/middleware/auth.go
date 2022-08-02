package middleware

import (
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/exception"
	"github.com/duxphp/duxgo/util/function"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

// Operate 操作记录
func Operate(Type string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == "GET" {
				return next(c)
			}
			err := core.Db.Model(&model.VisitorOperate{}).Create(&model.VisitorOperate{
				Type:   Type,
				UserId: cast.ToUint(c.Get("authID")),
				Url:    c.Request().RequestURI,
				Method: c.Request().Method,
				Params: cast.ToString(function.CtxBody(c)),
			}).Error
			if err != nil {
				exception.Error(err)
			}
			return next(c)
		}
	}
}
