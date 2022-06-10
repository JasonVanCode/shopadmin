package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type NideshopSearchHistory struct {
	Id      int       `orm:"column(id)" json:"id"`
	Keyword string    `orm:"column(keyword)" json:"keyword"`
	From    string    `orm:"column(from)" json:"from"`
	AddTime time.Time `orm:"column(add_time)" json:"add_time"`
	UserId  int       `orm:"column(user_id)" json:"user_id"`
}

func (*NideshopSearchHistory) TableName() string {
	return "nideshop_search_history"
}

func init() {
	orm.RegisterModel(new(NideshopSearchHistory))
}
