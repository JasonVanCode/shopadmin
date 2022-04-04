package services

import (
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
)

type CategoryService struct {
	BaseService
	CategoryRtnJson
}

//商品分类数据返回
type CategoryRtnJson struct {
	CurCategory     *models.NideshopCategory   `json:"currentCategory"`
	ParentCategory  *models.NideshopCategory   `json:"parentCategory"`
	BrotherCategory []*models.NideshopCategory `json:"brotherCategory"`
}

//根据id获取分类信息
func (c *CategoryService) GetCategoryById(id int) *models.NideshopCategory {
	var category models.NideshopCategory
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopCategory)).Filter("id", id).One(&category)
	if err != nil {
		return nil
	}
	return &category
}

//根据parent_id 获取分类信息
func (c *CategoryService) GetCategoryByParentId(pid int) []*models.NideshopCategory {
	var categoryList []*models.NideshopCategory
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopCategory)).Filter("parent_id", pid).All(&categoryList)
	if err != nil {
		return nil
	}
	return categoryList
}
