package services

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
)

type GoodsService struct {
	BaseService
}

type GoodsListRtnJson struct {
	GoodsList []*models.NideshopGoods `json:"goodsList"`
}

type Comment struct {
	Count int         `json:"count"`
	Data  CommentInfo `json:"data"`
}
type CommentInfo struct {
	Content  string                           `json:"content"`
	AddTime  string                           `json:"add_time"`
	NickName string                           `json:"nick_name"`
	Avatar   string                           `json:"avatar"`
	PicList  []*models.NideshopCommentPicture `json:"pic_list"`
}

type SkuRtnJson struct {
	ProductList       []*models.NideshopProduct   `json:"productList"`
	SpecificationList []*models.SpecificationItem `json:"specificationList"`
}

type DetailRtnJson struct {
	SkuRtnJson
	Goods          *models.NideshopGoods          `json:"info"`
	Galleries      []*models.NideshopGoodsGallery `json:"gallery"`
	Attribute      []map[string]string            `json:"attribute"`
	Issues         []*models.NideshopGoodsIssue   `json:"issue"`
	UserHasCollect int                            `json:"userHasCollect"`
	Comment        Comment                        `json:"comment"`
	Brand          *models.NideshopBrand          `json:"brand"`
}

//获取再售商品的总数
func (g *GoodsService) GetCount() int {
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopGoods)).Filter("is_on_sale", 1).Filter("is_delete", 0).Count()
	if err != nil {
		return 0
	}
	return int(num)
}

//根据商品id获取信息
func (g *GoodsService) GetGoodsById(id int) *models.NideshopGoods {
	var goods models.NideshopGoods
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopGoods)).Filter("id", id).One(&goods)
	if err != nil || goods.Id == 0 {
		return nil
	}
	return &goods
}

//获取商品列表数据
func (g *GoodsService) GetGoodsLists(page, size int, otherParams map[string]string) []*models.NideshopGoods {
	var goods []*models.NideshopGoods
	o := orm.NewOrm()
	cond := orm.NewCondition()
	if categoryId, ok := otherParams["categoryId"]; ok {
		if categoryId != "0" {
			cond = cond.And("category_id", categoryId)
		}

	}
	_, err := o.QueryTable(new(models.NideshopGoods)).SetCond(cond).Limit(size, page).All(&goods)
	fmt.Println(err)
	if err != nil {
		return nil
	}
	return goods
}

type RelatedRtnJson struct {
	GoodsList []*models.NideshopGoods `json:"goodsList"`
}

//获取相关商品的数据
func (g *GoodsService) GetRelatedGoods(goodsId int) []*models.NideshopGoods {
	var relatedLIst []*models.NideshopRelatedGoods
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopRelatedGoods)).RelatedSel().Filter("goods_id", goodsId).All(&relatedLIst)
	if err != nil || num == 0 {
		return nil
	}
	var goods []*models.NideshopGoods
	for _, v := range relatedLIst {
		goods = append(goods, v.RelatedGoods)
	}
	return goods
}

//根据分类获取商品
func (g *GoodsService) GetGoodsByCategoryId(categoryId int) []*models.NideshopGoods {
	var goods []*models.NideshopGoods
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopGoods)).Filter("category_id", categoryId).Limit(8).All(&goods)
	if err != nil || num == 0 {
		return nil
	}
	return goods
}
