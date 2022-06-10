package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
)

type AdminCategoryController struct {
	AdminBaseController
}

//分类列表数据
func (category *AdminCategoryController) GetIndex() {
	var categoryService services.CategoryService
	pageData, err := categoryService.GetAdminCategoryLists(category.pagination, category.otherQueryData)
	if err != nil {
		category.ErrLog(err)
		utils.HttpFail(400, "数据获取失败", category.Ctx)
		return
	}
	utils.HttpSuccess(pageData, category.Ctx)
}

//获取分类下的属性
func (category *AdminCategoryController) GetCategoryAttributeLists() {
	cateId, err := category.GetInt("cateId")
	name := category.GetString("name")
	if err != nil {
		category.ErrLog(err)
		utils.HttpFail(400, "参数有误", category.Ctx)
		return
	}
	var categoryService services.CategoryService
	data, err := categoryService.GetAttributeByCateId(cateId, name)
	if err != nil {
		category.ErrLog(err)
		utils.HttpFail(400, "暂无数据", category.Ctx)
		return
	}
	utils.HttpSuccess(data, category.Ctx)
}

//获取顶级分类的id 和 name
func (category *AdminCategoryController) GetCategoryIdName() {
	var categoryService services.CategoryService
	data := categoryService.GetTopCategoryIdName()
	utils.HttpSuccess(data, category.Ctx)
}

//分类新增或者编辑
func (category *AdminCategoryController) CreateOrUpdate() {
	userId := category.GetUserId()
	var body services.CategoryRequestBody
	if err := json.Unmarshal(category.Ctx.Input.RequestBody, &body); err != nil {
		category.ErrLog(err)
		utils.HttpFail(400, "请求参数解析失败", category.Ctx)
		return
	}
	v := validation.Validation{}
	if b, err := v.Valid(&body); err != nil || !b {
		fmt.Println("err---->", err, "b---->", b)
		utils.HttpFail(400, "必填项未提交", category.Ctx)
		return
	}
	var categoryService services.CategoryService
	var categoryInfo = new(models.NideshopCategory)
	var errMgs error
	if body.Id > 0 {
		categoryInfo, errMgs = categoryService.UpdateCategory(body, userId)
	} else {
		categoryInfo, errMgs = categoryService.CreateCategory(body, userId)
	}
	if errMgs != nil {
		utils.HttpFail(400, errMgs.Error(), category.Ctx)
		return
	}
	utils.HttpSuccess(categoryInfo, category.Ctx)
}

//获取分类数据
func (category *AdminCategoryController) GetEdit() {
	id, _ := category.GetInt("id")
	var categoryService services.CategoryService
	data := categoryService.GetCategoryById(id)
	utils.HttpSuccess(data, category.Ctx)
}

//获取分类以及子分类数据
func (category *AdminCategoryController) GetCategoryWithChildren() {
	var categoryService services.CategoryService
	data, err := categoryService.GetCategoryWithChildren()
	if err != nil {
		category.ErrLog(err)
		utils.HttpSuccess(nil, category.Ctx)
		return
	}
	utils.HttpSuccess(data, category.Ctx)
}

//删除分类
func (category *AdminCategoryController) DelCategory() {
	id, _ := category.GetInt("id")
	var categoryService services.CategoryService
	b, err := categoryService.DelCategory(id)
	if err != nil || !b {
		category.ErrLog(err)
		utils.HttpFail(400, "删除失败", category.Ctx)
	}
	utils.HttpSuccess(nil, category.Ctx)
}

//分类下面属性的新增或者编辑
func (category *AdminCategoryController) AttributeCreateOrUpdate() {
	var body services.RequestAttributeBody
	if err := json.Unmarshal(category.Ctx.Input.RequestBody, &body); err != nil {
		category.ErrLog(err)
		utils.HttpFail(400, "数据解析失败", category.Ctx)
		return
	}
	var attrService services.AttributeService
	var attInfo = new(models.NideshopAttribute)
	var errMsg error
	if body.Id > 0 {
		attInfo, errMsg = attrService.UpdateAttribute(body)
	} else {
		attInfo, errMsg = attrService.CreateAttribute(body)
	}
	if errMsg != nil {
		category.ErrLog(errMsg)
		utils.HttpFail(400, errMsg.Error(), category.Ctx)
		return
	}
	utils.HttpSuccess(attInfo, category.Ctx)
}

//删除属性
func (category *AdminCategoryController) DelAttribute() {
	var body services.AttributeDelBody
	if err := json.Unmarshal(category.Ctx.Input.RequestBody, &body); err != nil {
		category.ErrLog(err)
		utils.HttpFail(400, "数据解析失败", category.Ctx)
		return
	}
	if body.Id == "" {
		utils.HttpFail(400, "请提交要删除的数据", category.Ctx)
		return
	}
	var attrService services.AttributeService
	if err := attrService.DelAttribute(body.Id); err != nil {
		utils.HttpFail(400, err.Error(), category.Ctx)
	}
	utils.HttpSuccess(nil, category.Ctx)
}
