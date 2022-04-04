package utils

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"time"
)

//检测key是否存在map当中
func KeyInMapStringInt(m map[string]int, s string) bool {
	_, ok := m[s]
	return ok
}

//删除[]string 空元素
func RemoveEmpSliceString(s []string) []string {
	if len(s) == 0 {
		return nil
	}
	var result []string
	for _, v := range s {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}

//获取当前时间
func GetNowTime() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

//uuid生成唯一的串
func GenarateUUid() string {
	return uuid.NewV4().String()
}

//将[]orm.Parms 转 []int
func TransMapValueToSliceInt(data []orm.Params) []int {
	var res []int
	for _, item := range data {
		if val, ok := item["Id"].(int64); ok {
			res = append(res, int(val))
		}
		if val, ok := item["Id"].(int); ok {
			res = append(res, val)
		}
	}
	return res
}

//将[]orm.Parms 转 []int 获取对应的key 值的 value
func TransMapValueToSliceIntWithKey(data []orm.Params, key string) []int {
	var res []int
	if len(data) == 0 {
		return nil
	}
	if _, ok := data[0][key]; !ok {
		return nil
	}
	for _, item := range data {
		if val, ok := item[key].(int64); ok {
			res = append(res, int(val))
			continue
		}
		if val, ok := item[key].(int); ok {
			res = append(res, val)
			continue
		}
		if val, ok := item[key].(string); ok {
			intData, _ := strconv.Atoi(val)
			res = append(res, intData)
			continue
		}
	}
	return res
}

//将string 转 int
func TransStringToInt(value string) int {
	res, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return res
}

//将int转string
func TransIntToString(value int) string {
	return strconv.Itoa(value)
}

//将[]string 转 []int
func TransSliceStingToSliceInt(value []string) []int {
	var res []int
	for _, v := range value {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			return nil
		}
		res = append(res, vInt)
	}
	return res
}

//判断int 在[]int 种是否存在
func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func StringTimeTransToTime(stime string) time.Time {
	myTime, _ := time.ParseInLocation("2006-01-02", stime, time.Local)
	return myTime
}

func ReverseSliceInt(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

//返回一个string类型的值
func ReturnStringTypeValue(data interface{}) string {
	if str, ok := data.(string); ok {
		return str
	}
	return ""
}

//返回一个int类型的值

func ReturnIntTypeValue(data interface{}) int {
	switch v := data.(type) {
	case int:
		return v
	case int64:
		return int(v)
	case string:
		if vData, err := strconv.Atoi(v); err == nil {
			return vData
		} else {
			return 0
		}
	default:
		return 0
	}
}

//将[]orm.Params 转 map[int]stirng
//int 对应 key在map中对应的值
//string 对应value 在map中对应的值
func TransSliceOrmParamsToMapIntString(data []orm.Params, key, value string) map[int]string {
	if len(data) == 0 {
		return nil
	}
	if _, ok := data[0][key]; !ok {
		return nil
	}
	if _, ok := data[0][value]; !ok {
		return nil
	}
	resultData := make(map[int]string)
	for _, v := range data {
		resultDataKey := ReturnIntTypeValue(v[key])
		resultDataValue := ReturnStringTypeValue(v[value])
		resultData[resultDataKey] = resultDataValue
	}
	return resultData
}

//获取当前年份
func GetNowYear() string {
	year := time.Now().Year()
	return fmt.Sprintf("%02s", TransIntToString(year))
}

//获取当前月份
func GetNowMonth() string {
	month := time.Now().Month()
	return fmt.Sprintf("%02s", TransIntToString(int(month)))
}

//获取当前天
func GetNowDay() string {
	day := time.Now().Day()
	return fmt.Sprintf("%02s", TransIntToString(day))
}

//获取当前时
func GetNowHour() string {
	hour := time.Now().Hour()
	return fmt.Sprintf("%02s", TransIntToString(hour))
}

//获取当前分
func GetNowMinute() string {
	year := time.Now().Year()
	return fmt.Sprintf("%02s", TransIntToString(year))
}

//获取当前秒
func GetNowSecond() string {
	second := time.Now().Second()
	return fmt.Sprintf("%02s", TransIntToString(second))
}
