package controllers

import (
	"fmt"
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
)

type CategoryController struct {
	BaseController
}

type CurCategory struct {
	*models.NideshopCategory
	SubCategoryList []*models.NideshopCategory `json:"subCategoryList"`
}

type CateLogIndexRtnJson struct {
	CategoryList    []*models.NideshopCategory `json:"categoryList"`
	CurrentCategory *CurCategory               `json:"currentCategory"`
}

//获取分类首页的数据
func (c *CategoryController) GetCategoryList() {
	categoryId := c.GetString("id")
	fmt.Println(categoryId)
	var categoryService = new(services.CategoryService)

	//获取顶级分类的数据
	var categories = categoryService.GetCategoryByParentId(0)
	//当前分类数据
	var cuurentCategory *models.NideshopCategory
	if categoryId != "" {
		cuurentCategory = categoryService.GetCategoryById(utils.TransStringToInt(categoryId))
	}
	//如果没有选中制定分类，默认取其中一个
	if cuurentCategory == nil {
		cuurentCategory = categories[0]
	}
	//处理获取当前分类下面的数据
	curCategory := new(CurCategory)

	if cuurentCategory != nil && cuurentCategory.Id > 0 {
		curCategory.NideshopCategory = cuurentCategory
		curCategory.SubCategoryList = categoryService.GetCategoryByParentId(cuurentCategory.Id)
	}
	utils.HttpSuccess(CateLogIndexRtnJson{
		categories,
		curCategory,
	}, c.Ctx)
}
