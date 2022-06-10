package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
)

type AdminMenuController struct {
	AdminBaseController
}

//菜单首页数据
func (menu *AdminMenuController) GetIndex() {
	var menuService services.AdminMenuService
	pageData, err := menuService.GetMenuLists(menu.pagination, menu.otherQueryData)
	if err != nil {
		utils.HttpFail(400, err.Error(), menu.Ctx)
		return
	}
	utils.HttpSuccess(pageData, menu.Ctx)
}

//菜单列表 只有id 和 name字段 且只是顶级菜单
func (menu *AdminMenuController) GetMenuListIdName() {
	var menuService services.AdminMenuService
	utils.HttpSuccess(menuService.GetMenuIdNames(), menu.Ctx)
}

//新增或者编辑菜单数据
func (menu *AdminMenuController) CreateOrUpdate() {
	var requestBody services.MenuSaveBody
	if err := json.Unmarshal(menu.Ctx.Input.RequestBody, &requestBody); err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, err.Error(), menu.Ctx)
		return
	}
	v := validation.Validation{}
	if b, err := v.Valid(&requestBody); err != nil || !b {
		logs.Error(err.Error())
		utils.HttpFail(400, "必填参数未传", menu.Ctx)
		return
	}
	var menuService services.AdminMenuService
	var menuInfo *models.AdminMenu
	var errorMsg error
	//编辑
	if requestBody.Id > 0 {
		menuInfo, errorMsg = menuService.UpdateMenu(requestBody)
	} else {
		menuInfo, errorMsg = menuService.AddMenu(requestBody)
	}
	if errorMsg != nil {
		logs.Error(errorMsg.Error())
		utils.HttpFail(400, "数据添加失败", menu.Ctx)
		return
	}
	utils.HttpSuccess(menuInfo, menu.Ctx)
}

//删除单个、多个菜单
func (menu *AdminMenuController) DelMenu() {
	var delBody services.MenuDelBody
	if err := json.Unmarshal(menu.Ctx.Input.RequestBody, &delBody); err != nil {
		logs.Error(err.Error())
		utils.HttpFail(400, "参数解析错误", menu.Ctx)
		return
	}
	var menuService services.AdminMenuService
	b, err := menuService.DelMenu(delBody)
	if err != nil || !b {
		utils.HttpFail(400, "删除失败", menu.Ctx)
		return
	}
	utils.HttpSuccess(nil, menu.Ctx)
}

//分配菜单列表数据
func (menu *AdminMenuController) GetAllocMenuLists() {
	var menuService services.AdminMenuService
	allLists := menuService.GetAllShowMenulis()
	if len(allLists) == 0 {
		utils.HttpSuccess(nil, menu.Ctx)
		return
	}
	data := menuService.GetTreeMenuLists(0, allLists)
	utils.HttpSuccess(data, menu.Ctx)
}
