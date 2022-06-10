package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/services"
	"shopadmin/utils"
)

type AdminAuthController struct {
	AdminBaseController
}

//登录验证
func (a *AdminAuthController) Login() {
	var loginBody services.LoginBody
	if err := json.Unmarshal(a.Ctx.Input.RequestBody, &loginBody); err != nil {
		utils.HttpFail(400, err.Error(), a.Ctx)
	}
	//请求数据校验
	valid := validation.Validation{}
	if b, err := valid.Valid(&loginBody); err != nil || !b {
		utils.HttpFail(400, "请填写用户名和密码", a.Ctx)
	}
	var adminAuthAervice services.AdminAuthService
	userInfo, err := adminAuthAervice.Login(loginBody)
	if err != nil {
		utils.HttpFail(400, err.Error(), a.Ctx)
	}
	//生成jwt token
	token, err := services.GenerateToken(userInfo.Id, userInfo.Username, 0)
	if err != nil {
		utils.HttpFail(400, err.Error(), a.Ctx)
	}
	utils.HttpSuccess(struct {
		UserName string `json:"userName"`
		Token    string `json:"token"`
	}{
		userInfo.Username,
		token,
	}, a.Ctx)
}
