package controllers

import "shopadmin/utils"

type AdminIndexController struct {
	AdminBaseController
}

func (a *AdminIndexController) GetIndex() {
	utils.HttpFail(400, "测试", a.Ctx)
	a.Ctx.WriteString("aaaaa")
}
