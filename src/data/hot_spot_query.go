package data

func GetZhihuData(date string, limit int) ([]string, []ZhiHuHotSpotMinuteReport) {
	var contents []ZhiHuHotSpotMinuteReport
	var contentlist []string
	var result []ZhiHuHotSpotMinuteReport
	dbConn := GetDataDB("default")
	sql1 := "select content from ( " +
		"SELECT content,count(1) as c,sum(rank) as r FROM zhihu_hot_spot_minute_report where log_date = ? group by content" +
		") as t1 where c > 10 order by r/c limit ?;"
	dbConn.Raw(sql1, date, limit).Find(&contents)
	for x := range contents {
		contentlist = append(contentlist, contents[x].Content)
	}
	sql2 := "SELECT distinct on (log_hour,content) log_hour,content,rank FROM zhihu_hot_spot_minute_report where log_date = ? and content in (?) order by log_hour,content,log_min desc;"
	//sql2 := "SELECT log_hour,content,rank FROM zhihu_hot_spot_minute_report where log_date = ? and log_min = 30 and content in (?);"
	dbConn.Raw(sql2, date, contentlist).Find(&result)
	return contentlist, result
}

func GetWeiboData(date string, limit int) ([]string, []WeiBoHotSpotMinuteReport) {
	var contents []WeiBoHotSpotMinuteReport
	var contentlist []string
	var result []WeiBoHotSpotMinuteReport
	dbConn := GetDataDB("default")
	sql1 := "select content from ( " +
		"SELECT content,count(1) as c,sum(rank) as r FROM weibo_hot_spot_minute_report where log_date = ? group by content" +
		") as t1 where c > 10 order by r/c limit ?;"
	dbConn.Raw(sql1, date, limit).Find(&contents)
	for x := range contents {
		contentlist = append(contentlist, contents[x].Content)
	}
	sql2 := "SELECT distinct on (log_hour,content) log_hour,content,rank FROM weibo_hot_spot_minute_report where log_date = ? and content in (?) order by log_hour,content,log_min desc;"
	//sql2 := "SELECT log_hour,content,rank FROM zhihu_hot_spot_minute_report where log_date = ? and log_min = 30 and content in (?);"
	dbConn.Raw(sql2, date, contentlist).Find(&result)
	return contentlist, result
}

func GetBaiduData(date string, limit int) ([]string, []BaiDuHotSpotMinuteReport) {
	var contents []BaiDuHotSpotMinuteReport
	var contentlist []string
	var result []BaiDuHotSpotMinuteReport
	dbConn := GetDataDB("default")
	sql1 := "select content from ( " +
		"SELECT content,count(1) as c,sum(rank) as r FROM baidu_hot_spot_minute_report where log_date = ? group by content" +
		") as t1 where c > 10 order by r/c limit ?;"
	dbConn.Raw(sql1, date, limit).Find(&contents)
	for x := range contents {
		contentlist = append(contentlist, contents[x].Content)
	}
	sql2 := "SELECT distinct on (log_hour,content) log_hour,content,rank FROM baidu_hot_spot_minute_report where log_date = ? and content in (?) order by log_hour,content,log_min desc;"
	//sql2 := "SELECT log_hour,content,rank FROM zhihu_hot_spot_minute_report where log_date = ? and log_min = 30 and content in (?);"
	dbConn.Raw(sql2, date, contentlist).Find(&result)
	return contentlist, result
}
