package utils

import (
	"encoding/json"
	"github.com/beego/beego/v2/server/web/context"
)

type HttpData struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

//http错误返回的数据
func HttpFail(errNo int, errMsg string, ctx *context.Context) {
	h := HttpData{
		errNo,
		errMsg,
		nil,
	}
	jsonStr, _ := json.Marshal(h)
	ctx.Output.JSON(json.RawMessage(jsonStr), true, false)
}

//http成功返回数据
func HttpSuccess(data interface{}, ctx *context.Context) {
	h := HttpData{
		0,
		"",
		data,
	}
	jsonStr, err := json.Marshal(h)
	if err != nil {
		ctx.Output.JSON(err.Error(), true, false)
	}

	ctx.Output.JSON(json.RawMessage(jsonStr), true, false)

}
