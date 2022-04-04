package controllers

import (
	"shopadmin/services"
	"shopadmin/utils"
	"sort"
	"time"
)

type FootprintController struct {
	BaseController
}

//返回前端的数据
type IndexFootprintData struct {
	Name        string  `json:"name"`
	ListPicUrl  string  `json:"list_pic_url"`
	GoodsBrief  string  `json:"goods_brief"`
	RetailPrice float64 `json:"retail_price"`
	Id          int     `json:"id"`
	UserId      int     `json:"user_id"`
	GoodsId     int     `json:"goods_id"`
	AddTime     string  `json:"add_time"`
	AddTimeInt  int64   `json:"add_time_int"`
}

//获取足迹数据
func (f *FootprintController) GetFootprintList() {
	//登录人id
	userId := f.user.Id
	var footPrintService = new(services.FootPrintService)
	lists := footPrintService.GetFootprintList(userId)
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)
	beforeYesterday := yesterday.Add(-24 * time.Hour)
	var resultDataTwo = make(map[int64][]IndexFootprintData)
	var timeIntarrs []int
	for _, v := range lists {
		timeString := v.AddTime.Format(TimeDateFormat)
		//日期时间戳
		timeInt := utils.StringTimeTransToTime(timeString).Unix()

		switch {
		case utils.DateEqual(v.AddTime, now):
			timeString = "今天"
		case utils.DateEqual(v.AddTime, yesterday):
			timeString = "昨天"
		case utils.DateEqual(v.AddTime, beforeYesterday):
			timeString = "前天"
		}

		if _, ok := resultDataTwo[timeInt]; !ok {
			timeIntarrs = append(timeIntarrs, int(timeInt))
		}
		resultDataTwo[timeInt] = append(resultDataTwo[timeInt], IndexFootprintData{
			Name:        v.Goods.Name,
			ListPicUrl:  v.Goods.ListPicUrl,
			GoodsBrief:  v.Goods.GoodsBrief,
			RetailPrice: v.Goods.RetailPrice,
			Id:          v.Id,
			UserId:      v.UserId,
			GoodsId:     v.Goods.Id,
			AddTime:     timeString,
			AddTimeInt:  v.AddTime.Unix(),
		})
	}
	var finalData [][]IndexFootprintData
	sort.Ints(timeIntarrs)
	timeIntarrs = utils.ReverseSliceInt(timeIntarrs)
	for _, v := range timeIntarrs {
		listdata := resultDataTwo[int64(v)]
		sort.Slice(listdata, func(i, j int) bool {
			return listdata[i].AddTimeInt > listdata[j].AddTimeInt
		})
		finalData = append(finalData, listdata)
	}
	utils.HttpSuccess(finalData, f.Ctx)
}
