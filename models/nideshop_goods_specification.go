package models

import (
	"github.com/beego/beego/v2/client/orm"
)

//商品对应规格表值表
type NideshopGoodsSpecification struct {
	Id      int `orm:"column(id)" json:"id"`
	GoodsId int `orm:"column(goods_id)" json:"goods_id"`
	//SpecificationId int                    `orm:"column(specification_id)" json:"specification_id"`
	Value         string                 `orm:"column(value)" json:"value"`
	PicUrl        string                 `orm:"column(pic_url)" json:"pic_url"`
	Specification *NideshopSpecification `orm:"rel(fk)"`
}

func (*NideshopGoodsSpecification) TableName() string {
	return "nideshop_goods_specification"
}

func init() {
	orm.RegisterModel(new(NideshopGoodsSpecification))
}
