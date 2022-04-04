package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

//是否收藏
type NideshopCollect struct {
	Id          int       `orm:"column(id)" json:"id"`
	UserId      int       `orm:"column(user_id)" json:"user_id"`
	ValueId     int       `orm:"column(value_id)" json:"value_id"` //goods_id
	AddTime     time.Time `orm:"column(add_time)" json:"add_time"`
	IsAttention int       `orm:"column(is_attention)" json:"is_attention"` //是否关注
	TypeId      int       `orm:"column(type_id)" json:"type_id"`
}

func (*NideshopCollect) TableName() string {
	return "nideshop_collect"
}

func init() {
	orm.RegisterModel(new(NideshopCollect))
}
