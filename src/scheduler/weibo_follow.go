package scheduler

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go-go-go/src/data"
	"go-go-go/src/ding-talk"
	"go-go-go/src/weibo"
	"math/rand"
	"time"
)

type FollowWeibo struct {
}

var pages = 20
var seconds = 10

func (runner FollowWeibo) Run() {
	do := data.GetConfig("weibo_follow")
	chatId := getDingChatId()
	if do == "" {
		log.Info().Msg("start follow")
		userDatas := data.GetWeiboUserFollow("")
		for _, userData := range userDatas {
			var followMap = make(map[string]bool)
			for _, follower := range userData.WeiboFollows {
				followMap[follower.UID] = true
			}
			succNum := 0
			fail := 0
			uids, _ := weibo.GetUsers("互粉", userData.Cookie, pages)
			//uids, _ := weibo.GetUsersFromHufen(userData.Cookie, pages)
			//uids, _ := weibo.GetUsersFromCantSleep(userData.Cookie, pages)
			for _, uid := range uids {
				if _, ok := followMap[uid]; ok {
					log.Info().Str("uid", uid).Msg("followed")
					continue
				}
				var s = seconds + rand.Intn(3)
				time.Sleep(time.Duration(s) * time.Second)
				err := weibo.Follow(uid, userData.Cookie)
				if err == nil {
					data.UpdateWeiboFollower([]data.WeiboFollow{
						{
							WeiboUserId: userData.ID,
							UID:         uid,
						},
					})
					followMap[uid] = true
					succNum += 1
				} else {
					if fail += 1; fail >= 2 {
						if chatId != "" {
							ding_talk.SendDingMessage(chatId, fmt.Sprintf("%s follow break for %s", uid, getCodeMsg(err.Error())))
						}
						break
					}
				}
			}
			if chatId != "" {
				ding_talk.SendDingMessage(chatId, fmt.Sprintf("新增加关注 %d", succNum))
			}
		}
	} else {
		if chatId != "" {
			ding_talk.SendDingMessage(chatId, "忽略 follow")
		}
	}
}

func getCodeMsg(code string) string {
	if _, ok := codeMap[code]; ok {
		return codeMap[code]
	}
	return code
}

var codeMap = map[string]string{
	"100001": "关注太到今日上限了",
	"100027": "关注要验证码了",
}

func getDingChatId() string {
	chatId := ""
	var dings = data.GetDingChatId("own")
	if len(dings) > 0 {
		chatId = dings[0].ChatId
	}
	return chatId
}
