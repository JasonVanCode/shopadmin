package controllers

import (
	"shopadmin/services"
	"shopadmin/utils"
)

type CommentController struct {
	BaseController
}

//评价列表信息
func (c *CommentController) GetCommentList() {
	var commentService services.CommentService
	pageData := commentService.GetCommentList(c.pageData, c.getData)
	utils.HttpSuccess(pageData, c.Ctx)
}

type CountRtnJson struct {
	AllCount    int `json:"allCount"`
	HasPicCount int `json:"hasPicCount"`
}

//获取评价数量
func (c *CommentController) GetCommentCount() {
	var commentService services.CommentService
	typeId := c.getData["typeId"]
	valueId := c.getData["valueId"]
	allCount := commentService.GetCommentCount("0", valueId, typeId)
	hasPicCount := commentService.GetCommentCount("1", valueId, typeId)
	utils.HttpSuccess(CountRtnJson{
		allCount,
		hasPicCount,
	}, c.Ctx)
}
