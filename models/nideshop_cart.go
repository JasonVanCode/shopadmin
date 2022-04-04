package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type NideshopCart struct {
	Id                        int     `orm:"column(id)" json:"id"`
	UserId                    int     `orm:"column(user_id)" json:"user_id"`
	SessionId                 string  `orm:"column(session_id)" json:"session_id"`
	GoodsId                   int     `orm:"column(goods_id)" json:"goods_id"`
	GoodsSn                   string  `orm:"column(goods_sn)" json:"goods_sn"`
	ProductId                 int     `orm:"column(product_id)" json:"product_id"`
	GoodsName                 string  `orm:"column(goods_name)" json:"goods_name"`
	MarketPrice               float64 `orm:"column(market_price);digits(10);decimals(2);description(市场价格)" json:"market_price"`
	RetailPrice               float64 `orm:"column(retail_price);digits(10);decimals(2)" json:"retail_price"`
	Number                    int     `orm:"column(number)" json:"number"`
	GoodsSpecifitionNameValue string  `orm:"column(goods_specifition_name_value)" json:"goods_specifition_name_value"`
	GoodsSpecifitionIds       string  `orm:"column(goods_specifition_ids)" json:"goods_specifition_ids"`
	Checked                   int     `orm:"column(checked);default(1)" json:"checked"`
	ListPicUrl                string  `orm:"column(list_pic_url)" json:"list_pic_url"`
}

func (*NideshopCart) TableName() string {
	return "nideshop_cart"
}

func (c *NideshopCart) GetGoodsUrl() string {
	var goods NideshopGoods
	o := orm.NewOrm()
	err := o.QueryTable(new(NideshopGoods)).Filter("id", c.GoodsId).One(&goods)
	if err != nil || goods.Id != c.GoodsId {
		return ""
	}
	return goods.ListPicUrl
}

func init() {
	orm.RegisterModel(new(NideshopCart))
}
