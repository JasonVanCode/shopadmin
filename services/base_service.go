package services

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"math"
	"strings"
)

type BaseService struct {
	PageData
}

type PageData struct {
	NumsPerPage int         `json:"pageSize"`
	CurrentPage int         `json:"currentPage"`
	Count       int         `json:"count"`
	TotalPages  int         `json:"totalPages"`
	Data        interface{} `json:"data"`
}

func (b *BaseService) GetTotalPage(goodsCount, size int) int {
	allPage := float64(goodsCount) / float64(size)
	return int(math.Ceil(allPage))
}

//讲orm.params key 转小写
func (b *BaseService) TransMapKeyToLower(data []orm.Params) []orm.Params {
	if len(data) == 0 {
		return nil
	}
	var resultData []orm.Params
	for _, v := range data {
		newParam := make(orm.Params)
		for k2, v2 := range v {
			newParam[strings.ToLower(k2)] = v2
		}
		resultData = append(resultData, newParam)
	}
	return resultData
}

//处理map转json key的问题
func (b *BaseService) updateJsonKey(value []orm.Params) []orm.Params {
	for _, val := range value {
		for k, v := range val {
			switch k {
			case "Id":
				delete(val, "Id")
				val["id"] = v
			case "ScenePicUrl":
				delete(val, "ScenePicUrl")
				val["scene_pic_url"] = v
			case "Subtitle":
				delete(val, "Subtitle")
				val["subtitle"] = v
			case "Title":
				delete(val, "Title")
				val["title"] = v
			case "PriceInfo":
				delete(val, "PriceInfo")
				val["price_info"] = v
			}
		}

	}
	return value
}

//获取查询数据数量
func (b *BaseService) GetCounts(qs orm.QuerySeter, cond *orm.Condition) (int, error) {
	count, err := qs.SetCond(cond).Count()
	return int(count), err
}

//错误日志处理
func (*BaseService) ErrLog(err error) {
	logs.Error(err.Error() + "\n")
}
