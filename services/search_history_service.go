package services

import (
	"github.com/beego/beego/v2/client/orm"

	"shopadmin/models"
)

type SearchHistoryService struct {
	BaseService
}

//获取默认关键词
func (*SearchHistoryService) GetUserSearchHistory(userId int) []*models.NideshopSearchHistory {
	var data []*models.NideshopSearchHistory
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopSearchHistory)).Filter("user_id", userId).Distinct().Limit(10).All(&data)
	if err != nil || num == 0 {
		return nil
	}
	return data
}
