package services

import (
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
)

type UserCouponService struct {
	BaseService
}

//获取当前用户未使用的优惠券
func (*UserCouponService) GetUnusedCoupon(userId int) []*models.NideshopUserCoupon {
	var couponList []*models.NideshopUserCoupon
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopUserCoupon)).Filter("userId", userId).Filter("order_id", 0).All(&couponList)
	if err != nil || num == 0 {
		return nil
	}
	return couponList
}
