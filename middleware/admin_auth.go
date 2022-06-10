package middleware

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"shopadmin/services"
	"shopadmin/utils"
	"strings"
)

var adminExceptAuth = map[string]int{
	"admin/auth/login": 1,
}

var AdminUserId int

func AdminAuth() {
	AdminUserId = 0
	var filterFunc = func(ctx *context.Context) {
		requestUrl := parseUrl(ctx.Input.URL())
		//判断路由是否要验证
		if utils.KeyInMapStringInt(adminExceptAuth, requestUrl) {
			return
		}
		//获取token
		var token string
		token = ctx.Input.Header("Authorization")
		tokenSplit := strings.Fields(token)
		if len(tokenSplit) == 0 || tokenSplit[1] == "" {
			utils.HttpFail(401, "token 不存在", ctx)
		}
		//解析token
		jwtData, err := services.ParseToken(tokenSplit[1])
		if err != nil {
			utils.HttpFail(401, "token 过期", ctx)
		}
		//当前时间戳
		var unixTime int64
		unixTime = utils.GetNowTimestamp()

		if unixTime > jwtData.ExpiresAt {
			utils.HttpFail(401, "token 过期", ctx)
		}
		//刷新新token返回前端
		if unixTime-jwtData.ExpiresAt < 30 {
			newToken, _ := services.RefreshToken(tokenSplit[1])
			ctx.Output.Header("New-Authorization-Token", newToken)
		}
		AdminUserId = jwtData.UserID
		return
	}
	beego.InsertFilter("/admin/*", beego.BeforeRouter, filterFunc)
}
