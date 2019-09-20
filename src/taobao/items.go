package taobao

import (
	"encoding/json"
	"fmt"
	"go-go-go/src/single-cache"
	"go-go-go/src/utils"
	"net/url"
	"strconv"
	"time"
)

var count = 9

func GetItemList(floorId, word, taobaoToken string) Items {
	header := map[string]string{
		"cookie": taobaoToken,
	}
	b := map[string]interface{}{
		"pageNum":  0,
		"pageSize": 60,
		"refpid":   "mm_121000339_0_0",
	}
	if floorId != "" {
		b["floorId"] = floorId
	} else {
		b["floorId"] = "20392"
		b["variableMap"] = map[string]interface{}{
			"fn": "search",
			"q":  word,
			"_t": strconv.Itoa(int(time.Now().UnixNano() / 1000000)),
			//"coupon_amount_filter": "1~",
			//"toPage": 1,
			//"sort":   "sales:des",
		}
	}
	bStr, _ := json.Marshal(b)
	formData := url.Values{}
	formData.Add("_data_", string(bStr))
	var result Items
	a, _ := utils.QueryPostWithFormData("https://pub.alimama.com/openapi/json2/1/gateway.unionpub/optimus.material.json", nil, header, formData)
	var resultTemp Items
	fmt.Println(string(a))
	json.Unmarshal(a, &resultTemp)
	for _, item := range resultTemp.Model.Recommend.ResultList {
		if item.CouponAmount == "" {
			//log.Info().Str("item.ItemName", item.ItemName).Msg("")
			continue
		}
		_, ok := single_cache.Get(item.ItemId)
		if ok {
			continue
		} else {
			result.Model.Recommend.ResultList = append(result.Model.Recommend.ResultList, item)
			if len(result.Model.Recommend.ResultList) >= count {
				break
			}
		}
	}

	return result
}

func ShowSubstr(s string, l int) string {
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
				PromotionPrice    string `json:"promotionPrice"`    // 原价
				PriceAfterCoupon  string `json:"priceAfterCoupon"`  // 券后价
				CouponAmount      string `json:"couponAmount"`      // 券额
				CouponTotalCount  int    `json:"couponTotalCount"`  // 总数
				CouponRemainCount int    `json:"couponRemainCount"` // 剩余
			} `json:"resultList"`
		} `json:"recommend"`
	} `json:"model"`
}

func GetItemCode(itemId, taobaoToken string) Code {
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
		"cookie": taobaoToken,
	}
	a, _ := utils.QueryGet("https://pub.alimama.com/openapi/param2/1/gateway.unionpub/shareitem.json",
		params, header)
	var result Code
	json.Unmarshal(a, &result)
	return result

}

type Code struct {
	Data struct {
		ShortLinkInfo struct {
			CouponUrl string `json:"couponUrl"`
		} `json:"shortLinkInfo"`
		TaoTokenInfo struct {
			CouponUrl string `json:"couponUrl"`
		} `json:"taoTokenInfo"`
		UrlTransInfo struct {
			CouponUrl string `json:"couponUrl"`
		} `json:"urlTransInfo"`
	} `json:"data"`
}
