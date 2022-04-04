package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type NideshopAttribute struct {
	Id                  int    `orm:"column(id)" json:"id"`
	AttributeCategoryId int    `orm:"column(attribute_category_id)" json:"attribute_category_id"`
	Name                string `orm:"column(name)" json:"name"`
	InputType           int    `orm:"column(input_type)" json:"input_type"`
	Values              string `orm:"column(values)" json:"values"`
	SortOrder           int    `orm:"column(sort_order)" json:"sort_order"`
}

func (*NideshopAttribute) TableName() string {
	return "nideshop_attribute"
}

func init() {
	orm.RegisterModel(new(NideshopAttribute))
}
