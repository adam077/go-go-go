package scheduler

import (
	"go-go-go/src/data"
	"testing"
)

func TestRunner(t *testing.T) {
	//tt()
	t3()
}

func t1() {
	HotSpotRunner{}.Run()
}

func t2() {
	//FollowWeibo{}.Run()
	//WeiboChat{}.Run()
	//WeiboTopicRunner{}.Run()
	//WeiboLoginChecker{1}.Run()
	groupSender()
}

func t3() {
	//doForWord("婴儿手推车", taobao.Token, weibo.Cookie)
	taobaoToken := data.GetConfig("taobao_cookie")
	weiboCookie := data.GetConfig("weibo_cookie")
	doForWord(ItemType{"食品", []string{"零食"}, "19452"}, taobaoToken, weiboCookie)
}
