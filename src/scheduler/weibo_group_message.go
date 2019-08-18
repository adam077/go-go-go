package scheduler

import (
	"fmt"
	"go-go-go/src/weibo"
	"math/rand"
	"time"
)

// 4074460921980762 抛弃

func tt() {
	users := weibo.GetUsersFromGroup("4074460921980762", WeiboCookie)
	var i = 0
	followMap, _ := weibo.GetUserStatus(WeiboCookie, "6491407817")
	for _, userUid := range users {
		if _, ok := followMap[userUid]; ok {
			fmt.Println("continue " + userUid)
			continue
		}
		var sleep = 2 + rand.Intn(4)
		time.Sleep(time.Duration(sleep) * time.Second)
		err := weibo.Follow(userUid, WeiboCookie)
		if err != nil {
			fmt.Println(err)
			break
		}
		weibo.SendMessage(WeiboCookie, userUid, "哈喽，在互粉群中看到了你~可以互粉吗~")
		fmt.Println(userUid)
		i += 1
		if i > 80 {
			break
		}
	}
}
