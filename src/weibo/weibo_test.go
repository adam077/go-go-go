package weibo

import (
	"fmt"
	"testing"
	"time"
)

var cookie = "SINAGLOBAL=7188266670065.255.1499346921203; login_sid_t=8ab13d9a1abb502c3be871806d4c3f00; cross_origin_proto=SSL; _s_tentry=login.sina.com.cn; Apache=2154066483990.2334.1565109327603; ULV=1565109327616:39:1:1:2154066483990.2334.1565109327603:1563680713146; UOR=,weibo.com,login.sina.com.cn; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9Whaco_OwkDTB-SH898JGmvo5JpX5K2hUgL.FoqX1K2XehMReKM2dJLoI7y1dcia9GLzUBtt; ALF=1596647717; SSOLoginState=1565111718; SCF=AkQsjpu9Rz_J_k6nLpDzVEDRK1LUfVBstmGqJWe8pgYWvl8nEHEhFLbgYA-NJh_qREgzi3eP52jeVlLO7lI6zwQ.; SUB=_2A25wTcX2DeRhGeBK4lMV8CnEyjuIHXVTOrA-rDV8PUNbmtAKLXDTkW9NR2dre3AsDvQRuHhc2t2vCl8D3zdsbXLu; SUHB=00xp9GV9fpJd7x; un=13003638736; wvr=6; webim_unReadCount=%7B%22time%22%3A1565112780544%2C%22dm_pub_total%22%3A0%2C%22chat_group_client%22%3A0%2C%22allcountNum%22%3A1%2C%22msgbox%22%3A0%7D; WBStorage=edfd723f2928ec64|undefined"

func TestWeibo(t *testing.T) {
	uids, _ := GetUsers("互粉", cookie, 3)
	for _, uid := range uids {
		time.Sleep(10 * time.Second)
		err := Follow(uid, cookie)
		if err != nil {
			fmt.Println(err)
		}
	}
}
