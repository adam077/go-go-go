package taobao

import (
	"fmt"
	"go-go-go/src/utils"
	"strconv"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	t1()
}

func t1() {
	t := strconv.Itoa(int(time.Now().UnixNano() / 1000000))
	a, err := utils.QueryGet("https://widget.1688.com/front/getJsonComponent.json?callback=jQuery17207196444384864313_"+t+"&props=%7B%22loginId%22%3A%22%E9%A9%B0%E7%8C%AB1003%22%2C%22pageNo%22%3A1%2C%22keyword%22%3A%22%22%2C%22memberId%22%3A%22%22%2C%22offerType%22%3A%22normal%22%2C%22supportDF%22%3Atrue%2C%22supportNY%22%3Afalse%2C%22supportCYS%22%3Afalse%7D&namespace=AlifeCsbcDxManagentOfferListActionsQueryOffers&widgetId=AlifeCsbcDxManagentOfferListActionsQueryOffers&methodName=execute",
		map[string]string{
			"Referer":        "https://guanjia.1688.com/page/offers.htm?spm=a2615.7691456.0.0.48257403cC7i6M&sellerLoginId=%E9%A9%B0%E7%8C%AB1003&offerType=normal",
			"Sec-Fetch-Mode": "no-cors",
			"User-Agent":     "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36",
		})
	if err != nil {
	}
	fmt.Println(string(a))

}
