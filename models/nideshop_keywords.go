package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type NideshopKeywords struct {
	Id        int    `orm:"column(id)" json:"id"`
	Keyword   string `orm:"column(keyword)" json:"keyword"`
	IsHot     int    `orm:"column(is_hot)" json:"is_hot"`
	IsDefault int    `orm:"column(is_default)" json:"is_default"`
	IsShow    int    `orm:"column(is_show)" json:"is_show"`
	SortOrder int    `orm:"column(sort_order)" json:"sort_order"`
	Type      int    `orm:"column(type)" json:"type"`
	SchemeUrl string `orm:"column(scheme_url)" json:"scheme_url"`
}

func (*NideshopKeywords) TableName() string {
	return "nideshop_keywords"
}

func init() {
	orm.RegisterModel(new(NideshopKeywords))
}
