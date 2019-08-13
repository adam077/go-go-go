package weibo

import (
	"errors"
	"github.com/rs/zerolog/log"
	"go-go-go/src/utils"
	"strconv"
)

func GetUserStatus(cookie, uid string) (map[string]string, error) {
	pages := 100
	// 获取用户状态
	result := make(map[string]string)
	for i := 1; i <= pages; i++ {
		var urlStr = "https://weibo.com/p/100505" + uid + "/myfollow?Pl_Official_RelationMyfollow__95_page=" + strconv.Itoa(i)
		a, err := utils.QueryGet(urlStr, map[string]string{"Cookie": cookie})
		if err != nil {
			return result, err
		}
		uids := utils.FindBetween(a, "&uid=", "&sex")
		// 正则表达式中匹配一个反斜杠要用四个反斜杠  "<span class=\"S_txt1\">", "<\/span>")  但是这样又不好replace所以使用先替代再查找
		status := utils.FindBetween(a, "<span class=\\\"S_txt1\\\">", "<\\/span>")
		if len(uids) != len(status) {
			return result, errors.New("用户与状态的长度不一致")
		}
		for x := range uids {
			result[uids[x]] = status[x]
		}
		if len(uids) < 30 {
			break
		}
	}
	log.Info().Int("len", len(result)).Msg("")
	return result, nil
}

func GetFans() {

}
