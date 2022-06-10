package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
)

type AdminUserController struct {
	AdminBaseController
}

//用户列表
func (user *AdminUserController) GetIndex() {
	var userService services.AdminUserService
	pageData, err := userService.GetUserLists(user.pagination, user.otherQueryData)
	if err != nil {
		utils.HttpFail(400, err.Error(), user.Ctx)
		return
	}
	utils.HttpSuccess(pageData, user.Ctx)
}

//新增/编辑 用户
func (user *AdminUserController) CreateOrUpdate() {
	var requestBody services.UserRequestBody
	if err := json.Unmarshal(user.Ctx.Input.RequestBody, &requestBody); err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, err.Error(), user.Ctx)
		return
	}
	v := validation.Validation{}
	if b, err := v.Valid(&requestBody); err != nil || !b {
		logs.Error(err.Error())
		utils.HttpFail(400, "必填参数未传", user.Ctx)
		return
	}
	if requestBody.Password != requestBody.RePassword {
		utils.HttpFail(400, "密码不一致", user.Ctx)
		return
	}
	var userService services.AdminUserService
	var userInfo *models.AdminUser
	var errorMsg error
	if requestBody.Id > 0 {
		userInfo, errorMsg = userService.UpdateUser(requestBody)
	} else {
		userInfo, errorMsg = userService.AddUser(requestBody)
	}
	if errorMsg != nil {
		logs.Error(errorMsg.Error())
		utils.HttpFail(400, errorMsg.Error(), user.Ctx)
		return
	}
	utils.HttpSuccess(userInfo, user.Ctx)
}

//删除单个、多个用户数据
func (user *AdminUserController) DelUser() {
	var delBody services.UserDelBody
	if err := json.Unmarshal(user.Ctx.Input.RequestBody, &delBody); err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, err.Error(), user.Ctx)
		return
	}
	if delBody.Id == "" {
		utils.HttpFail(400, "暂无该数据", user.Ctx)
		return
	}
	var userService services.AdminUserService
	b, err := userService.DelUser(delBody.Id)
	if err != nil || !b {
		utils.HttpFail(400, "删除失败", user.Ctx)
		return
	}
	utils.HttpSuccess(nil, user.Ctx)
}

//修改状态
func (user *AdminUserController) ChangeStatus() {
	var statusBody services.UserStatusBody
	if err := json.Unmarshal(user.Ctx.Input.RequestBody, &statusBody); err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, err.Error(), user.Ctx)
		return
	}
	var userService services.AdminUserService
	b, err := userService.ChangeStatus(statusBody)
	if err != nil || !b {
		utils.HttpFail(400, "修改失败", user.Ctx)
		return
	}
	utils.HttpSuccess(nil, user.Ctx)
}

//添加角色
func (user *AdminUserController) ChangeRole() {
	var roleBody services.UserRoleBody
	if err := json.Unmarshal(user.Ctx.Input.RequestBody, &roleBody); err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, err.Error(), user.Ctx)
		return
	}
	var userService services.AdminUserService
	b, err := userService.ChangeUserRole(roleBody)
	if err != nil || !b {
		utils.HttpFail(400, err.Error(), user.Ctx)
		return
	}
	utils.HttpSuccess(nil, user.Ctx)
}
