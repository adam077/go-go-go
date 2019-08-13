package scheduler

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go-go-go/src/data"
	"go-go-go/src/ding-talk"
	"go-go-go/src/utils"
	"go-go-go/src/weibo"
)

type WeiboLoginChecker struct {
	doUpdate int
}

func (runner WeiboLoginChecker) Run() {
	defer utils.CommonRecover()
	if data.GetConfig(data.SchedulerWeiboLoginChecker) == "" {
		return
	}
	log.Info().Msg("start WeiboLoginChecker")
	userDatas := data.GetWeiboUserFollow("")
	for _, me := range userDatas {
		shouldbe15, _ := weibo.GetWeiboTopicForOne(me.Cookie, "all", 1)
		if len(shouldbe15) < 1 {
			ding_talk.SendDingMessage("c594857d21850991a7a15de920ab3c69b626fbb5d0f7bd3671dc0861bf13fab3", "loginFail "+me.Name)
			if runner.doUpdate > 0 {
				TriggerLogin(me)
				//var urlStr = "http://o26166b137.wicp.vip:44082/lv1/lv2/trigger_login?name=" + me.Name
				//a, _ := utils.QueryGet(urlStr, nil)
				//ding_talk.SendDingMessage("c594857d21850991a7a15de920ab3c69b626fbb5d0f7bd3671dc0861bf13fab3", "loginService triggered "+string(a))
			}
		}
	}
}

func TriggerLogin(one *data.WeiboUser) {
	one.Cookie = weibo.GetCookie(one.LoginName, one.Password)
	ding_talk.SendDingMessage("c594857d21850991a7a15de920ab3c69b626fbb5d0f7bd3671dc0861bf13fab3", "cookie gen "+one.Cookie)
	err := data.UpdateCookie(one.ID, one.Cookie)
	fmt.Println(err)
}
