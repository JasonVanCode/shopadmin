package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type NideshopChannel struct {
	Id        int    `orm:"column(id)" json:"id"`
	Name      string `orm:"column(name)" json:"name"`
	Url       string `orm:"column(url)" json:"url"`
	IconUrl   string `orm:"column(icon_url)" json:"icon_url"`
	SortOrder int    `orm:"column(sort_order)" json:"sort_order"`
}

func (*NideshopChannel) TableName() string {
	return "nideshop_channel"
}

func init() {
	orm.RegisterModel(new(NideshopChannel))
}
