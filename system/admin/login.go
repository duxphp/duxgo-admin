package admin

import (
	"errors"
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/exception"
	"github.com/duxphp/duxgo/middleware"
	"github.com/duxphp/duxgo/response"
	"github.com/duxphp/duxgo/util"
	"github.com/duxphp/duxgo/util/function"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Login 账号登录
func Login(ctx echo.Context) error {
	var params struct {
		Username string `json:"username" validate:"required" validateMsg:"请输入账号"`
		Password string `json:"password" validate:"required" validateMsg:"请输入密码"`
	}
	if err := util.RequestParser(ctx, &params); err != nil {
		return err
	}

	var user model.SystemUser
	err := core.Db.Model(&model.SystemUser{Username: params.Username}).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return exception.BusinessError("账号或密码错误")
	}
	if err != nil {
		return err
	}

	if !function.HashVerify(user.Password, []byte(params.Password)) {
		return exception.BusinessError("账号或密码错误")
	}

	token, err := middleware.NewJWT().MakeToken("admin", user.ID)
	if err != nil {
		return exception.BusinessError(err.Error())
	}

	return response.New(ctx).Send("ok", map[string]any{
		"userInfo": map[string]any{
			"user_id":     user.ID,
			"avatar":      "",
			"avatar_text": "A",
			"username":    user.Username,
			"nickname":    user.Nickname,
			"rolename":    "管理组",
		},
		"token": "Bearer " + token,
	})
}

// LoginCheck 登录检测
func LoginCheck(ctx echo.Context) error {
	var count int64
	core.Db.Model(&model.SystemUser{}).Count(&count)
	return response.New(ctx).Send("ok", map[string]any{
		"register": count <= 0,
	})
}

// LoginLogout 登录退出
func LoginLogout(ctx echo.Context) error {
	return nil
}
