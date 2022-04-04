package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

//用户优惠券
type NideshopUserCoupon struct {
	Id           int       `orm:"column(id)" json:"id"`
	CouponId     int       `orm:"column(coupon_id)" json:"coupon_id"`
	CouponNumber string    `orm:"column(coupon_number)" json:"coupon_number"`
	UserId       int       `orm:"column(user_id)" json:"user_id"`
	UsedTime     time.Time `orm:"column(used_time)" json:"used_time"`
	OrderId      int       `orm:"column(order_id)" json:"order_id"`
}

func (*NideshopUserCoupon) TableName() string {
	return "nideshop_user_coupon"
}

func init() {
	orm.RegisterModel(new(NideshopUserCoupon))
}
