package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type NideshopGoodsGallery struct {
	Id        int    `orm:"column(id)" json:"id"`
	GoodsId   int    `orm:"column(goods_id)" json:"goods_id"`
	ImgUrl    string `orm:"column(img_url)" json:"img_url"`
	ImgDesc   string `orm:"column(img_desc)" json:"img_desc"`
	SortOrder int    `orm:"column(sort_order)" json:"sort_order"`
}

func (*NideshopGoodsGallery) TableName() string {
	return "nideshop_goods_gallery"
}

func init() {
	orm.RegisterModel(new(NideshopGoodsGallery))
}
