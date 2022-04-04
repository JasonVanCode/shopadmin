package models

import (
	"github.com/beego/beego/v2/client/orm"
)

//商品对应规格表值表
type NideshopOrderGoods struct {
	Id                        int     `orm:"column(id)" json:"id"`
	OrderId                   int     `orm:"column(order_id)" json:"order_id"`
	GoodsId                   int     `orm:"column(goods_id)" json:"goods_id"`
	GoodsName                 string  `orm:"column(goods_name)" json:"goods_name"`
	GoodsSn                   string  `orm:"column(goods_sn)" json:"goods_sn"`
	ProductId                 int     `orm:"column(product_id)" json:"product_id"`
	Number                    int     `orm:"column(number)" json:"number"`
	MarketPrice               float64 `orm:"column(market_price);digits(10);decimals(2)" json:"market_price"`
	RetailPrice               float64 `orm:"column(retail_price);digits(10);decimals(2)" json:"retail_price"`
	GoodsSpecifitionNameValue string  `orm:"column(goods_specifition_name_value)" json:"goods_specifition_name_value"`
	IsReal                    int     `orm:"column(is_real)" json:"is_real"`
	GoodsSpecifitionIds       string  `orm:"column(goods_specifition_ids)" json:"goods_specifition_ids"`
	ListPicUrl                string  `orm:"column(list_pic_url)" json:"list_pic_url"`
}

func (*NideshopOrderGoods) TableName() string {
	return "nideshop_order_goods"
}

func init() {
	orm.RegisterModel(new(NideshopOrderGoods))
}
