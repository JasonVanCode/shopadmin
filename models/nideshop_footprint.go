package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

//商品浏览足迹
type NideshopFootprint struct {
	Id     int `orm:"column(id)" json:"id"`
	UserId int `orm:"column(user_id)" json:"user_id"`
	//GoodsId  string         `orm:"column(goods_id)" json:"goods_id"`
	AddTime time.Time      `orm:"column(add_time)" json:"add_time"`
	Goods   *NideshopGoods `orm:"rel(fk)"`
}

func (*NideshopFootprint) TableName() string {
	return "nideshop_footprint"
}

func init() {
	orm.RegisterModel(new(NideshopFootprint))
}
