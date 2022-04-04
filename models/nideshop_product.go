package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type NideshopProduct struct {
	Id                    int     `orm:"column(id)" json:"id"`
	GoodsId               int     `orm:"column(goods_id)" json:"goods_id"`
	GoodsSpecificationIds string  `orm:"column(goods_specification_ids)" json:"goods_specification_ids"`
	GoodsSn               string  `orm:"column(goods_sn)" json:"goods_sn"`
	GoodsNumber           int     `orm:"column(goods_number)" json:"goods_number"`
	RetailPrice           float64 `orm:"column(retail_price);digits(10);decimals(2);" json:"retail_price"`
}

func (*NideshopProduct) TableName() string {
	return "nideshop_product"
}

func init() {
	orm.RegisterModel(new(NideshopProduct))
}
