package data

import (
	"testing"
)

func TestMigrateTable(t *testing.T) {
	//t1()
	//t2()
	//t3()
	//t4()
	//t5()
}

func t1() {
	db := GetDataDB("default")

	db.AutoMigrate(&WeiBoHotSpotMinuteReport{})
	db.Exec("CREATE UNIQUE INDEX weibo_hot_spot_minute_report_unique_idx ON weibo_hot_spot_minute_report " +
		"USING btree (log_date, log_hour, log_min, content);")
	db.AutoMigrate(&BaiDuHotSpotMinuteReport{})
	db.Exec("CREATE UNIQUE INDEX baidu_hot_spot_minute_report_unique_idx ON baidu_hot_spot_minute_report " +
		"USING btree (log_date, log_hour, log_min, content);")
	db.AutoMigrate(&ZhiHuHotSpotMinuteReport{})
	db.Exec("CREATE UNIQUE INDEX zhihu_hot_spot_minute_report_unique_idx ON zhihu_hot_spot_minute_report " +
		"USING btree (log_date, log_hour, log_min, content);")
}

func t2() {
	db := GetDataDB("default")
	db.AutoMigrate(&EatWhatTable{})
	db.AutoMigrate(&Ding{})
}

func t3() {
	db := GetDataDB("config")
	db.AutoMigrate(&ConfigTable{})
}

func t4() {
	db := GetDataDB("default")
	MigrateTable(db, &WeiboUser{})
	MigrateTable(db, &WeiboFollow{})
	MigrateTable(db, &WeiboChat{})

	//one := &WeiboUser{}
	//one.ID = "e2e32ebd-69cb-4855-b4f4-97d5ec2e71fa"
	//one.Name = "NobodyHu"
	//db.Create(one)
}

func t5() {
	db := GetDataDB("default")

	db.AutoMigrate(&WeiBoTopicMinuteReport{})
	db.Exec("CREATE UNIQUE INDEX weibo_topic_minute_report_unique_idx ON weibo_topic_minute_report " +
		"USING btree (log_date, log_hour, log_min, cat, content);")

}
