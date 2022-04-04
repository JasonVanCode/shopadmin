package models

import (
	"github.com/beego/beego/v2/client/orm"
)

//相关产品
type NideshopRelatedGoods struct {
	Id      int `orm:"column(id)" json:"id"`
	GoodsId int `orm:"column(goods_id)" json:"goods_id"`
	//RelatedGoodsId int            `orm:"column(related_goods_id)" json:"related_goods_id"`
	RelatedGoods *NideshopGoods `orm:"rel(fk)"`
}

func (*NideshopRelatedGoods) TableName() string {
	return "nideshop_related_goods"
}

func init() {
	orm.RegisterModel(new(NideshopRelatedGoods))
}
