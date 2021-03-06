package weibo

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go-go-go/src/utils"
	"strconv"
)

func GetWeiboTopic(cookie string) (map[string]map[int]string, map[string]string) {
	/*
		获取微博话题
	*/
	var err error
	errList := make(map[string]string, 0)
	result := make(map[string]map[int]string)
	for cat, name := range m {
		result[name], err = GetWeiboTopicForOne(cookie, cat, 100)
		if err != nil {
			errList[m[cat]] = err.Error()
		}
	}
	return result, errList
}

func GetWeiboTopicForOne(cookie, cat string, pages int) (map[int]string, error) {
	/*
		获取微博话题列表
	*/
	var err error
	var a []byte
	result := make(map[int]string)
	rank := 0
	for i := 1; i <= pages; i++ {
		var urlStr = "https://d.weibo.com/231650_ctg1_-_" + cat + "?Pl_Discover_Pt6Rank__4_page=" + strconv.Itoa(i)
		a, err = utils.QueryGet(urlStr, nil, map[string]string{"Cookie": cookie})
		if err != nil {
			log.Warn().Msg(err.Error())
			a, err = utils.QueryGet(urlStr, nil, map[string]string{"Cookie": cookie})
		}
		topics := utils.FindBetween(a, " alt=\\\"#", "#\\\" class=\\\"pic")
		for _, topic := range topics {
			rank += 1
			result[rank] = topic
		}
		if len(topics) < 15 {
			break
		}
	}
	if cat == "all" && len(result) < 1 {
		err = errors.New(m[cat] + "no topics")
	}
	log.Info().Str("cat", cat).Int("num", len(result)).Msg("")
	return result, err

}

var m = map[string]string{
	"all":   "话题榜",
	"epoch": "新时代",
	"1":     "社会",
	"138":   "互联网",
	"3":     "科普",
	"131":   "数码",
	"146":   "音乐",
	"7":     "财经",
	"2":     "明星",
	"102":   "综艺",
	"101":   "电视剧",
	"100":   "电影",
	"117":   "汽车",
	"98":    "体育",
	"111":   "运动健身",
	"113":   "健康",
	"144":   "军事",
	"123":   "美图",
	"5":     "情感",
	"140":   "搞笑",
	"126":   "游戏",
	"93":    "旅游",
	"116":   "育儿",
	"133":   "教育",
	"91":    "美食",
	"6":     "公益",
	"137":   "房产",
	"145":   "星座",
	"94":    "读书",
	"142":   "艺术",
	"114":   "时尚美妆",
	"97":    "动漫",
	"128":   "萌宠",
	"120":   "生活记录",
	"9":     "创意征集",
	"8":     "其他",
}
