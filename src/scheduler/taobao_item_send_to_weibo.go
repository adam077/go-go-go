package scheduler

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go-go-go/src/data"
	"go-go-go/src/ding-talk"
	"go-go-go/src/taobao"
	"go-go-go/src/utils"
	"go-go-go/src/weibo"
	"math"
	"strconv"
	"strings"
)

type ItemType struct {
	Name    string
	Topics  []string
	FloorId string
}

type TaobaoToWeibo struct {
}

func (runner TaobaoToWeibo) Run() {
	defer utils.CommonRecover()
	if data.GetConfig(data.SchedulerTaobaoToWeibo) == "" {
		return
	}
	taobaoToken := data.GetConfig("taobao_cookie")
	weiboCookie := data.GetConfig("weibo_cookie")
	log.Info().Msg("start TaobaoToWeibo")
	for _, one := range []ItemType{
		{"食品", []string{"零食"}, "19452"},
		{"家居家装", []string{"家居"}, "19461"},
		{"母婴", []string{"母婴", "宝妈"}, "19457"},
		{"数码家电", []string{"数码", "手机"}, "19458"},
		//{"全部", []string{}, "19450"},
	} {
		doForWord(one, taobaoToken, weiboCookie)
	}
}

func doForWord(word ItemType, taobaoCookie, weiboCookie string) {
	var items taobao.Items
	items = taobao.GetItemList(word.FloorId, word.Name, taobaoCookie)
	weiboContent := "0|打折比例| "
	for _, topic := range word.Topics {
		weiboContent += fmt.Sprintf("#%s# ", topic)
	}
	weiboContent += "\n"
	i := 0
	picList := make([]string, 0)
	toComment := make([]string, 0)
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
			for i := 0; i < 10; i++ {
				if i < sale {
					weiboContentOneLine += "!"
				} else {
					weiboContentOneLine += " "
				}
			}
			log.Info().Str("itemId", item.ItemId).Msg("item execute")
			weiboContentOneLine += fmt.Sprintf("|原价(%s)", item.PromotionPrice)
			item.ItemName = strings.Replace(item.ItemName, "<span class=H>", "", -1)
			item.ItemName = strings.Replace(item.ItemName, "</span>", "", -1)
			weiboContentOneLine += taobao.ShowSubstr(item.ItemName, 18)
			weiboContentOneLine += "...\n"
			picList = append(picList, weibo.GetUrlPicId("http:"+item.Pic))
			log.Info().Str("itemId", item.ItemId).Msg("item code got")
			code := taobao.GetItemCode(item.ItemId, taobaoCookie)
			toComment = append(toComment, fmt.Sprintf("%d 复制#%s#至App或点击 %s",
				i, code.Data.TaoTokenInfo.CouponUrl, code.Data.ShortLinkInfo.CouponUrl))
			weiboContent += weiboContentOneLine
		}
	}

	for x := range toComment {
		weiboContent += fmt.Sprintf("\n %s", toComment[x])
	}
	fmt.Println(weiboContent)
	if len(picList) <= 0 {
		log.Error().Msg("no targets")
		ding_talk.SendDingMessage(data.GetConfig("ding"), "no targets")
		return
	}
	mid, err := weibo.SendWeiBo(weiboCookie, weiboContent, picList)
	if err != nil {
		ding_talk.SendDingMessage(data.GetConfig("ding"), mid+"weibo err")
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
