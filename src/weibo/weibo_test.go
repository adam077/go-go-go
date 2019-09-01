package weibo

import (
	"fmt"
	"testing"
	"time"
)

func TestWeibo(t *testing.T) {
	t4()
}

func t1() {
	uids, _ := GetUsers("互粉", cookie, 3)
	for _, uid := range uids {
		time.Sleep(10 * time.Second)
		err := Follow(uid, cookie)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func t2() {
	fmt.Println(GetUserStatus(cookie, "6491407817"))
}

func t3() {
	GetUsersFromHufen(cookie, 3)
}

func t4() {

	mid, _ := SendWeiBo(cookie, time.Now().String(), nil)
	SendWeiBoComment(cookie, mid, time.Now().String())
}

func t5() {
	fmt.Println(GetUrlPicId("https://wx2.sinaimg.cn/mw690/0075jhI5ly1g635ehui6bj30pr1f8101.jpg"))
}
