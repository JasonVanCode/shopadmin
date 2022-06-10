package services

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
	"strings"
)

type AttributeService struct {
	BaseService
}

type RequestAttributeBody struct {
	Id                  int    `json:"id"`
	AttributeCategoryId int    `json:"attribute_category_id"`
	Name                string `json:"name"`
	SortOrder           int    `json:"sort_order"`
}

type AttributeDelBody struct {
	Id string `json:"id"`
}

//判断属性名是否重复
func (*AttributeService) IsAttributeNameExists(id int, name string) bool {
	var attrInfo models.NideshopAttribute
	var qSeter orm.QuerySeter
	o := orm.NewOrm()
	qSeter = o.QueryTable(new(models.NideshopAttribute)).Filter("name", name)
	if id > 0 {
		qSeter = qSeter.Exclude("id", id)
	}
	err := qSeter.One(&attrInfo, "id")
	if err != nil || attrInfo.Id == 0 {
		return false
	}
	return true
}

//新增属性
func (attr *AttributeService) CreateAttribute(body RequestAttributeBody) (*models.NideshopAttribute, error) {
	if attr.IsAttributeNameExists(0, body.Name) {
		return nil, errors.New("该属性名已经存在")
	}
	var attInfo = models.NideshopAttribute{
		Name:                body.Name,
		AttributeCategoryId: body.AttributeCategoryId,
		SortOrder:           body.SortOrder,
	}
	o := orm.NewOrm()
	_, err := o.Insert(&attInfo)
	if err != nil {
		return nil, err
	}
	return &attInfo, nil
}

//编辑属性
func (attr *AttributeService) UpdateAttribute(body RequestAttributeBody) (*models.NideshopAttribute, error) {
	if attr.IsAttributeNameExists(body.Id, body.Name) {
		return nil, errors.New("该属性名已经存在")
	}
	var attInfo = models.NideshopAttribute{
		Id:                  body.Id,
		Name:                body.Name,
		AttributeCategoryId: body.AttributeCategoryId,
		SortOrder:           body.SortOrder,
	}
	o := orm.NewOrm()
	_, err := o.Update(&attInfo, "id", "name", "attribute_category_id", "sort_order")
	if err != nil {
		return nil, err
	}
	return &attInfo, nil
}

//删除属性
func (attr *AttributeService) DelAttribute(id string) error {
	var ids_arr = strings.Split(id, ",")
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopAttribute)).Filter("id__in", ids_arr).Delete()
	if err != nil {
		return err
	}
	return nil
}
