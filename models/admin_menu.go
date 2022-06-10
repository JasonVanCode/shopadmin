package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type AdminMenu struct {
	Id          int         `orm:"column(id);" json:"id"`
	ParentId    int         `orm:"column(parent_id);" json:"parent_id"`
	Name        string      `orm:"column(name);" json:"name"`
	Url         string      `orm:"column(url);" json:"url"`
	Icon        string      `orm:"column(icon);" json:"icon"`
	Status      int         `orm:"column(status) " json:"status"`
	SortId      int         `orm:"column(sort_id) " json:"sort_id"`
	Level       int         `orm:"column(level) " json:"level"`
	LogMethod   string      `orm:"column(log_method);" json:"log_method"`
	CreatedTime time.Time   `orm:"column(created_time)" json:"created_time"`
	UpdateTime  time.Time   `orm:"column(update_time);" json:"update_time"`
	Children    []AdminMenu `orm:"-" json:"children"`
}

func (*AdminMenu) TableName() string {
	return "admin_menu"
}

func init() {
	orm.RegisterModel(new(AdminMenu))
}
