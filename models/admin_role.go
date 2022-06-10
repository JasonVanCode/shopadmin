package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type AdminRole struct {
	Id          int       `orm:"column(id);" json:"id"`
	Name        string    `orm:"column(name);" json:"name"`
	Description string    `orm:"column(description);" json:"description"`
	Url         string    `orm:"column(url);" json:"url"`
	NumOfUsers  int       `orm:"column(num_of_users) " json:"num_of_users"`
	Status      int       `orm:"column(status) " json:"status"`
	CreatedTime time.Time `orm:"column(created_time)" json:"created_time"`
	UpdateTime  time.Time `orm:"column(update_time);" json:"update_time"`
}

// NoDeletionId 禁止删除的数据id
func (*AdminRole) NoDeletionId() []int {
	return []int{1}
}

func (*AdminRole) TableName() string {
	return "admin_role"
}

func init() {
	orm.RegisterModel(new(AdminRole))
}
