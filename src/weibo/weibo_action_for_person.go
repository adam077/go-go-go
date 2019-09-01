package weibo

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"go-go-go/src/utils"
	"net/url"
	"strconv"
)

func Follow(uid, cookie string) error {
	/*
		关注一个人
	*/
	headers := map[string]string{
		"Cookie":  cookie,
		"Referer": "https://weibo.com/u/" + uid,
	}

	form := url.Values{}
	form.Add("uid", uid)
	form.Add("refer_flag", "1005050001_")
	result, err := utils.QueryPostWithFormData("https://weibo.com/aj/f/followed", nil, headers, form)
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
	/*
		发送私信
	*/
	headers := map[string]string{
		"Cookie":  cookie,
		"Referer": "https://api.weibo.com/chat/",
	}

	form := url.Values{}
	form.Add("text", msg)
	form.Add("uid", uid)
	form.Add("source", "209678993")
	result, err := utils.QueryPostWithFormData("https://api.weibo.com/webim/2/direct_messages/new.json", nil, headers, form)
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

func SendMessageToGroup(cookie, group, msg string) error {
	/*
		发群消息
	*/
	headers := map[string]string{
		"Cookie":  cookie,
		"Referer": "https://api.weibo.com/chat/",
	}

	form := url.Values{}
	form.Add("content", msg)
	form.Add("id", group)
	form.Add("source", "209678993")
	result, err := utils.QueryPostWithFormData("https://api.weibo.com/webim/groupchat/send_message.json", nil, headers, form)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	var codeResult MsgResp
	json.Unmarshal(result, &codeResult)
	if codeResult.MsgStatus != 0 {
		errMsg := strconv.Itoa(codeResult.MsgStatus)
		log.Error().Str("group", group).Msg(errMsg)
		return errors.New(errMsg)
	}
	log.Info().Str("group", group).Msg("")
	return nil

}
