package services

import (
	"github.com/beego/beego/v2/client/orm"
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
