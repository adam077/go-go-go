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
	uids, _ := GetUsers("互粉", Cookie, 3)
	for _, uid := range uids {
		time.Sleep(10 * time.Second)
		err := Follow(uid, Cookie)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func t2() {
	fmt.Println(GetUserStatus(Cookie, "6491407817"))
}

func t3() {
	GetUsersFromHufen(Cookie, 3)
}

func t4() {

	mid, _ := SendWeiBo(Cookie, time.Now().String(), nil)
	SendWeiBoComment(Cookie, mid, time.Now().String())
}

func t5() {
	fmt.Println(GetUrlPicId("https://wx2.sinaimg.cn/mw690/0075jhI5ly1g635ehui6bj30pr1f8101.jpg"))
}
