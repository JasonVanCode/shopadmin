package controllers

import (
	"fmt"
	"shopadmin/services"
	"shopadmin/utils"
)

type TopicController struct {
	BaseController
}

//获取专题数据
func (t *TopicController) GetTopicList() {
	var page, size = 1, 10
	if t.pageData != nil {
		page, size = t.pageData["page"], t.pageData["size"]
	}
	//处理数据
	data, err := new(services.TopicService).GetPageList(page, size)
	fmt.Println(*data, err)
	if err != nil {
		utils.HttpFail(500, err.Error(), t.Ctx)
	}
	utils.HttpSuccess(data, t.Ctx)
}
