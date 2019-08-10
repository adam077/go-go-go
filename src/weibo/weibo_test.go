package weibo

import (
	"fmt"
	"testing"
	"time"
)

var cookie = "SINAGLOBAL=7188266670065.255.1499346921203; un=13003638736; wvr=6; UOR=,weibo.com,www.baidu.com; login_sid_t=6ebb73623df1b7903faf62fa49d6cb86; cross_origin_proto=SSL; _s_tentry=passport.weibo.com; Apache=5085805618345.12.1565282215469; ULV=1565282215481:40:2:2:5085805618345.12.1565282215469:1565109327616; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9Whaco_OwkDTB-SH898JGmvo5JpX5K2hUgL.FoqX1K2XehMReKM2dJLoI7y1dcia9GLzUBtt; ALF=1596818235; SSOLoginState=1565282236; SCF=AkQsjpu9Rz_J_k6nLpDzVEDRK1LUfVBstmGqJWe8pgYWrHKWS6If-rGo13XjH039baVoIqFoSyCCMEg4AS9ux2Y.; SUB=_2A25wSD_sDeRhGeBK4lMV8CnEyjuIHXVTPBYkrDV8PUNbmtAKLVH5kW9NR2drexkw8j12XSgMliOh155SiHXbvN5s; SUHB=0mxUzbDdWd5JHg; webim_unReadCount=%7B%22time%22%3A1565282272301%2C%22dm_pub_total%22%3A0%2C%22chat_group_client%22%3A0%2C%22allcountNum%22%3A0%2C%22msgbox%22%3A0%7D"

func TestWeibo(t *testing.T) {
	t2()
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
