package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
	"shopadmin/utils"
)

type IndexController struct {
	BaseController
}

type newCategoryList struct {
	Id        int                    `json:"id"`
	Name      string                 `json:"name"`
	GoodsList []models.NideshopGoods `json:"goodsList"`
}

//首页返回数据
type resultdata struct {
	Newgoods     []models.NideshopGoods   `json:"newGoodsList"`
	Hotgoods     []models.NideshopGoods   `json:"hotGoodsList"`
	Topicgppds   []models.NideshopTopic   `json:"topicList"`
	Brands       []models.NideshopBrand   `json:"brandList"`
	ChannelLists []models.NideshopChannel `json:"channel"`
	CategoryList []newCategoryList        `json:"categoryList"`
	BannerLists  []models.NideshopAd      `json:"banner"`
}

func (i *IndexController) Index() {

	o := orm.NewOrm()
	//新产品
	var goods = new(models.NideshopGoods)
	var newgoods []models.NideshopGoods
	o.QueryTable(goods).Filter("is_new", 1).Limit(4).All(&newgoods)

	//热门产品
	var hotgoods []models.NideshopGoods
	o.QueryTable(goods).Filter("is_hot", 1).Limit(3).All(&hotgoods)
	//专题精选
	var topicgppds []models.NideshopTopic
	o.QueryTable(new(models.NideshopTopic)).Limit(3).All(&topicgppds)
	//品牌
	var brands []models.NideshopBrand
	_, err := o.QueryTable(new(models.NideshopBrand)).Filter("is_new", 1).Limit(4).All(&brands)
	fmt.Println(err)
	var channelLists []models.NideshopChannel
	o.QueryTable(new(models.NideshopChannel)).OrderBy("sort_order").All(&channelLists)
	//广告栏
	var bannerLists []models.NideshopAd
	o.QueryTable(new(models.NideshopAd)).Filter("ad_position_id", 1).All(&bannerLists)

	//产品分类
	var category = new(models.NideshopCategory)
	var categoryList []models.NideshopCategory
	o.QueryTable(category).Exclude("name", "推荐").Filter("parent_id", 0).All(&categoryList)

	var newGoodsList []newCategoryList

	for _, categoryItem := range categoryList {
		var ids []orm.Params
		o.QueryTable(category).Filter("parent_id", categoryItem.Id).Values(&ids, "id")

		idsInt := utils.TransMapValueToSliceInt(ids)
		var cgoods []models.NideshopGoods
		o.QueryTable(goods).Filter("category_id__in", idsInt).Limit(7).All(&cgoods)

		newGoodsList = append(newGoodsList, newCategoryList{
			categoryItem.Id,
			categoryItem.Name,
			cgoods,
		})
	}
	utils.HttpSuccess(&resultdata{
		Newgoods:     newgoods,
		Hotgoods:     hotgoods,
		Topicgppds:   topicgppds,
		Brands:       brands,
		ChannelLists: channelLists,
		CategoryList: newGoodsList,
		BannerLists:  bannerLists,
	}, i.Ctx)
}
