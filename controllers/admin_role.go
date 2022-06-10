package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
)

type AdminRoleController struct {
	AdminBaseController
}

//角色列表
func (role *AdminRoleController) GetIndex() {
	var roleService services.AdminRoleService
	pageData, err := roleService.GetRoleLists(role.pagination, role.otherQueryData)
	if err != nil {
		utils.HttpFail(400, err.Error(), role.Ctx)
	}
	utils.HttpSuccess(pageData, role.Ctx)
}

//角色列表 只有id 和 name字段
func (role *AdminRoleController) GetRoleListIdName() {
	var roleService services.AdminRoleService
	utils.HttpSuccess(roleService.GetRoleIdNames(), role.Ctx)
}

//新增/编辑 角色
func (role *AdminRoleController) CreateOrUpdate() {
	var requestBody services.RoleRequestBody
	if err := json.Unmarshal(role.Ctx.Input.RequestBody, &requestBody); err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, err.Error(), role.Ctx)
	}
	v := validation.Validation{}
	if b, err := v.Valid(&requestBody); err != nil || !b {
		logs.Error(err.Error())
		utils.HttpFail(400, "必填参数未传", role.Ctx)
	}
	var roleService services.AdminRoleService
	var roleInfo *models.AdminRole
	var errorMsg error
	if requestBody.Id > 0 {
		roleInfo, errorMsg = roleService.UpdateRole(requestBody)
	} else {
		roleInfo, errorMsg = roleService.AddRole(requestBody)
	}
	if errorMsg != nil {
		logs.Error(errorMsg.Error())
		utils.HttpFail(400, errorMsg.Error(), role.Ctx)
	}
	utils.HttpSuccess(roleInfo, role.Ctx)
}

//修改状态
func (role *AdminRoleController) ChangeStatus() {
	var statusBody services.RoleStatusBody
	if err := json.Unmarshal(role.Ctx.Input.RequestBody, &statusBody); err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, err.Error(), role.Ctx)
	}
	var roleService services.AdminRoleService
	b, err := roleService.ChangeStatus(statusBody)
	if err != nil || !b {
		utils.HttpFail(400, "修改失败", role.Ctx)
	}
	utils.HttpSuccess(nil, role.Ctx)
}

//删除单个、多个用户数据
func (role *AdminRoleController) DelRole() {
	var delBody services.RoleDelBody
	if err := json.Unmarshal(role.Ctx.Input.RequestBody, &delBody); err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, err.Error(), role.Ctx)
		return
	}
	if delBody.Id == "" {
		utils.HttpFail(400, "暂无该数据", role.Ctx)
	}
	var roleService services.AdminRoleService
	b, err := roleService.DelRole(delBody.Id)
	if err != nil || !b {
		utils.HttpFail(400, "删除失败", role.Ctx)
	}
	utils.HttpSuccess(nil, role.Ctx)
}

//获取当前角色信息
func (role *AdminRoleController) GetRoleInfo() {
	id, err := role.GetInt("id")
	if err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, "请求参数有误", role.Ctx)
		return
	}
	var roleService services.AdminRoleService
	roleInfo, err := roleService.GetRoleById(id)
	if err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, "获取数据失败", role.Ctx)
		return
	}
	utils.HttpSuccess(roleInfo, role.Ctx)
}

//角色菜单分配
func (role *AdminRoleController) AllocMenus() {
	var body services.RoleMenusBody
	if err := json.Unmarshal(role.Ctx.Input.RequestBody, &body); err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, "参数解析错误", role.Ctx)
		return
	}
	if body.RoleId == 0 || body.Ids == "" {
		utils.HttpFail(400, "请选择要保存的菜单", role.Ctx)
		return
	}
	var roleService services.AdminRoleService
	b, err := roleService.SaveMenus(body)
	if err != nil || !b {
		logs.Error(err.Error())
		utils.HttpFail(400, "修改失败", role.Ctx)
		return
	}
	utils.HttpSuccess(nil, role.Ctx)
}
