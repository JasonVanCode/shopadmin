package services

import (
	"encoding/base64"
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
	"shopadmin/utils"
)

type CommentService struct {
	BaseService
}

//[]*models.NideshopComment
//获取当前商品的评价数
type CommentCount struct {
	CountNum int
}

type CommenListtRtnJson struct {
	Content  string                           `json:"content"`
	TypeId   int                              `json:"type_id"`
	ValueId  int                              `json:"value_id"`
	Id       int                              `json:"id"`
	AddTime  string                           `json:"add_time"`
	UserInfo map[string]interface{}           `json:"user_info"`
	PicList  []*models.NideshopCommentPicture `json:"pic_list"`
}

func (c *CommentService) GetCommentList(pageData map[string]int, getData map[string]string) PageData {
	var resultJson []CommenListtRtnJson
	size := pageData["size"]
	page := pageData["page"]
	typeId := getData["typeId"]
	showType := getData["showType"] // 0所有评论  1 带图片的评论
	valueId := getData["valueId"]
	o := orm.NewOrm()
	var sqlResult []orm.Params
	var selectSql string
	//获取当前评论数量
	var countNum = c.GetCommentCount(showType, valueId, typeId)
	if countNum == 0 {
		return PageData{
			size,
			page,
			0,
			0,
			nil,
		}
	}
	////下面是sql信息
	//如果只要图片信息
	if showType == "1" {
		selectSql = "SELECT c.* ,u.username,u.nickname,u.avatar from nideshop_comment as c INNER JOIN nideshop_comment_picture as p on p.comment_id = c.id INNER JOIN nideshop_user as u on u.id = c.user_id where c.value_id = ? and c.type_id = ? limit ?,?"
	} else {
		selectSql = "select c.*,u.username,u.nickname,u.avatar from nideshop_comment as c INNER JOIN nideshop_user as u on u.id = c.user_id where c.value_id = ? and c.type_id = ? limit ?,?"
	}
	if _, err := o.Raw(selectSql, valueId, typeId, (page-1)*size, size).Values(&sqlResult); err != nil {
		return PageData{
			size,
			page,
			0,
			0,
			nil,
		}
	}
	commentIds := utils.TransMapValueToSliceIntWithKey(sqlResult, "id")
	//获取存在评论的图片
	_, commentPics := c.GetCommentPics(commentIds)

	for _, v := range sqlResult {
		commentId := utils.ReturnIntTypeValue(v["id"])
		content := utils.ReturnStringTypeValue(v["content"])
		if content != "" {
			decodeContent, _ := base64.StdEncoding.DecodeString(content)
			content = string(decodeContent)
		}
		userInfo := map[string]interface{}{
			"username": v["username"],
			"nickname": v["nickname"],
			"avatar":   v["avatar"],
		}
		var pics []*models.NideshopCommentPicture
		if picItem, ok := commentPics[commentId]; ok {
			pics = picItem
		}
		resultJson = append(resultJson, CommenListtRtnJson{
			Content:  content,
			TypeId:   utils.ReturnIntTypeValue(v["type_id"]),
			ValueId:  utils.ReturnIntTypeValue(v["value_id"]),
			Id:       commentId,
			AddTime:  utils.ReturnStringTypeValue(v["add_time"]),
			UserInfo: userInfo,
			PicList:  pics,
		})
	}
	return PageData{
		size,
		page,
		countNum,
		(countNum + size - 1) / size,
		resultJson,
	}

}

//获取评论数量
//typ 0 全部评论数量  typ 1 有图片的数量
func (c *CommentService) GetCommentCount(typ, valueId, typeId string) int {
	var count CommentCount
	var sqlString string
	o := orm.NewOrm()
	if typ == "0" {
		sqlString = "SELECT count(*) as count_num from nideshop_comment where value_id = ? and type_id = ?"
	} else {
		sqlString = "SELECT count(c.id) as count_num from nideshop_comment as c INNER JOIN nideshop_comment_picture as p on p.comment_id = c.id where c.value_id = ? and c.type_id = ?"
	}
	if err := o.Raw(sqlString, valueId, typeId).QueryRow(&count); err != nil || count.CountNum == 0 {
		return 0
	}
	return count.CountNum
}

//返回有图片的comment id 以及图片数据
func (c *CommentService) GetCommentPics(commentIds []int) ([]int, map[int][]*models.NideshopCommentPicture) {
	if len(commentIds) == 0 {
		return nil, nil
	}
	var commentPics []*models.NideshopCommentPicture
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopCommentPicture)).Filter("comment_id__in", commentIds).RelatedSel().All(&commentPics)

	if err != nil || num == 0 {
		return nil, nil
	}
	var mapCommentPics = make(map[int][]*models.NideshopCommentPicture)

	var Ids []int
	for _, v := range commentPics {
		if !utils.ContainsInt(Ids, v.Comment.Id) {
			Ids = append(Ids, v.Comment.Id)
		}
		if pics, ok := mapCommentPics[v.Comment.Id]; ok {
			pics = append(pics, v)
		} else {
			mapCommentPics[v.Comment.Id] = []*models.NideshopCommentPicture{
				v,
			}
		}
	}
	return Ids, mapCommentPics
}
