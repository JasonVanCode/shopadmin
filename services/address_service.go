package services

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/models"
)

type AddressService struct {
	BaseService
	DeleteBody
}

//api 请求体
type AddressSaveBody struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Mobile     string `json:"mobile"`
	ProvinceId int    `json:"province_id"`
	CityId     int    `json:"city_id"`
	DistrictId int    `json:"district_id"`
	Address    string `json:"address"`
	IsDefault  bool   `json:"is_default"`
}

func (body *AddressSaveBody) Valid(v *validation.Validation) {
	if body.Name == "" {
		v.SetError("姓名", "不能为空")
	}
	if body.Address == "" {
		v.SetError("地址", "不能为空")
	}
	if body.Mobile == "" {
		v.SetError("手机号", "不能为空")
	}
	if body.ProvinceId == 0 {
		v.SetError("省份", "不能为空")
	}
	if body.CityId == 0 {
		v.SetError("城市", "不能为空")
	}
	if body.DistrictId == 0 {
		v.SetError("区县", "不能为空")
	}
}

//获取地址信息
func (*AddressService) GetAddressLIst(userId int) []*models.NideshopAddress {
	var data []*models.NideshopAddress
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopAddress)).Filter("user_id", userId).All(&data)
	if err != nil {
		return nil
	}
	return data
}

//插入数据
func (*AddressService) AddAddress(data *AddressSaveBody, userId int) int {
	o := orm.NewOrm()
	var isDefault int
	if data.IsDefault {
		isDefault = 1
	} else {
		isDefault = 0
	}
	saveData := models.NideshopAddress{
		Name:       data.Name,
		Mobile:     data.Mobile,
		Address:    data.Address,
		ProvinceId: data.ProvinceId,
		DistrictId: data.DistrictId,
		CityId:     data.CityId,
		IsDefault:  isDefault,
		UserId:     userId,
	}
	num, err := o.Insert(&saveData)
	if err != nil || num == 0 {
		return 0
	}
	return int(num)
}

//编辑数据
func (*AddressService) UpdateAddress(data *AddressSaveBody) error {
	var isDefault int
	if data.IsDefault {
		isDefault = 1
	} else {
		isDefault = 0
	}
	o := orm.NewOrm()
	updateData := models.NideshopAddress{
		Id:         data.Id,
		Name:       data.Name,
		Mobile:     data.Mobile,
		Address:    data.Address,
		ProvinceId: data.ProvinceId,
		DistrictId: data.DistrictId,
		CityId:     data.CityId,
		IsDefault:  isDefault,
	}
	_, err := o.Update(&updateData)
	if err != nil {
		return err
	}
	return nil
}

//更新默认地址
func (*AddressService) ChangeDefaultAddress(userId, nowDefaultAddId int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopAddress)).Filter("user_id", userId).Exclude("id", nowDefaultAddId).Update(orm.Params{
		"is_default": 0,
	})
	if err != nil {
		return err
	}
	return nil
}

type DeleteBody struct {
	Id int
}

//删除数据
func (*AddressService) DeleteAddress(id int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopAddress)).Filter("id", id).Delete()
	if err != nil {
		return err
	}
	return nil
}

//获取默认地址
func (*AddressService) GetDefaultAddress(userId int) *models.NideshopAddress {
	var add models.NideshopAddress
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopAddress)).Filter("user_id", userId).Filter("is_default", 1).One(&add)
	if err != nil {
		return nil
	}
	return &add
}

func (*AddressService) GetAddressById(id int) *models.NideshopAddress {
	var add models.NideshopAddress
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopAddress)).Filter("id", id).One(&add)
	if err != nil {
		return nil
	}
	return &add
}
