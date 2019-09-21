package utils

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func CommonRecover() {
	if err := recover(); err != nil {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)

		if err, ok := err.(error); ok {
			log.Error().Str("error", err.(error).Error()).Str("stack", string(buf[:n])).Msg("goroutine unexpected panic.")
		} else {
			log.Error().Str("stack", string(buf[:n])).Msg("goroutine unexpected error when recover.")
		}
	}
}

func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func KeyDiff(left, right map[string]interface{}) bool {
	if len(left) != len(right) {
		return false
	}
	for key := range left {
		if _, ok := right[key]; !ok {
			return false
		}
	}
	return true
}

var ShangHaiLocation, _ = time.LoadLocation("Asia/Shanghai")

func GetLocalTimeFromShanghaiString(raw string) time.Time {
	result, _ := time.ParseInLocation(TIME_SHANGHAI_FORMAT, raw, ShangHaiLocation)
	return result.Local()
}

func GetShanghaiTimeString(raw time.Time) string {
	return raw.In(ShangHaiLocation).Format(TIME_SHANGHAI_FORMAT)
}

const (
	TIME_SHANGHAI_FORMAT = "2006-01-02T15:04:05+08:00"
	DATE_FORMAT          = "2006-01-02"
	TIME_DEFAULT_FORMAT  = "2006-01-02 15:04:05"
)

func GetTimeDateString(t time.Time) string {
	// 时间转化为日期字符串
	return t.Format(DATE_FORMAT)
}

func GetNowTime() (int64, string, int, int) {
	now := time.Now()
	today := GetTimeDateString(now)
	hour := now.Hour()
	min := (now.Minute() / 5) * 5
	return now.Unix(), today, hour, min
}

func GetUUID() string {
	return uuid.NewV4().String()
}

func FindBetween(str []byte, left, right string) []string {

	// result := make([]string, 0)
	// // var r = fmt.Sprintf("%s[^%s]+%s", left, right, right) // 这个只适合单个

	// var r = fmt.Sprintf("%s.*?%s", left, right) // 懒惰限定符
	// reg := regexp.MustCompile(r)
	// for _, one := range reg.FindAll(str, -1) {
	// 	result = append(result, strings.Replace(strings.Replace(string(one), left, "", -1), right, "", -1))
	// }
	// // https://www.cnblogs.com/jkko123/p/8329515.html
	// // https://www.cnblogs.com/golove/p/3270918.html
	// // https://www.runoob.com/regexp/regexp-syntax.html
	// // https://www.debuggex.com/
	// return result

	return FindBetweenInstr(string(str), left, right)
}

func FindBetweenInstr(str string, left, right string) []string {
	newLeft := "aaaLeft"
	newRight := "aaaRight"
	str = strings.Replace(str, left, newLeft, -1)
	str = strings.Replace(str, right, newRight, -1)

	result := make([]string, 0)
	// var r = fmt.Sprintf("%s[^%s]+%s", left, right, right) // 这个只适合单个

	var r = fmt.Sprintf("%s.*?%s", newLeft, newRight) // 懒惰限定符
	reg := regexp.MustCompile(r)
	for _, one := range reg.FindAll([]byte(str), -1) {
		result = append(result, strings.Replace(strings.Replace(string(one), newLeft, "", -1), newRight, "", -1))
	}
	// https://www.cnblogs.com/jkko123/p/8329515.html
	// https://www.cnblogs.com/golove/p/3270918.html
	// https://www.runoob.com/regexp/regexp-syntax.html
	// https://www.debuggex.com/
	return result
}
