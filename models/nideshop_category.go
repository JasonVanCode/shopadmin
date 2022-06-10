package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type NideshopCategory struct {
	Id           int                 `orm:"column(id)" json:"id"`
	Name         string              `orm:"column(name)" json:"name"`
	Keywords     string              `orm:"column(keywords)" json:"keywords"`
	FrontDesc    string              `orm:"column(front_desc)" json:"front_desc"`
	ParentId     int                 `orm:"column(parent_id)" json:"parent_id"`
	SortOrder    int                 `orm:"column(sort_order)" json:"sort_order"`
	ShowIndex    int                 `orm:"column(show_index)" json:"show_index"`
	IsShow       int                 `orm:"column(is_show)" json:"is_show"`
	BannerUrl    string              `orm:"column(banner_url)" json:"banner_url"`
	IconUrl      string              `orm:"column(icon_url)" json:"icon_url"`
	ImgUrl       string              `orm:"column(img_url)" json:"img_url"`
	WapBannerUrl string              `orm:"column(wap_banner_url)" json:"wap_banner_url"`
	Level        string              `orm:"column(level)" json:"level"`
	Type         int                 `orm:"column(type)" json:"type"`
	FrontName    string              `orm:"column(front_name)" json:"front_name"`
	Children     []*NideshopCategory `orm:"-" json:"children"`
}

func (*NideshopCategory) TableName() string {
	return "nideshop_category"
}

func (category *NideshopCategory) GetAttributeLists(name string) ([]NideshopAttribute, error) {
	var data []NideshopAttribute
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond = cond.And("attribute_category_id", category.Id)
	if name != "" {
		cond = cond.And("name__contains", name)
	}
	_, err := o.QueryTable(new(NideshopAttribute)).SetCond(cond).All(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func init() {
	orm.RegisterModel(new(NideshopCategory))
}
