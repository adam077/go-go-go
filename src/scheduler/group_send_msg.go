package scheduler

import (
	"github.com/rs/zerolog/log"
	"go-go-go/src/data"
	"go-go-go/src/utils"
	"go-go-go/src/weibo"
)

type GroupSender struct {
}

func (runner GroupSender) Run() {
	defer utils.CommonRecover()
	if data.GetConfig(data.SchedulerWeiboGroupSender) == "" {
		return
	}
	log.Info().Msg("start GroupSender")
	groupSender()
}

func groupSender() {
	groups := weibo.GetGroupsFromContacts(weibo.Cookie)
	for _, group := range groups {
		weibo.SendMessageToGroup(weibo.Cookie, group, "有！粉！必！回！")
	}
}
