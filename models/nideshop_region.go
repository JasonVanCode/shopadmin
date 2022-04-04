package models

import (
	"github.com/beego/beego/v2/client/orm"
)

//商品对应规格表值表
type NideshopRegion struct {
	Id       int    `orm:"column(id)" json:"id"`
	ParentId int    `orm:"column(parent_id)" json:"parent_id"`
	Name     string `orm:"column(name)" json:"name"`
	AgencyId int    `orm:"column(agency_id)" json:"agency_id"`
	Type     int    `orm:"column(type)" json:"type"`
}

func (*NideshopRegion) TableName() string {
	return "nideshop_region"
}

func init() {
	orm.RegisterModel(new(NideshopRegion))
}
