package controllers

import (
	"shopadmin/services"
	"shopadmin/utils"
)

type RegionController struct {
	BaseController
}

//获取区域数据
func (r *RegionController) GetRegionList() {
	parentId, _ := r.GetInt("parentId")
	var regionService = new(services.RegionService)
	resData := regionService.GetRegionByParentId(parentId)

	utils.HttpSuccess(&resData, r.Ctx)
}
