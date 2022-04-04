package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type NideshopGoodsAttribute struct {
	Id      int `orm:"column(id)" json:"id"`
	GoodsId int `orm:"column(goods_id)" json:"goods_id"`
	//AttributeId int    `orm:"column(attribute_id)" json:"attribute_id"`
	Attribute *NideshopAttribute `orm:"rel(fk)"`
	Value     string             `orm:"column(value)" json:"value"`
}

func (*NideshopGoodsAttribute) TableName() string {
	return "nideshop_goods_attribute"
}

func init() {
	orm.RegisterModel(new(NideshopGoodsAttribute))
}
