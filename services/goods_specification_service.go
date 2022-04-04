package services

import (
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
	"strings"
)

type GoodsSpecificationService struct {
	BaseService
}

//查询该商品对应的规格数据
func (*GoodsSpecificationService) GetGoodsSpecification(goodsId int, specificationIds []string) []*models.NideshopGoodsSpecification {

	var specifications []*models.NideshopGoodsSpecification
	o := orm.NewOrm()
	nums, err := o.QueryTable(new(models.NideshopGoodsSpecification)).Filter("goods_id", goodsId).Filter("id__in", specificationIds).All(&specifications)
	if err != nil || nums == 0 {
		return nil
	}
	return specifications
}

//拼接规格数据
func (*GoodsSpecificationService) HandleSpecificationValues(data []*models.NideshopGoodsSpecification) string {
	var strs []string
	for _, v := range data {
		strs = append(strs, v.Value)
	}
	if len(strs) == 0 {
		return ""
	}
	return strings.Join(strs, ";")
}
