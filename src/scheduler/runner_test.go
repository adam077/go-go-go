package scheduler

import (
	"testing"
)

func TestRunner(t *testing.T) {
	//tt()
	t2()
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
