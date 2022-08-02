package admin

import (
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/exception"
	"github.com/duxphp/duxgo/middleware"
	"github.com/duxphp/duxgo/response"
	"github.com/duxphp/duxgo/util"
	"github.com/duxphp/duxgo/util/function"
	"github.com/labstack/echo/v4"
)

// Register 用户注册
func Register(ctx echo.Context) error {
	var params struct {
		Username string `json:"username" validate:"required" validateMsg:"请输入账号"`
		Password string `json:"password" validate:"required" validateMsg:"请输入密码"`
	}
	if err := util.RequestParser(ctx, &params); err != nil {
		return err
	}

	hash := function.HashEncode([]byte(params.Password))
	user := model.SystemUser{Nickname: "管理员", Username: params.Username, Password: hash}
	if err := core.Db.Create(&user).Error; err != nil {
		return err
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
