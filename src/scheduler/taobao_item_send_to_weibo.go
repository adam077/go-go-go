package scheduler

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go-go-go/src/data"
	"go-go-go/src/ding-talk"
	"go-go-go/src/single-cache"
	"go-go-go/src/taobao"
	"go-go-go/src/utils"
	"go-go-go/src/weibo"
	"math"
	"strconv"
	"strings"
	"time"
)

type ItemType struct {
	Name    string
	Topics  []string
	FloorId string
}

var ItemTypes = []ItemType{
	{"食品", []string{"零食"}, "19452"},
	//{"家居家装", []string{"家居"}, "19461"},
	{"母婴", []string{"母婴", "宝妈"}, "19457"},
	//{"数码家电", []string{"数码", "手机"}, "19458"},
	//{"全部", []string{}, "19450"},
}

type TaobaoToWeibo struct {
	One ItemType
}

var mins = 90

func (runner TaobaoToWeibo) Run() {
	defer utils.CommonRecover()
	if data.GetConfig(data.SchedulerTaobaoToWeibo) == "" {
		return
	}
	taobaoToken := data.GetConfig("taobao_cookie")
	weiboCookie := data.GetConfig("weibo_cookie")
	log.Info().Msg("start TaobaoToWeibo")
	for _, one := range []ItemType{runner.One} {
		doForWord(one, taobaoToken, weiboCookie)
	}
}

func doForWord(word ItemType, taobaoCookie, weiboCookie string) {
	var items taobao.Items
	items = taobao.GetItemList(word.FloorId, word.Name, taobaoCookie)
	weiboContent := "0|折扣|(原价) "
	for _, topic := range word.Topics {
		weiboContent += fmt.Sprintf("#%s# ", topic)
	}
	weiboContent += time.Now().Format(utils.TIME_DEFAULT_FORMAT) + "\n"
	weiboContent += "copy code to taobao or click url\n"
	i := 0
	picList := make([]string, 0)
	//toComment := make([]string, 0)
	log.Info().Int("count", len(items.Model.Recommend.ResultList)).Msg("getResultList")
	for _, item := range items.Model.Recommend.ResultList {
		weiboContentOneLine := ""
		after, _ := strconv.ParseFloat(item.PriceAfterCoupon, 64)
		coupon, _ := strconv.ParseFloat(item.CouponAmount, 64)
		all := after + coupon
		if all <= 0 {
			continue
		}
		sale := int(math.Floor(10 * coupon / all))
		if sale > 0 {
			i += 1
			weiboContentOneLine += fmt.Sprintf("%d|", i)

			//for i := 0; i < 10; i++ {
			//	if i < sale {
			//		weiboContentOneLine += "!"
			//	} else {
			//		weiboContentOneLine += " "
			//	}
			//}
			weiboContentOneLine += fmt.Sprintf("%d折", sale)

			weiboContentOneLine += fmt.Sprintf("|(%s)", item.PromotionPrice)
			item.ItemName = strings.Replace(item.ItemName, "<span class=H>", "", -1)
			item.ItemName = strings.Replace(item.ItemName, "</span>", "", -1)
			weiboContentOneLine += taobao.ShowSubstr(item.ItemName, 10)
			weiboContentOneLine += fmt.Sprintf(" ... %s\n", "http://o26166b137.wicp.vip/tb/"+item.ItemId)
			picList = append(picList, weibo.GetUrlPicId("http:"+item.Pic))
			log.Info().Str("itemId", item.ItemId).Msg("item code got")
			//code := taobao.GetItemCode(item.ItemId, taobaoCookie)
			//weiboContentOneLine += fmt.Sprintf("%d #%s# | %s \n", i, code.Data.TaoTokenInfo.CouponUrl, code.Data.ShortLinkInfo.CouponUrl)
			//toComment = append(toComment, fmt.Sprintf("%d copy#%s#OrClick %s", i, code.Data.TaoTokenInfo.CouponUrl, code.Data.ShortLinkInfo.CouponUrl))
			//toComment = append(toComment, fmt.Sprintf("%d #%s# | %s", i, code.Data.TaoTokenInfo.CouponUrl, code.Data.UrlTransInfo.CouponUrl))
			weiboContent += weiboContentOneLine
		}
	}

	//weiboContent += fmt.Sprintf("\n %s \n", " Copy       |       Click")
	//for x := range toComment {
	//	weiboContent += fmt.Sprintf("\n %s", toComment[x])
	//}
	fmt.Println(weiboContent)
	if len(picList) <= 0 {
		log.Error().Msg("no targets")
		ding_talk.SendDingMessage(data.GetConfig("ding"), "no targets")
		return
	}
	for _, item := range items.Model.Recommend.ResultList {
		single_cache.Set(item.ItemId, "", 60*mins)
	}
	mid, err := weibo.SendWeiBo(weiboCookie, weiboContent, picList)
	if err != nil {
		ding_talk.SendDingMessage(data.GetConfig("ding"), mid+"weibo send err")
		return
	}
	//time.Sleep(10 * time.Second)
	//log.Info().Str("mid", mid).Msg("send weibo")
	//for x := range toComment {
	//	time.Sleep(5 * time.Second)
	//	reverseInd := len(toComment) - 1 - x
	//	log.Info().Str("comment", toComment[reverseInd]).Msg("send weibo comment")
	//	weibo.SendWeiBoComment(weiboCookie, mid, toComment[reverseInd])
	//}
}
