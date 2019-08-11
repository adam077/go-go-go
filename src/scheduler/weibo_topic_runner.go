package scheduler

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"go-go-go/src/data"
	"go-go-go/src/ding-talk"
	"go-go-go/src/utils"
	"go-go-go/src/weibo"
)

type WeiboTopicRunner struct {
}

func (runner WeiboTopicRunner) Run() {
	userDatas := data.GetWeiboUserFollow("NobodyHu")
	for _, me := range userDatas {
		cookie := weibo.GetCookie(me.LoginName, me.Password)
		GetWeiBoTopic(cookie)
	}

}

func GetWeiBoTopic(cookie string) {
	var err error

	ts, today, hour, min := utils.GetNowTime()
	datas, errMap := weibo.GetWeiboTopic(cookie)
	if len(errMap) > 0 {
		errStr, _ := json.Marshal(errMap)
		ding_talk.SendDingMessage("c594857d21850991a7a15de920ab3c69b626fbb5d0f7bd3671dc0861bf13fab3", string(errStr))
	}
	if len(datas) == 0 {
		return
	}
	tb := data.WeiBoTopicMinuteReport{}.TableName()
	cols := []string{
		"create_time",
		"update_time",

		"log_date",
		"log_hour",
		"log_min",

		"cat",
		"content",
		"rank",

		"request_ts",
	}
	values := make([][]string, 0)
	keys := []string{"log_date", "log_hour", "log_min", "cat", "content"}
	updateCols := utils.GetUpdateTail(cols)
	for cat, v1 := range datas {
		for rank, topic := range v1 {
			temp := []string{
				"now()",
				"now()",

				utils.GetStr(&today),
				utils.GetInt(&hour),
				utils.GetInt(&min),

				utils.GetStr(&cat),
				utils.GetStr(&topic),
				utils.GetInt(&rank),

				fmt.Sprintf("to_timestamp(%d)", ts),
			}
			values = append(values, temp)
		}
	}
	if len(values) == 0 {
		return
	}
	sqlStr := utils.CreateBatchSql(tb, cols, values, keys, updateCols)

	db := data.GetDataDB("default")
	err = db.Exec(sqlStr).Error
	if err != nil {
		log.Error().Err(err).Str("sql", sqlStr)
	}
	return
}
