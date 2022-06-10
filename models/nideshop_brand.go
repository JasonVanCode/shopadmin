package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"shopadmin/utils"
)

type NideshopBrand struct {
	Id             int     ` orm:"column(id)" json:"id"`
	Name           string  ` orm:"column(name)" json:"name"`
	ListPicUrl     string  ` orm:"column(list_pic_url)" json:"list_pic_url"`
	SimpleDesc     string  ` orm:"column(simple_desc)" json:"simple_desc"`
	PicUrl         string  ` orm:"column(pic_url)" json:"pic_url"`
	SortOrder      int     ` orm:"column(sort_order)" json:"sort_order"`
	IsShow         int     ` orm:"column(is_show)" json:"is_show"`
	FloorPrice     float64 ` orm:"column(floor_price);digits(10);decimals(2)" json:"floor_price"`
	AppListPicurl  string  ` orm:"column(app_list_pic_url)" json:"app_list_pic_url"`
	IsNew          int     ` orm:"column(is_new)" json:"is_new"`
	NewPicUrl      string  ` orm:"column(new_pic_url)" json:"new_pic_url"`
	NewSortOrder   int     ` orm:"column(new_sort_order)" json:"new_sort_order"`
	GoodsCounts    int     `orm:"-" json:"goods_counts"`
	CommentsCounts int     `orm:"-" json:"comments_counts"`
}

func (*NideshopBrand) TableName() string {
	return "nideshop_brand"
}

//获取该品牌商品数量
func (brand *NideshopBrand) GetGoodsCounts() int {
	o := orm.NewOrm()
	num, err := o.QueryTable(new(NideshopGoods)).Filter("brand_id", brand.Id).Count()
	if err != nil {
		logs.Error(err.Error() + "\n")
		return 0
	}
	return int(num)
}

//获取该品牌评价数量
func (brand *NideshopBrand) GetCommentsCount() int {
	var goodsIds []orm.Params
	o := orm.NewOrm()
	num, err := o.QueryTable(new(NideshopGoods)).Filter("brand_id", brand.Id).Values(&goodsIds, "id")
	if err != nil {
		logs.Error(err.Error() + "\n")
		return 0
	}
	if num == 0 {
		return 0
	}
	ids := utils.TransMapValueToSliceIntWithKey(goodsIds, "Id")
	num, err = o.QueryTable(new(NideshopComment)).Filter("value_id__in", ids).Count()
	if err != nil {
		logs.Error(err.Error() + "\n")
		return 0
	}
	return int(num)
}

func init() {
	orm.RegisterModel(new(NideshopBrand))
}
