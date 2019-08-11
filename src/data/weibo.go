package data

import (
	"github.com/rs/zerolog/log"
	"go-go-go/src/api_query"
	"go-go-go/src/utils"
)

type WeiboFollow struct {
	BaseModelIncrementID
	WeiboUserId string `gorm:"column:weibo_user_id;type:text"`
	UID         string `gorm:"column:uid;type:text"`
}

func (WeiboFollow) TableName() string {
	return "weibo_follow"
}

func (*WeiboFollow) UniqueIndexes() map[string][]string {
	return map[string][]string{"unique_index_weibo_follow": {"weibo_user_id", "uid"}}
}

func UpdateWeiboFollower(targets []WeiboFollow) error {
	datas := api_query.GetWeibo()

	if len(datas) == 0 {
		return nil
	}
	tb := WeiboFollow{}.TableName()
	cols := []string{
		"create_time",
		"update_time",

		"weibo_user_id",
		"uid",
	}
	values := make([][]string, 0)
	keys := []string{"weibo_user_id", "uid"}
	updateCols := utils.GetUpdateTail(cols)
	for _, v1 := range targets {
		temp := []string{
			"now()",
			"now()",

			utils.GetStr(&v1.WeiboUserId),
			utils.GetStr(&v1.UID),
		}
		values = append(values, temp)
	}
	sqlStr := utils.CreateBatchSql(tb, cols, values, keys, updateCols)
	db := GetDataDB("default")
	err := db.Exec(sqlStr).Error
	if err != nil {
		log.Error().Err(err).Str("sql", sqlStr)
	}
	return err
}

type WeiboUser struct {
	BaseModelUUID
	Name      string `gorm:"column:name;type:text"`
	LoginName string `gorm:"column:login_name;type:text"`
	Password  string `gorm:"column:password;type:text"`
	Uid       string `gorm:"column:uid;type:text"`

	Status int `gorm:"column:status;type:int4"`

	WeiboFollows []*WeiboFollow `gorm:"foreignkey:WeiboUserId;association_foreignkey:ID"`
}

func (WeiboUser) TableName() string {
	return "weibo_user"
}

func GetWeiboUserFollow(name string) []*WeiboUser {
	var err error
	dataConn := GetDataDB("default")
	var temp []*WeiboUser
	var pre = dataConn.Preload("WeiboFollows").Where("status = ?", 1)
	if name == "" {
		err = pre.Find(&temp).Error
	} else {
		err = pre.Find(&temp, "name = ?", name).Error
	}
	if err != nil {
		panic(err)
	}
	return temp

}

type WeiboChat struct {
	BaseModelIncrementID
	WeiboUserId string `gorm:"column:weibo_user_id;type:text"`
	UID         string `gorm:"column:uid;type:text"`
}

func (WeiboChat) TableName() string {
	return "weibo_chat"
}

func (*WeiboChat) UniqueIndexes() map[string][]string {
	return map[string][]string{"unique_index_weibo_chat": {"weibo_user_id", "uid"}}
}

func UpdateWeiboChat(targets []WeiboChat) error {
	datas := api_query.GetWeibo()

	if len(datas) == 0 {
		return nil
	}
	tb := WeiboChat{}.TableName()
	cols := []string{
		"create_time",
		"update_time",

		"weibo_user_id",
		"uid",
	}
	values := make([][]string, 0)
	keys := []string{"weibo_user_id", "uid"}
	updateCols := utils.GetUpdateTail(cols)
	for _, v1 := range targets {
		temp := []string{
			"now()",
			"now()",

			utils.GetStr(&v1.WeiboUserId),
			utils.GetStr(&v1.UID),
		}
		values = append(values, temp)
	}
	sqlStr := utils.CreateBatchSql(tb, cols, values, keys, updateCols)
	db := GetDataDB("default")
	err := db.Exec(sqlStr).Error
	if err != nil {
		log.Error().Err(err).Str("sql", sqlStr)
	}
	return err
}

func GetWeiboUserChat(weiboUserId string) []*WeiboChat {
	dataConn := GetDataDB("default")
	var temp []*WeiboChat
	err := dataConn.Where("weibo_user_id = ?", weiboUserId).Find(&temp).Error
	if err != nil {
		panic(err)
	}
	return temp

}
