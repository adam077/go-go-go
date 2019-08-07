package scheduler

import "testing"

func TestRunner(t *testing.T) {
	t2()
}

func t1() {
	HotSpotRunner{}.Run()
}

func t2() {
	FollowWeibo{}.Run()
}
