package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type NideshopGoods struct {
	Id                int       `orm:"column(id)" json:"id"`
	Name              string    `orm:"column(name)" json:"name"`
	GoodsSn           string    `orm:"column(goods_sn)" json:"goods_sn"`
	CategoryId        int       `orm:"column(category_id)" json:"category_id"`
	BrandId           int       `orm:"column(brand_id)" json:"brand_id"`
	GoodsNumber       int       `orm:"column(goods_number)" json:"goods_number"`
	Keywords          string    `orm:"column(keywords)" json:"keywords"`
	GoodsBrief        string    `orm:"column(goods_brief)" json:"goods_brief"`
	GoodsDesc         string    `orm:"column(goods_desc)" json:"goods_desc"`
	IsOnSale          int       `orm:"column(is_on_sale);description(1 在售 0 下架)" json:"is_on_sale"`
	AddTime           time.Time `orm:"column(add_time)" json:"add_time"`
	SortOrder         int       `orm:"column(sort_order)" json:"sort_order"`
	IsDelete          int       `orm:"column(is_delete)" json:"is_delete"`
	AttributeCategory int       `orm:"column(attribute_category)" json:"attribute_category"`
	CounterPrice      float64   `orm:"column(counter_price);digits(10);decimals(2);description(专柜价格)" json:"counter_price"`
	ExtraPrice        float64   `orm:"column(extra_price);digits(10);decimals(2);description(附加价格)" json:"extra_price"`
	IsNew             int       `orm:"column(is_new)" json:"is_new"`
	GoodsUnit         string    `orm:"column(goods_unit);description(商品单位)" json:"goods_unit"`
	PrimaryPicUrl     string    `orm:"column(primary_pic_url);description(商品主图)" json:"primary_pic_url"`
	ListPicUrl        string    `orm:"column(list_pic_url);description(商品列表图)" json:"list_pic_url"`
	RetailPrice       float64   `orm:"column(retail_price);digits(10);decimals(2);description(零售价格)" json:"retail_price"`
	SellVolume        int       `orm:"column(sell_volume);description(销售量)" json:"sell_volume"`
	PrimaryProductId  int       `orm:"column(primary_product_id);description(主sku　product_id)" json:"primary_product_id"`
	UnitPrice         float64   `orm:"column(unit_price);digits(10);decimals(2);description(单位价格，单价)" json:"unit_price"`
	PromotionDesc     string    `orm:"column(promotion_desc)" json:"promotion_desc"`
	PromotionTag      string    `orm:"column(promotion_tag)" json:"promotion_tag"`
	AppExclusivePrice float64   `orm:"column(app_exclusive_price);digits(10);decimals(2);description(APP专享价)" json:"app_exclusive_price"`
	IsAppExclusive    int       `orm:"column(is_app_exclusive);description(是否是APP专属)" json:"is_app_exclusive"`
	IsLimited         int       `orm:"column(is_limited)" json:"is_limited"`
	IsHot             int       `orm:"column(is_hot)" json:"is_hot"`
}

func (*NideshopGoods) TableName() string {
	return "nideshop_goods"
}

//获取商品详情展示图片
func (n *NideshopGoods) GetGoodsGallery(limit int) []*NideshopGoodsGallery {
	var gallery []*NideshopGoodsGallery
	o := orm.NewOrm()
	_, err := o.QueryTable(new(NideshopGoodsGallery)).Filter("goods_id", n.Id).Limit(limit).All(&gallery)
	if err != nil {
		return nil
	}
	return gallery
}

//获取商品属性
func (n *NideshopGoods) GetGoodsAttribute() []map[string]string {
	var attribute []*NideshopGoodsAttribute
	o := orm.NewOrm()
	_, err := o.QueryTable(new(NideshopGoodsAttribute)).Filter("goods_id", n.Id).RelatedSel().All(&attribute)
	if err != nil {
		return nil
	}

	var resultMap []map[string]string
	for _, v := range attribute {
		resultMap = append(resultMap, map[string]string{
			"name":  v.Attribute.Name,
			"value": v.Value,
		})
	}
	return resultMap
}

//获取商品品牌
func (n *NideshopGoods) GetGoodsBrand() *NideshopBrand {
	var brand NideshopBrand
	o := orm.NewOrm()
	err := o.QueryTable(new(NideshopBrand)).Filter("id", n.BrandId).One(&brand)
	if err != nil {
		return nil
	}
	return &brand
}

//获取当前商品总评价数
func (n *NideshopGoods) GetCommentCount() int {
	o := orm.NewOrm()
	num, err := o.QueryTable(new(NideshopComment)).Filter("value_id", n.Id).Filter("type_id", 0).Count()
	if err != nil {
		return 0
	}
	return int(num)
}

//获取随机一个热门评价
func (n *NideshopGoods) GetOneHotComment() *NideshopComment {
	var hotComment NideshopComment
	o := orm.NewOrm()
	err := o.QueryTable(new(NideshopComment)).Filter("id", 1).RelatedSel().Filter("type_id", 0).One(&hotComment)
	if err != nil {
		return nil
	}
	//载入对应的图片数据
	o.LoadRelated(&hotComment, "CommentPictures")
	return &hotComment
}

//获取该商品的产品信息
func (n *NideshopGoods) GetProductList() []*NideshopProduct {
	var productList []*NideshopProduct
	o := orm.NewOrm()
	num, err := o.QueryTable(new(NideshopProduct)).Filter("goods_id", n.Id).All(&productList)
	if err != nil || num == 0 {
		return nil
	}
	return productList
}

//返回的数据
type SpecificationItem struct {
	Specification_id int                          `json:"specification_id"`
	Name             string                       `json:"name"`
	List             []NideshopGoodsSpecification `json:"list"`
}

//获取商品的规格信息
func (n *NideshopGoods) GetGoodsSpecificationList() []*SpecificationItem {
	var list []NideshopGoodsSpecification
	o := orm.NewOrm()
	num, err := o.QueryTable(new(NideshopGoodsSpecification)).RelatedSel().Filter("goods_id", n.Id).All(&list)
	var data []*SpecificationItem
	//确保有数据
	resData := make(map[int]*SpecificationItem)
	if err == nil && num > 0 {
		for _, v := range list {
			specificationId := v.Specification.Id
			if item, ok := resData[specificationId]; ok {
				item.List = append(item.List, v)
			} else {
				fmt.Println("b", v.Id)
				resData[specificationId] = &SpecificationItem{
					specificationId,
					v.Specification.Name,
					[]NideshopGoodsSpecification{
						v,
					},
				}
			}
		}
		//存切片中
		for _, v := range resData {
			data = append(data, v)
		}
	}
	return data
}

func init() {
	orm.RegisterModel(new(NideshopGoods))
}
