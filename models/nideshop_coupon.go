package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

//优惠券种类
type NideshopCoupon struct {
	Id             int       `orm:"column(id)" json:"id"`
	Name           string    `orm:"column(name)" json:"name"`
	TypeMoney      float64   `orm:"column(type_money);digits(10);decimals(2)" json:"type_money"`
	SendType       int       `orm:"column(send_type)" json:"send_type"`
	MinAmount      float64   `orm:"column(min_amount);digits(10);decimals(2)" json:"min_amount"` //最低使用金额
	MaxAmount      float64   `orm:"column(max_amount);digits(10);decimals(2)" json:"max_amount"`
	SendStartDate  time.Time `orm:"column(send_start_date)" json:"send_start_date"`
	SendEndDate    time.Time `orm:"column(send_end_date)" json:"send_end_date"`
	UseStartDate   time.Time `orm:"column(use_start_date)" json:"use_start_date"`
	UseEndDate     time.Time `orm:"column(use_end_date)" json:"use_end_date"`
	MinGoodsAmount float64   `orm:"column(min_goods_amount);digits(10);decimals(2)" json:"min_goods_amount"`
}

func (*NideshopCoupon) NideshopCoupon() string {
	return "nideshop_coupon"
}

func init() {
	orm.RegisterModel(new(NideshopCoupon))
}
