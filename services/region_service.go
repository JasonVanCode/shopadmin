package services

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
)

type RegionService struct {
	BaseService
}

//根据parent_id获取数据
func (*RegionService) GetRegionByParentId(pid int) []models.NideshopRegion {
	var regionLists []models.NideshopRegion
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopRegion)).Filter("parent_id", pid).All(&regionLists)
	if err != nil || num == 0 {
		return nil
	}
	return regionLists
}

//根据id获取数据
func (*RegionService) GetRegionById(id int) *models.NideshopRegion {
	var region models.NideshopRegion
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopRegion)).Filter("id", id).One(&region)
	if err != nil {
		return nil
	}
	return &region
}

//查询多条id数据
func (*RegionService) GetRegionByIds(ids []int) []*models.NideshopRegion {
	var regionLists []*models.NideshopRegion

	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopRegion)).Filter("id__in", ids).All(&regionLists)
	if err != nil || num == 0 {
		return nil
	}
	return regionLists
}

func (r *RegionService) SliceAddress(ids []int) string {
	var provice, city, district string
	fmt.Println(ids)
	regionLists := r.GetRegionByIds(ids)
	for _, v := range regionLists {
		switch v.Type {
		case 1:
			provice = v.Name
		case 2:
			city = v.Name
		case 3:
			district = v.Name
		}
	}
	return provice + city + district
}
