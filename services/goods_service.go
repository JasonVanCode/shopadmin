package services

import (
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
	"shopadmin/utils"
)

type GoodsService struct {
	BaseService
}

type GoodsListRtnJson struct {
	GoodsList []*models.NideshopGoods `json:"goodsList"`
	*PageData
	FilterCategories []FilterCategory `json:"filterCategory"`
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

type FilterCategory struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
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
func (g *GoodsService) GetGoodsLists(page, size int, otherParams map[string]string) (*PageData, []FilterCategory) {
	var categoryService CategoryService
	var goods []*models.NideshopGoods
	var qs orm.QuerySeter
	o := orm.NewOrm()
	qs = o.QueryTable(new(models.NideshopGoods))
	cond := orm.NewCondition()
	cond = cond.And("is_on_sale", 1)
	cond = cond.And("is_delete", 0)
	if categoryId, ok := otherParams["categoryId"]; ok {
		if categoryId != "0" {
			categoryIdInt := utils.TransStringToInt(categoryId)
			categoryIds := categoryService.GetCategoryIds(categoryService.GetCategoryByParentId(categoryIdInt))
			cond = cond.And("category_id__in", categoryIds)
		}
	}
	//关键词
	if keyworld, ok := otherParams["keyword"]; ok {
		cond = cond.And("name__icontains", keyworld)
	}
	//排序字段
	if sort, ok := otherParams["sort"]; ok {
		if sort == "price" {
			sort = "retail_price"
			if otherParams["order"] == "desc" {
				sort = "-" + sort
			}
		} else {
			//默认排序
			sort = "-id"
		}
		qs = qs.OrderBy(sort)
	}
	//获取所有分类的数据
	var categoryids []orm.Params
	qs.SetCond(cond).Distinct().Values(&categoryids, "category_id")
	categoryIntIds := utils.TransMapValueToSliceIntWithKey(categoryids, "CategoryId")

	//分类数据
	var filterCategories []FilterCategory
	if len(categoryIntIds) > 0 {
		parentIds := categoryService.GetParentCategoryIds(categoryIntIds)
		topParentLists := categoryService.GetCategoryByIds(parentIds)

		for _, value := range topParentLists {
			id := utils.ReturnIntTypeValue(value["Id"])
			filterCategories = append(filterCategories, FilterCategory{
				Id:      id,
				Name:    value["Name"].(string),
				Checked: false,
			})
		}
	}

	_, err := qs.SetCond(cond).Limit(size, (page-1)*size).All(&goods)
	goodsCount, _ := o.QueryTable(new(models.NideshopGoods)).SetCond(cond).Count()

	if err != nil {
		return nil, nil
	}
	return &PageData{
		size,
		page,
		int(goodsCount),
		g.GetTotalPage(int(goodsCount), size),
		goods,
	}, filterCategories
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

//---------------------后台商品管理列表--------------------------------

//商品状态请求数据
type GoodsStatusBody struct {
	Id       int `json:"id"`
	IsHot    int `json:"is_hot"`
	IsNew    int `json:"is_new"`
	IsOnSale int `json:"is_on_sale"`
	Type     int `json:"type"`
}

//商品列表
func (goods *GoodsService) GetAdminGoodsIndex(pageInfo map[string]int, otherParams map[string]string) (*PageData, error) {
	var page = pageInfo["page"]
	var size = pageInfo["size"]
	var goodsLists []models.NideshopGoods
	var qSeter orm.QuerySeter
	o := orm.NewOrm()
	cond := orm.NewCondition()
	//商品名
	if name, ok := otherParams["name"]; ok {
		cond = cond.And("name__icontains", name)
	}
	//商品编号
	if goods_no, ok := otherParams["goods_no"]; ok {
		cond = cond.And("goods_sn", goods_no)
	}
	//品牌
	if brand, ok := otherParams["brand_id"]; ok {
		cond = cond.And("brand_id", brand)
	}
	//分类
	if category_id, ok := otherParams["cat_id"]; ok {
		cond = cond.And("category_id", category_id)
	}
	//是否在售
	if is_on_sale, ok := otherParams["is_on_sale"]; ok {
		cond = cond.And("is_on_sale", is_on_sale)
	}
	qSeter = o.QueryTable(new(models.NideshopGoods))
	_, err := qSeter.SetCond(cond).OrderBy("id").Limit(size, (page-1)*size).All(&goodsLists)
	if err != nil {
		return nil, err
	}
	allCount, err := goods.GetCounts(qSeter, cond)
	if err != nil {
		return nil, err
	}
	return &PageData{
		NumsPerPage: size,
		CurrentPage: page,
		Count:       allCount,
		TotalPages:  goods.GetTotalPage(allCount, size),
		Data:        goodsLists,
	}, nil
}

//处理商品状态
func (goods *GoodsService) HandleGoodsStatus(body GoodsStatusBody) error {
	//需要修改的字段
	fileld, value := goods.GetChangeField(body)
	updateData := orm.Params{
		fileld: value,
	}
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopGoods)).Filter("id", body.Id).Update(updateData)
	if err != nil {
		return err
	}
	return nil
}

//1:修改是否上架 2:是否新品 3:是否热门
func (*GoodsService) GetChangeField(body GoodsStatusBody) (string, int) {
	switch body.Type {
	case 1:
		return "is_on_sale", body.IsOnSale
	case 2:
		return "is_new", body.IsNew
	case 3:
		return "is_hot", body.IsHot
	}
	return "", 0
}
