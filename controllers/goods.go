package controllers

import (
	"encoding/base64"
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
)

type GoodsController struct {
	BaseController
}

type CountJson struct {
	GoodsCount int `json:"goodsCount"`
}

//获取再售商品的数量
func (g *GoodsController) GetCount() {
	var goodsService = new(services.GoodsService)
	num := goodsService.GetCount()
	utils.HttpSuccess(&CountJson{num}, g.Ctx)
}

//商品分类数据
func (g *GoodsController) GetCategory() {
	goodsId := utils.TransStringToInt(g.GetString("id"))
	if goodsId == 0 {
		utils.HttpSuccess(nil, g.Ctx)
	}
	var categoryService = new(services.CategoryService)
	var currentCtegory = categoryService.GetCategoryById(goodsId)
	var parentCtegory = categoryService.GetCategoryById(currentCtegory.ParentId)
	var brotherCtegorys = categoryService.GetCategoryByParentId(currentCtegory.ParentId)
	rtnJson := categoryService.CategoryRtnJson
	rtnJson.CurCategory = currentCtegory
	rtnJson.ParentCategory = parentCtegory
	rtnJson.BrotherCategory = brotherCtegorys
	utils.HttpSuccess(rtnJson, g.Ctx)
}

//商品列表数据
func (g *GoodsController) GetGoodsList() {
	page, size := g.pageData["page"], g.pageData["size"]
	var goodsService services.GoodsService
	goods, filterCategories := goodsService.GetGoodsLists(page, size, g.getData)
	resultData := services.GoodsListRtnJson{
		goods.Data.([]*models.NideshopGoods),
		goods,
		filterCategories,
	}
	utils.HttpSuccess(resultData, g.Ctx)
}

//商品详情
func (g *GoodsController) GetGoodsDetail() {
	//当前登录用户
	userId := g.user.Id
	id := utils.TransStringToInt(g.GetString("id"))
	if id == 0 {
		utils.HttpFail(400, "该商品不存在", g.Ctx)
	}
	var goodsService = new(services.GoodsService)
	//当前商品
	nowGood := goodsService.GetGoodsById(id)
	if nowGood == nil {
		utils.HttpFail(400, "该商品不存在", g.Ctx)
	}
	//商品问题咨询
	var issueService = new(services.GoodsIssueService)
	issues := issueService.GetIssueList(0)
	//商品展示图片
	galleryList := nowGood.GetGoodsGallery(4)
	//商品属性
	attribute := nowGood.GetGoodsAttribute()
	//商品品牌
	goodBrand := nowGood.GetGoodsBrand()
	//该商品评论数量
	var commentInfo services.CommentInfo
	commentCount := nowGood.GetCommentCount()
	if commentCount > 0 {
		hotComment := nowGood.GetOneHotComment()
		decodeContent, _ := base64.StdEncoding.DecodeString(hotComment.Content)
		commentInfo = services.CommentInfo{
			Content:  string(decodeContent),
			AddTime:  hotComment.AddTime.Format(TimeFormat),
			NickName: hotComment.User.Nickname,
			Avatar:   hotComment.User.Avatar,
			PicList:  hotComment.CommentPictures,
		}
	}
	comment := services.Comment{
		Count: commentCount,
		Data:  commentInfo,
	}
	//查看当前用户是否收藏该商品
	var colletService = new(services.CollectService)
	hasCollectd := colletService.IsUserCollect(userId, 0, id)

	//产品信息
	productList := nowGood.GetProductList()
	//商品规格信息
	slist := nowGood.GetGoodsSpecificationList()

	//添加足迹信息
	var footPrint = new(services.FootPrintService)
	footPrint.AddFootprint(userId, id)

	utils.HttpSuccess(services.DetailRtnJson{
		services.SkuRtnJson{
			productList,
			slist,
		},
		nowGood,
		galleryList,
		attribute,
		issues,
		hasCollectd,
		comment,
		goodBrand,
	}, g.Ctx)
}

//获取相关商品
func (g *GoodsController) GetRelatedGoods() {
	id := utils.TransStringToInt(g.GetString("id"))
	if id == 0 {
		utils.HttpFail(400, "该商品不存在", g.Ctx)
	}
	var goodsService = new(services.GoodsService)

	nowGoods := goodsService.GetGoodsById(id)
	if nowGoods == nil {
		utils.HttpFail(400, "该商品不存在", g.Ctx)
	}
	var relatedList []*models.NideshopGoods
	relatedList = goodsService.GetRelatedGoods(id)
	//找同分类的产品
	if relatedList == nil {
		relatedList = goodsService.GetGoodsByCategoryId(nowGoods.CategoryId)
	}
	utils.HttpSuccess(services.RelatedRtnJson{GoodsList: relatedList}, g.Ctx)
}
