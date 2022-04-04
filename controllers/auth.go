package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
	"time"
)

type AuthController struct {
	BaseController
}

//使用微信登录
func (a *AuthController) LoginByWeixin() {
	var authBody services.AuthBody
	if err := json.Unmarshal(a.Ctx.Input.RequestBody, &authBody); err != nil {
		utils.HttpFail(500, err.Error(), a.Ctx)
	}
	//处理微信授权登录，解析加密的用户信息
	userIinfo, err := services.Login(authBody.Code, authBody.UserInfo)
	if err != nil {
		utils.HttpFail(500, err.Error(), a.Ctx)
	}
	fmt.Println(userIinfo)
	//获取用户的请求ip地址
	clientIp := a.Ctx.Input.IP()
	//根据openid 去查询是否存在该用户
	var user models.Nideshop_user
	o := orm.NewOrm()
	err = o.QueryTable(new(models.Nideshop_user)).Filter("weixin_openid", userIinfo.OpenId).One(&user)
	if err == orm.ErrNoRows {
		//新建数据
		newUser := models.Nideshop_user{
			Username:     utils.GenarateUUid(),
			RegisterIp:   clientIp,
			WeixinOpenid: userIinfo.OpenId,
			Nickname:     userIinfo.NickName,
			Mobile:       "",
			Avatar:       userIinfo.AvatarUrl,
			Gender:       userIinfo.Gender,
			Password:     "",
			RegisterTime: time.Now(),
		}
		_, err := o.Insert(&newUser)
		if err != nil {
			utils.HttpFail(500, err.Error(), a.Ctx)
		}
		//查询这条新建的数据
		o.QueryTable(new(models.Nideshop_user)).Filter("weixin_openid", userIinfo.OpenId).One(&user)
	}

	//将用户的信息保存在map中
	userMap := map[string]interface{}{
		"id":       user.Id,
		"username": user.Username,
		"nickname": user.Nickname,
		"gender":   user.Gender,
		"avatar":   user.Avatar,
		"birthday": user.Birthday,
	}
	fmt.Println(userMap)
	user.LastLoginTime = time.Now()
	user.LastLoginIp = clientIp
	//更新表表数据
	if _, err := o.Update(&user); err != nil {
		utils.HttpFail(500, err.Error(), a.Ctx)

	}
	//生成jwt token
	var token string
	if token, err = services.GenerateToken(user.Id, user.Username, 0); err != nil {
		utils.HttpFail(500, err.Error(), a.Ctx)
	}
	utils.HttpSuccess(map[string]interface{}{
		"token":    token,
		"userInfo": userMap,
	}, a.Ctx)
}
