package taobao

import (
	"encoding/json"
	"fmt"
	"go-go-go/src/utils"
	"math"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var count = 9

func getItemList(word string) {
	header := map[string]string{
		"cookie": TaobaoToken,
	}
	b := map[string]interface{}{
		"floorId":  "20392",
		"pageNum":  0,
		"pageSize": 60,
		"refpid":   "mm_121000339_0_0",
		"variableMap": map[string]interface{}{
			"fn": "search",
			"q":  word,
			"_t": strconv.Itoa(int(time.Now().UnixNano() / 1000000)),
			//"coupon_amount_filter": "1~",
			//"toPage": 1,
			//"sort":   "sales:des",
		},
	}
	bStr, _ := json.Marshal(b)
	formData := url.Values{}
	formData.Add("_data_", string(bStr))
	do := 0
	for {
		a, err := utils.QueryPostWithFormData("https://pub.alimama.com/openapi/json2/1/gateway.unionpub/optimus.material.json",
			nil, header, formData)
		if err != nil {
		}
		var result Items
		//fmt.Println(string(a))
		json.Unmarshal(a, &result)
		fmt.Println("|折扣力度|" + word)
		for _, item := range result.Model.Recommend.ResultList {
			if item.CouponAmount == "" {
				//log.Info().Str("item.ItemName", item.ItemName).Msg("")
				continue
			}
			after, _ := strconv.ParseFloat(item.PriceAfterCoupon, 64)
			coupon, _ := strconv.ParseFloat(item.CouponAmount, 64)
			all := after + coupon
			if all <= 0 {
				continue
			}
			sale := int(math.Floor(10 * coupon / all))
			if sale > 0 {
				fmt.Print("|")
				for i := 0; i < 10; i++ {
					if i < sale {
						fmt.Print("!")
					} else {
						fmt.Print(" ")
					}
				}
				do += 1
				fmt.Print("|")
				item.ItemName = strings.Replace(item.ItemName, "<span class=H>", "", -1)
				item.ItemName = strings.Replace(item.ItemName, "</span>", "", -1)
				fmt.Println(show_substr(item.ItemName, 30))
				//fmt.Print(" ")
				//fmt.Println(getItemCode(item.ItemId))
				if do >= count {
					break
				}
			}
		}
		if do >= count {
			break
		}
		do += 1
	}
}

func show_substr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}

		if sl+rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

type Items struct {
	Model struct {
		Recommend struct {
			ResultList []struct {
				ItemId            string `json:"itemId"`
				ItemName          string `json:"itemName"`
				Pic               string `json:"pic"`
				PriceAfterCoupon  string `json:"priceAfterCoupon"`  // 券后价
				CouponAmount      string `json:"couponAmount"`      // 券额
				CouponTotalCount  int    `json:"couponTotalCount"`  // 总数
				CouponRemainCount int    `json:"couponRemainCount"` // 剩余
			} `json:"resultList"`
		} `json:"recommend"`
	} `json:"model"`
}

func getItemCode(itemId string) string {
	params := map[string]string{
		"shareUserType":    "1",
		"unionBizCode":     "union_pub",
		"shareSceneCode":   "item_search",
		"materialId":       itemId,
		"tkClickSceneCode": "qtz_pub_search",
		"siteId":           "723000131",
		"adzoneId":         "109285850418",
		"materialType":     "1",
		"needQueryQtz":     "true",
	}
	header := map[string]string{
		"cookie": TaobaoToken,
	}
	a, err := utils.QueryGet("https://pub.alimama.com/openapi/param2/1/gateway.unionpub/shareitem.json",
		params, header)
	if err != nil {
	}
	var result Code
	json.Unmarshal(a, &result)
	return result.Data.ShortLinkInfo.CouponUrl

}

type Code struct {
	Data struct {
		ShortLinkInfo struct {
			CouponUrl string `json:"couponUrl"`
		} `json:"shortLinkInfo"`
		TaoTokenInfo struct {
			CouponUrl string `json:"couponUrl"`
		} `json:"taoTokenInfo"`
	} `json:"data"`
}
