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

var pages = 3
var seconds = 7

func (runner FollowWeibo) Run() {
	d := data.GetConfig("weibo_follow")
	if d != "" {
		do()
	} else {
		chatId := getDingChatId()
		if chatId != "" {
			go ding_talk.SendDingMessage(chatId, "忽略 follow")
		}
	}
}

func do() {
	log.Info().Msg("start follow")
	userDatas := data.GetWeiboUserFollow("NobodyHu")
	for _, userData := range userDatas {
		var followMap = make(map[string]bool)
		for _, follower := range userData.WeiboFollows {
			followMap[follower.UID] = true
		}
		succNum := 0
		fail := 0
		uids, _ := weibo.GetUsers("互粉", userData.Cookie, pages)
		for _, uid := range uids {
			if _, ok := followMap[uid]; ok {
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
				fail += 1
				if fail >= 2 {
					chatId := getDingChatId()
					if chatId != "" {
						go ding_talk.SendDingMessage(chatId, fmt.Sprintf("%s follow break for %s", uid, err.Error()))
					}
					break
				}
			}
		}

		chatId := getDingChatId()
		if chatId != "" {
			go ding_talk.SendDingMessage(chatId, fmt.Sprintf("新增加关注 %d", succNum))
		}

	}
}

func getDingChatId() string {
	chatId := ""
	var dings = data.GetDingChatId("own")
	if len(dings) > 0 {
		chatId = dings[0].ChatId
	}
	return chatId
}

func getFollowMap() map[string]bool {
	// 查询
	return map[string]bool{}
}

func getCookie() string {
	//return cookie
	return data.GetConfig("weibo_cookie")
}
