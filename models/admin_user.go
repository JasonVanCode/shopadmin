package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type AdminUser struct {
	Id          int       `orm:"column(id);" json:"id"`
	Username    string    `orm:"column(username);" json:"username"`
	Password    string    `orm:"column(password);" json:"password"`
	Nickname    string    `orm:"column(nickname);" json:"nickname"`
	Avatar      string    `orm:"column(avatar);" json:"avatar"`
	Role        string    `orm:"column(role) " json:"role"`
	DeleteTime  time.Time `orm:"column(delete_time)" json:"delete_time"`
	Status      int       `orm:"column(status) " json:"status"`
	CreatedTime time.Time `orm:"column(created_time)" json:"created_time"`
	UpdateTime  time.Time `orm:"column(update_time);" json:"update_time"`
}

// NoDeletionId 禁止删除的数据id
func (*AdminUser) NoDeletionId() []int {
	return []int{1}
}

func (*AdminUser) TableName() string {
	return "admin_user"
}

func init() {
	orm.RegisterModel(new(AdminUser))
}
