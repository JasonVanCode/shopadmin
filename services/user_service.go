package services

import (
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
)

//根据id获取用户的信息
func GetNideshop_userById(id int) (v *models.Nideshop_user, err error) {
	o := orm.NewOrm()
	v = &models.Nideshop_user{Id: id}
	if err = o.QueryTable(new(models.Nideshop_user)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
