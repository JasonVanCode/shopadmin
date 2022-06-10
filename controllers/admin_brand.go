package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
)

type AdminBrandController struct {
	AdminBaseController
}

//后台品牌列表
func (brand *AdminBrandController) GetIndex() {
	var brandService services.BrandService
	pageData, err := brandService.GetBrandLists(brand.pagination, brand.otherQueryData)
	if err != nil {
		brand.ErrLog(err)
		utils.HttpFail(400, "数据获取失败", brand.Ctx)
		return

	}
	utils.HttpSuccess(pageData, brand.Ctx)
}

//添加/编辑品牌
func (brand *AdminBrandController) CreateOrUpdate() {
	var userId = brand.GetUserId()
	var requestBody services.BrandSaveBody
	if err := json.Unmarshal(brand.Ctx.Input.RequestBody, &requestBody); err != nil {
		brand.ErrLog(err)
		utils.HttpFail(400, err.Error(), brand.Ctx)
	}
	v := validation.Validation{}
	if b, err := v.Valid(&requestBody); err != nil || !b {
		brand.ErrLog(err)
		utils.HttpFail(400, "必填参数未传", brand.Ctx)
		return
	}
	var brandService services.BrandService
	var brandInfo = new(models.NideshopBrand)
	var errorMsg error
	if requestBody.Id > 0 {
		brandInfo, errorMsg = brandService.UpdateBrand(requestBody, userId)
	} else {
		brandInfo, errorMsg = brandService.CreateBrand(requestBody, userId)
	}
	if errorMsg != nil {
		brand.ErrLog(errorMsg)
		utils.HttpFail(400, errorMsg.Error(), brand.Ctx)
	}

	utils.HttpSuccess(brandInfo, brand.Ctx)
}

//获取品牌数据
func (brand *AdminBrandController) GetEdit() {
	id, _ := brand.GetInt("id")
	var brandService services.BrandService
	data, err := brandService.GetBrandById(id)
	if err != nil {
		brand.ErrLog(err)
		utils.HttpFail(400, "暂无该条数据", brand.Ctx)
		return
	}
	utils.HttpSuccess(data, brand.Ctx)
}

//获取所有商品列表
func (brand *AdminBrandController) GetBrandIdNames() {
	var brandService services.BrandService
	data := brandService.GetAllBrandIdNames()
	utils.HttpSuccess(data, brand.Ctx)
}
