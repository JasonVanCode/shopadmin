package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/services"
	"shopadmin/utils"
)

type AdminGoodsController struct {
	AdminBaseController
}

//获取列表
func (goods *AdminGoodsController) GetIndex() {
	var goodsService services.GoodsService
	data, err := goodsService.GetAdminGoodsIndex(goods.pagination, goods.otherQueryData)
	if err != nil {
		goods.ErrLog(err)
		utils.HttpFail(400, "数据获取失败", goods.Ctx)
		return
	}
	//品牌数据
	var brandService services.BrandService
	brandLists := brandService.GetAllBrandIdNames()
	//分类数据
	var catService services.CategoryService
	catLists := catService.GetAllCategoryIdNames()
	resultdata := struct {
		Data       *services.PageData `json:"data"`
		BrandLists []orm.Params       `json:"brand_lists"`
		CatLists   []orm.Params       `json:"cat_lists"`
	}{
		data,
		brandLists,
		catLists,
	}
	utils.HttpSuccess(resultdata, goods.Ctx)
}

//修改商品一些状态
func (goods *AdminGoodsController) ChangeGoodsStatus() {
	var body services.GoodsStatusBody
	if err := json.Unmarshal(goods.Ctx.Input.RequestBody, &body); err != nil {
		goods.ErrLog(err)
		utils.HttpFail(400, "数据解析失败", goods.Ctx)
		return
	}
	var goodsService services.GoodsService
	if err := goodsService.HandleGoodsStatus(body); err != nil {
		goods.ErrLog(err)
		utils.HttpFail(400, err.Error(), goods.Ctx)
		return
	}
	utils.HttpSuccess(nil, goods.Ctx)
}
