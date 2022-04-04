package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type NideshopGoodsIssue struct {
	Id       int    `orm:"column(id)" json:"id"`
	GoodsId  string `orm:"column(goods_id)" json:"goods_id"`
	Question string `orm:"column(question)" json:"question"`
	Answer   string `orm:"column(answer)" json:"answer"`
}

func (*NideshopGoodsIssue) TableName() string {
	return "nideshop_goods_issue"
}

func init() {
	orm.RegisterModel(new(NideshopGoodsIssue))
}
