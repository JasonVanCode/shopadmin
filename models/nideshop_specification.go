package models

import (
	"github.com/beego/beego/v2/client/orm"
)

//规格表
type NideshopSpecification struct {
	Id        int    `orm:"column(id)" json:"id"`
	Name      string `orm:"column(name)" json:"name"`
	SortOrder int    `orm:"column(sort_order)" json:"sort_order"`
}

func (*NideshopSpecification) TableName() string {
	return "nideshop_specification"
}

func init() {
	orm.RegisterModel(new(NideshopSpecification))
}
