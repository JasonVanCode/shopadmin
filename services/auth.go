package services

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	beego "github.com/beego/beego/v2/server/web"
	"shopadmin/utils"
)

//下面是微信授权登录请求的数据
type AuthBody struct {
	Code     string `json:"code"`
	UserInfo `json:"userInfo"`
}

type UserInfo struct {
	ErrMsg        string `json:"errMsg"`
	RawData       string `json:"rawData"`
	Signature     string `json:"signature"`
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
	CloudID       string `json:"cloudID"`
	RealUserInfo  `json:"userInfo"`
}

type RealUserInfo struct {
	NickName  string `json:"nickName"`
	Gender    int    `json:"gender"`
	Language  string `json:"language"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	AvatarUrl string `json:"avatarUrl"`
	OpenId    string `json:"openId"`
	WaterMark `json:"watermark"`
}

type WaterMark struct {
	Timestamp int    `json:"timestamp"`
	Appid     string `json:"appid"`
}

//https://api.weixin.qq.com/sns/jscode2session 接口返回的数据
type WXLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

//测试aes 加密解密
//st := utils.AesCBCEncrypt([]byte("aaaaa"), []byte("AAAAAAAAAAAAAAAA"), []byte("AAAAAAAAAAAAAAAA"))
//st1, _ := base64.StdEncoding.DecodeString(st)
//ddddd := utils.AesCBCDecrypt(st1, []byte("AAAAAAAAAAAAAAAA"), []byte("AAAAAAAAAAAAAAAA"))
//fmt.Println(string(ddddd))

//处理微信的登录请求
func Login(code string, info UserInfo) (*RealUserInfo, error) {

	appid, _ := beego.AppConfig.String("weixin::appid")
	secret, _ := beego.AppConfig.String("weixin::secret")
	//调用jscode2session 接口
	//成功：{"session_key":"DZReJGxqHtNpjFjDmiVkAQ==","openid":"orQu25OVzhAgieJZwThAH-iJ59-g"}
	//失败：{"errcode":40029,"errmsg":"invalid code, rid: 62315303-6bf82c14-411401fc"}
	str := httplib.Get("https://api.weixin.qq.com/sns/jscode2session")
	//请求参数
	str.Param("grant_type", "authorization_code")
	str.Param("js_code", code)
	str.Param("secret", secret)
	str.Param("appid", appid)
	//将返回数据解析
	var wxRes WXLoginResponse
	str.ToJSON(&wxRes)
	//通过请求的userInfo 里面的rowdata 和 jscode2session 接口返回的 session_key sha1 加密
	//返回的结果和 userInfo 的Signature 进行比较
	sha := sha1.New()
	sha.Write([]byte(info.RawData + wxRes.SessionKey))
	//判断数据是否一致
	if info.Signature != hex.EncodeToString(sha.Sum(nil)) {
		fmt.Println("验证失败")
	}
	//通过AES解密出ncryptedData里面的真实数据
	return DecryptUserInfoData(wxRes.SessionKey, info.Iv, info.EncryptedData)
}

//aes 解密 数据
func DecryptUserInfoData(sessionKey string, iv string, encryptedData string) (*RealUserInfo, error) {
	sessionKeyByte, _ := base64.StdEncoding.DecodeString(sessionKey)
	ivByte, _ := base64.StdEncoding.DecodeString(iv)
	ncryptedDataByte, _ := base64.StdEncoding.DecodeString(encryptedData)

	userIinfoByte, err := utils.AesCBCDecrypt(ncryptedDataByte, sessionKeyByte, ivByte)
	if err != nil {
		return nil, err
	}

	var userInfo RealUserInfo
	if err := json.Unmarshal(userIinfoByte, &userInfo); err != nil {
		return nil, err
	}
	return &userInfo, nil
}
