package services

import (
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
)

type ProductService struct {
	BaseService
}

//根据id获取分类信息
func (c *ProductService) GetProductById(id int) *models.NideshopProduct {
	var product models.NideshopProduct
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopProduct)).Filter("id", id).One(&product)
	if err != nil || product.Id != id {
		return nil
	}
	return &product
}
