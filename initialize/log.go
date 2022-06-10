package initialize

import (
	"github.com/beego/beego/v2/core/logs"
)

//初始化日志
func init() {
	f := &logs.PatternLogFormatter{
		Pattern:    "%F:%n|%w%t>> %m\n",
		WhenFormat: "2006-01-02",
	}
	logs.RegisterFormatter("pattern", f)
	_ = logs.SetGlobalFormatter("pattern")
	_ = logs.SetLogger(logs.AdapterFile, `{"filename":"logs/log.log","formatter": "pattern"}`)
}
