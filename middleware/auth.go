package middleware

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"shopadmin/services"
	"shopadmin/utils"
	"strings"
	"time"
)

//不需要验证的路由
var authExcept = map[string]int{
	"api/auth/loginByWeixin": 1,
	//"api/address/save":       2,
	"api/cart/add": 3,
}

var UserId int

//请求的登录验证
func Auth() {
	var filterFunc = func(ctx *context.Context) {
		requestUrl := parseUrl(ctx.Input.URL())
		if utils.KeyInMapStringInt(authExcept, requestUrl) {
			return
		}
		token := ctx.Input.Header("x-nideshop-token")
		if token == "" {
			utils.HttpFail(401, "token not found", ctx)
		}
		//解析token，获取到用户的信息
		info, err := services.ParseToken(token)
		if err != nil || info.UserID == 0 {
			utils.HttpFail(401, err.Error(), ctx)
		}
		UserId = info.UserID

		//判断token是否快过期
		timeNowUnix := time.Now().Unix()
		expiresAt := info.ExpiresAt
		if timeNowUnix-expiresAt < 30 {
			var newToken string
			if newToken, err = services.RefreshToken(token); err != nil {
				logs.Error("generate new token fail")
				return
			}
			//将新生成的token返回给前端
			ctx.Output.Header("new-x-nideshop-token", newToken)
		}

	}
	beego.InsertFilter("/api/*", beego.BeforeRouter, filterFunc)
}

//重新拼接url
func parseUrl(url string) string {
	//urlArr := strings.Split(url, "/")
	urlArr := utils.RemoveEmpSliceString(strings.Split(url, "/"))
	if urlArr == nil {
		return ""
	}
	return strings.Join(urlArr, "/")
}
