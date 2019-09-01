package weibo

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"go-go-go/src/utils"
	"net/url"
	"strings"
)

func SendWeiBo(cookie, text string, picIds []string) (string, error) {
	/*
		发一个微博
	*/
	headers := map[string]string{
		"Cookie":  cookie,
		"Referer": "https://weibo.com/u/6491407817/home?wvr=5",
	}

	form := url.Values{}
	form.Add("text", text)
	if picIds != nil {
		form.Add("pic_id", strings.Join(picIds, "|"))
	}
	result, err := utils.QueryPostWithFormData("https://weibo.com/aj/mblog/add", nil, headers, form)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}
	var codeResult FollowResp
	json.Unmarshal(result, &codeResult)
	if codeResult.Code != "100000" {
		log.Error().Str("text", text).Msg(codeResult.Code)
		return "", errors.New(codeResult.Code)
	}
	status := utils.FindBetween(result, "mid=\\\"", "\\\"  action-type")
	if len(status) > 0 {
		return status[0], nil
	}
	return "", errors.New("no mid found")
}

func SendWeiBoComment(cookie, mid, comment string) error {
	/*
		发一个微博评论
	*/
	headers := map[string]string{
		"Cookie":  cookie,
		"Referer": "https://weibo.com/",
	}

	form := url.Values{}
	form.Add("mid", mid)
	form.Add("content", comment)
	result, err := utils.QueryPostWithFormData("https://weibo.com/aj/v6/comment/add", nil, headers, form)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	var codeResult FollowResp
	json.Unmarshal(result, &codeResult)
	if codeResult.Code != "100000" {
		log.Error().Str("mid", mid).Msg(string(result))
		return errors.New(codeResult.Code)
	}
	log.Info().Str("mid", mid).Msg("")
	return nil
}
