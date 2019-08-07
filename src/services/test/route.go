package test

import "github.com/gin-gonic/gin"

var MonitorRoutes = map[string]map[string]gin.HandlersChain{
	"get_persons": {
		"GET": gin.HandlersChain{GetPersons},
	},
	"": {
		"GET": gin.HandlersChain{Haha1},
	},
	"hah": {
		"GET": gin.HandlersChain{Haha2},
	},
	"echarts/zhihu": {
		"GET": gin.HandlersChain{GetZhihuEcharts},
	},
	"echarts/weibo": {
		"GET": gin.HandlersChain{GetWeiboEcharts},
	},
	"echarts/baidu": {
		"GET": gin.HandlersChain{GetBaiduEcharts},
	},
	"set_eat_what": {
		"GET": gin.HandlersChain{EatWhat},
	},
	"eat_what": {
		"GET": gin.HandlersChain{SeeEatWhat},
	},
	"test_mc_set": {
		"GET": gin.HandlersChain{TestMemCacheSet},
	},
	"test_mc_expire": {
		"GET": gin.HandlersChain{TestMemCacheExpire},
	},
}
