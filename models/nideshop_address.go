package models

import (
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/utils"
)

//商品对应规格表值表
type NideshopAddress struct {
	Id         int    `orm:"column(id)" json:"id"`
	Name       string `orm:"column(name)" json:"name"`
	UserId     int    `orm:"column(user_id)" json:"user_id"`
	CountryId  int    `orm:"column(country_id)" json:"country_id"`
	ProvinceId int    `orm:"column(province_id)" json:"province_id"`
	CityId     int    `orm:"column(city_id)" json:"city_id"`
	DistrictId int    `orm:"column(district_id)" json:"district_id"`
	Address    string `orm:"column(address)" json:"address"`
	Mobile     string `orm:"column(mobile)" json:"mobile"`
	IsDefault  int    `orm:"column(is_default)" json:"is_default"`
}

func (*NideshopAddress) TableName() string {
	return "nideshop_address"
}

//拼接地址
func (a *NideshopAddress) SplicingAddress() []int {
	return []int{a.ProvinceId, a.CityId, a.DistrictId}
}

//获取对应省市区名称
func (a *NideshopAddress) GetPCDNames() map[int]string {
	ids := a.SplicingAddress()
	var res []orm.Params
	o := orm.NewOrm()
	_, err := o.QueryTable(new(NideshopRegion)).Filter("id__in", ids).Values(&res, "id", "name")
	if err != nil {
		return nil
	}
	return utils.TransSliceOrmParamsToMapIntString(res, "Id", "Name")
}

func init() {
	orm.RegisterModel(new(NideshopAddress))
}
