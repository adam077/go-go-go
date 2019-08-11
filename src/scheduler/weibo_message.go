package scheduler

import (
	"github.com/rs/zerolog/log"
	"go-go-go/src/data"
	"go-go-go/src/weibo"
	"math/rand"
	"time"
)

type WeiboChat struct {
}

func (runner WeiboChat) Run() {
	log.Info().Msg("start follow")
	userDatas := data.GetWeiboUserFollow("")
	for _, userData := range userDatas {
		cookie := weibo.GetCookie(userData.LoginName, userData.Password)
		chated := data.GetWeiboUserChat(userData.ID)
		chatedMap := make(map[string]bool)
		for _, chatedOne := range chated {
			chatedMap[chatedOne.UID] = true
		}
		uids, _ := weibo.GetUserStatus(cookie, userData.Uid)
		for uid, status := range uids {
			if _, ok := chatedMap[uid]; ok {
				continue
			}
			if status == "已关注" {
				var s = 1 + rand.Intn(3)
				time.Sleep(time.Duration(s) * time.Second)
				huFen(userData.ID, cookie, uid)
			}
		}
	}
}

func huFen(id, cookie, uid string) {
	err := weibo.SendMessage(cookie, uid, "哈喽，欢迎互粉哦~")
	if err != nil {
		log.Warn().Msg("aaa")
	} else {
		data.UpdateWeiboChat([]data.WeiboChat{
			{
				WeiboUserId: id,
				UID:         uid,
			},
		})
	}
}
