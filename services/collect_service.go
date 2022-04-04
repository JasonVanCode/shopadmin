package services

import (
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
)

type CollectService struct {
	BaseService
}

//判断用户是否收藏该商品
func (*CollectService) IsUserCollect(userId, typeId, goodsId int) int {
	o := orm.NewOrm()
	var collect models.NideshopCollect
	err := o.QueryTable(new(models.NideshopCollect)).Filter("user_id", userId).Filter("type_id", typeId).Filter("value_id", goodsId).One(&collect)
	if err != nil || collect.Id == 0 {
		return 0
	}
	return 1
}
