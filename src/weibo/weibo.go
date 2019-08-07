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
