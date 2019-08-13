package weibo

import (
	"github.com/rs/zerolog/log"
	"go-go-go/src/utils"
	"strconv"
	"strings"
)

var splits = []string{"href=\"//weibo.com/", "?refer_flag"}

func GetUsers(topic string, cookie string, pages int) ([]string, error) {
	result := make([]string, 0)
	aMap := make(map[int]bool)
	for i := 1; i <= pages; i++ {
		a, err := utils.QueryGet("https://s.weibo.com/weibo/"+topic+"?topnav=1&wvr=6&b=1&page="+strconv.Itoa(i), map[string]string{"Cookie": cookie})
		if err != nil {
			return result, err
		}
		splitsResult := splitByss(string(a), splits)
		for x := range splitsResult {
			uid, err := strconv.Atoi(splitsResult[x])
			if err == nil {
				if _, ok := aMap[uid]; !ok {
					result = append(result, strconv.Itoa(uid))
				}
				aMap[uid] = true
			}
		}
	}
	log.Info().Int("len", len(result)).Msg("")
	return result, nil
}

func GetUsersFromHufen(cookie string, pages int) ([]string, error) {
	result := make([]string, 0)
	for i := 1; i <= pages; i++ {
		a, err := utils.QueryGet("https://s.weibo.com/realtime?q=%23%E4%BA%92%E7%B2%89%23&rd=realtime&tw=realtime&Refer=weibo_realtime&page="+strconv.Itoa(i), map[string]string{"Cookie": cookie})
		if err != nil {
			return result, err
		}
		uids := utils.FindBetween(a, "uid=", "&mid=")
		result = append(result, uids...)
	}
	log.Info().Int("len", len(result)).Msg("")
	return result, nil
}

func GetUsersFromCantSleep(cookie string, pages int) ([]string, error) {
	result := make([]string, 0)
	for i := 1; i <= pages; i++ {
		a, err := utils.QueryGet("https://s.weibo.com/realtime?q=%E7%9D%A1%E4%B8%8D%E7%9D%80&rd=realtime&tw=realtime&Refer=weibo_realtime&page="+strconv.Itoa(i), map[string]string{"Cookie": cookie})
		if err != nil {
			return result, err
		}
		uids := utils.FindBetween(a, "uid=", "&mid=")
		result = append(result, uids...)
	}
	log.Info().Int("len", len(result)).Msg("")
	return result, nil
}

func GetUsersFromRealTimeWord(word, cookie string, pages int) ([]string, error) {
	result := make([]string, 0)
	for i := 1; i <= pages; i++ {
		a, err := utils.QueryGet("https://s.weibo.com/realtime?q="+word+"&rd=realtime&tw=realtime&Refer=weibo_realtime&page="+strconv.Itoa(i), map[string]string{"Cookie": cookie})
		if err != nil {
			return result, err
		}
		uids := utils.FindBetween(a, "uid=", "&mid=")
		result = append(result, uids...)
	}
	log.Info().Int("len", len(result)).Msg("")
	return result, nil
}

func splitByss(str string, splits []string) []string {
	if len(splits) < 1 {
		return []string{str}
	} else if len(splits) == 1 {
		return strings.Split(str, splits[0])
	} else {
		result := make([]string, 0)
		temp := strings.Split(str, splits[len(splits)-1])
		for x := range temp {
			result = append(result, splitByss(temp[x], splits[0:len(splits)-1])...)
		}
		return result
	}
}

func GetUsersFronGroup() ([]string, error) {
	result := make([]string, 0)
	return result, nil
}