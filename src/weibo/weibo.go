package weibo

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go-go-go/src/utils"
	"net/url"
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

func Follow(uid, cookie string) error {
	headers := map[string]string{
		"Cookie":  cookie,
		"Referer": "https://weibo.com/u/" + uid,
	}

	form := url.Values{}
	form.Add("uid", uid)
	form.Add("refer_flag", "1005050001_")
	result, err := utils.QueryPost("https://weibo.com/aj/f/followed", headers, "x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	var codeResult FollowResp
	json.Unmarshal(result, &codeResult)
	if codeResult.Code != "100000" {
		log.Error().Str("uid", uid).Msg(codeResult.Code)
		return errors.New(codeResult.Code)
	}
	log.Info().Str("uid", uid).Msg("")
	return nil
}

type FollowResp struct {
	Code string `json:"code"`
}

func SendMessage(cookie, uid, msg string) error {
	headers := map[string]string{
		"Cookie":  cookie,
		"Referer": "https://api.weibo.com/chat/",
	}

	form := url.Values{}
	form.Add("text", msg)
	form.Add("uid", uid)
	form.Add("source", "209678993")
	result, err := utils.QueryPost("https://api.weibo.com/webim/2/direct_messages/new.json", headers, "x-www-form-urlencoded", strings.NewReader(form.Encode()))
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	var codeResult MsgResp
	json.Unmarshal(result, &codeResult)
	if codeResult.MsgStatus != 0 {
		errMsg := strconv.Itoa(codeResult.MsgStatus)
		log.Error().Str("uid", uid).Msg(errMsg)
		return errors.New(errMsg)
	}
	log.Info().Str("uid", uid).Msg("")
	return nil

}

type MsgResp struct {
	MsgStatus int `json:"msg_status"`
}

func GetUserStatus(cookie, uid string) (map[string]string, error) {
	pages := 100
	// 获取用户状态
	result := make(map[string]string)
	for i := 1; i <= pages; i++ {
		var urlStr = "https://weibo.com/p/100505" + uid + "/myfollow?Pl_Official_RelationMyfollow__95_page=" + strconv.Itoa(i)
		a, err := utils.QueryGet(urlStr, map[string]string{"Cookie": cookie})
		if err != nil {
			return result, err
		}
		uids := utils.FindBetween(a, "&uid=", "&sex")
		// 正则表达式中匹配一个反斜杠要用四个反斜杠  "<span class=\"S_txt1\">", "<\/span>")  但是这样又不好replace所以使用先替代再查找
		status := utils.FindBetween(a, "<span class=\\\"S_txt1\\\">", "<\\/span>")
		if len(uids) != len(status) {
			return result, errors.New("用户与状态的长度不一致")
		}
		for x := range uids {
			result[uids[x]] = status[x]
		}
		if len(uids) < 30 {
			break
		}
	}
	log.Info().Int("len", len(result)).Msg("")
	return result, nil
}
