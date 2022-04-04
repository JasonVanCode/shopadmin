package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/services"
	"shopadmin/utils"
)

type AddressController struct {
	BaseController
}

//获取地址信息
func (a *AddressController) GetAddressList() {
	var addressService = new(services.AddressService)
	var regionService = new(services.RegionService)
	data := addressService.GetAddressLIst(a.user.Id)
	var resultData []ListData
	for _, v := range data {
		fmt.Println(*v)
		resultData = append(resultData, ListData{
			"id":          v.Id,
			"name":        v.Name,
			"is_default":  v.IsDefault,
			"full_region": regionService.SliceAddress(v.SplicingAddress()),
			"mobile":      v.Mobile,
			"address":     v.Address,
		})
	}
	utils.HttpSuccess(resultData, a.Ctx)
}

//地址保存和修改
func (a *AddressController) SaveAddress() {
	var addressService = new(services.AddressService)
	var body services.AddressSaveBody
	if err := json.Unmarshal(a.Ctx.Input.RequestBody, &body); err != nil {
		utils.HttpFail(400, err.Error(), a.Ctx)
	}
	//验证必填数据是否填写
	valid := validation.Validation{}
	if ok, err := valid.Valid(body); err != nil || !ok {
		utils.HttpFail(400, "请求数据有误", a.Ctx)
	}
	userId := a.user.Id
	//判断是修改数据还是新添加数据
	if body.Id == 0 {
		body.Id = addressService.AddAddress(&body, userId)
		if body.Id == 0 {
			utils.HttpFail(400, "数据插入失败", a.Ctx)
		}
	} else {
		if err := addressService.UpdateAddress(&body); err != nil {
			utils.HttpFail(400, err.Error(), a.Ctx)
		}
	}

	//默认地址判断
	if body.IsDefault {
		if err := addressService.ChangeDefaultAddress(userId, body.Id); err != nil {
			utils.HttpFail(400, err.Error(), a.Ctx)
		}
	}
	utils.HttpSuccess(nil, a.Ctx)
}

//删除地址
func (a *AddressController) DeleteAddress() {
	addressService := new(services.AddressService)
	deleteBody := addressService.DeleteBody
	if err := json.Unmarshal(a.Ctx.Input.RequestBody, &deleteBody); err != nil || deleteBody.Id == 0 {
		utils.HttpFail(400, "请求数据有误", a.Ctx)
	}
	if err := addressService.DeleteAddress(deleteBody.Id); err != nil {
		utils.HttpFail(400, err.Error(), a.Ctx)
	}
	utils.HttpSuccess(nil, a.Ctx)
}
