package test

import (
	"fmt"
	"github.com/chenjiandongx/go-echarts/charts"
	"github.com/gin-gonic/gin"
	"github.com/gomemcache/memcache"
	"github.com/rs/zerolog/log"
	"go-go-go/src/data"
	"go-go-go/src/scheduler"
	"go-go-go/src/single-cache"
	"go-go-go/src/utils"
	"net/http"
	"sort"
	"strconv"
	//"tianrang-client-center/src/common"
	"time"
)

type Person struct {
	Name          string `json:"name"`
	Score         int    `json:"score"`
	ScoreTHisTurn *int   `json:"scoreThisTurn"`
}

func GetPersons(c *gin.Context) {
	result := make([]Person, 0)
	asd := 15
	result = append(result, Person{
		Name:          "Adam Zhao",
		Score:         40,
		ScoreTHisTurn: &asd,
	})
	result = append(result, Person{
		Name:  "Adam Qian",
		Score: 30,
	})
	result = append(result, Person{
		Name:  "Adam Sun",
		Score: 20,
	})
	result = append(result, Person{
		Name:  "Adam Li",
		Score: 5,
	})
	utils.SuccessResp(c, "", result)
}

func GetEcharts(c *gin.Context) {
	n := []string{"hhh", "ddd", "asdf"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "yohaha"})
	bar.AddXAxis(n).
		AddYAxis("test1", []int{1, 2, 3})
	err := bar.XYReversal().Render(c.Writer)
	if err != nil {
		log.Warn().Msg(err.Error())
		utils.ErrorResp(c, http.StatusBadRequest, 400, err.Error())
	}
}

func GetZhihuEcharts(c *gin.Context) {
	params := struct {
		LogDate string `form:"logDate"`
		Limit   int    `form:"limit"`
	}{}
	if err := c.ShouldBindQuery(&params); err != nil {
		utils.ErrorResp(c, http.StatusBadRequest, 0, "")
		return
	}
	if params.LogDate == "" {
		params.LogDate = utils.GetTimeDateString(time.Now())
	}
	if params.Limit == 0 {
		params.Limit = 3
	}
	contents, datas := data.GetZhihuData(params.LogDate, params.Limit)
	hourList := make([]string, 0)
	for hour := 0; hour < 24; hour++ {
		hourList = append(hourList, strconv.Itoa(hour))
	}
	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "知乎", Bottom: "center"})
	contentMap := make(map[string]map[int]int)
	for _, ontData := range datas {
		if _, ok := contentMap[ontData.Content]; !ok {
			contentMap[ontData.Content] = make(map[int]int)
		}
		contentMap[ontData.Content][ontData.LogHour] = ontData.Rank
	}
	bar.AddXAxis(hourList)
	for _, content := range contents {
		hourMap := contentMap[content]
		rankData := make([]int, 0)
		for hour := 0; hour < 24; hour++ {
			rankData = append(rankData, GetRank(hourMap, hour))
		}
		bar.AddYAxis(content, rankData)
	}
	err := bar.Render(c.Writer)
	if err != nil {
		log.Warn().Msg(err.Error())
		utils.ErrorResp(c, http.StatusBadRequest, 400, err.Error())
	}
}

func GetWeiboEcharts(c *gin.Context) {
	params := struct {
		LogDate string `form:"logDate"`
		Limit   int    `form:"limit"`
	}{}
	if err := c.ShouldBindQuery(&params); err != nil {
		utils.ErrorResp(c, http.StatusBadRequest, 0, "")
		return
	}
	if params.LogDate == "" {
		params.LogDate = utils.GetTimeDateString(time.Now())
	}
	if params.Limit == 0 {
		params.Limit = 3
	}
	contents, datas := data.GetWeiboData(params.LogDate, params.Limit)
	hourList := make([]string, 0)
	for hour := 0; hour < 24; hour++ {
		hourList = append(hourList, strconv.Itoa(hour))
	}
	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "微博", Bottom: "center"})
	contentMap := make(map[string]map[int]int)
	for _, ontData := range datas {
		if _, ok := contentMap[ontData.Content]; !ok {
			contentMap[ontData.Content] = make(map[int]int)
		}
		contentMap[ontData.Content][ontData.LogHour] = ontData.Rank
	}
	bar.AddXAxis(hourList)
	for _, content := range contents {
		hourMap := contentMap[content]
		rankData := make([]int, 0)
		for hour := 0; hour < 24; hour++ {
			rankData = append(rankData, GetRank(hourMap, hour))
		}
		bar.AddYAxis(content, rankData)
	}
	err := bar.Render(c.Writer)
	if err != nil {
		log.Warn().Msg(err.Error())
		utils.ErrorResp(c, http.StatusBadRequest, 400, err.Error())
	}
}

func GetBaiduEcharts(c *gin.Context) {
	params := struct {
		LogDate string `form:"logDate"`
		Limit   int    `form:"limit"`
	}{}
	if err := c.ShouldBindQuery(&params); err != nil {
		utils.ErrorResp(c, http.StatusBadRequest, 0, "")
		return
	}
	if params.LogDate == "" {
		params.LogDate = utils.GetTimeDateString(time.Now())
	}
	if params.Limit == 0 {
		params.Limit = 3
	}
	contents, datas := data.GetBaiduData(params.LogDate, params.Limit)
	hourList := make([]string, 0)
	for hour := 0; hour < 24; hour++ {
		hourList = append(hourList, strconv.Itoa(hour))
	}
	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "百度", Bottom: "center"})
	contentMap := make(map[string]map[int]int)
	for _, ontData := range datas {
		if _, ok := contentMap[ontData.Content]; !ok {
			contentMap[ontData.Content] = make(map[int]int)
		}
		contentMap[ontData.Content][ontData.LogHour] = ontData.Rank
	}
	bar.AddXAxis(hourList)
	for _, content := range contents {
		hourMap := contentMap[content]
		rankData := make([]int, 0)
		for hour := 0; hour < 24; hour++ {
			rankData = append(rankData, GetRank(hourMap, hour))
		}
		bar.AddYAxis(content, rankData)
	}
	err := bar.Render(c.Writer)
	if err != nil {
		log.Warn().Msg(err.Error())
		utils.ErrorResp(c, http.StatusBadRequest, 400, err.Error())
	}
}

func GetRank(rankMap map[int]int, hour int) int {
	if rank, ok := rankMap[hour]; !ok {
		return 0
	} else {
		return 50 - rank
	}
}

func EatWhat(c *gin.Context) {
	params := struct {
		Eat string `form:"eat"`
		Set string `form:"set"`
	}{}
	c.ShouldBindQuery(&params)
	if params.Set != "" {
		switch params.Set {
		case scheduler.Choose:
			scheduler.Task1()
		case scheduler.Result:
			scheduler.Task2()
		case scheduler.ResetResult:
			scheduler.ResetTask()
		}
	} else {
		succ := scheduler.EnrichEatMap(c.ClientIP(), params.Eat)
		if succ {
			scheduler.Task1()
		}
	}
	SeeEatWhat(c)
}

func SeeEatWhat(c *gin.Context) {
	names := data.GetEatNames()
	if len(names) == 0 {
		return
	}

	result, ipList := scheduler.GetSortedEats(names)
	name := make([]string, 0)
	num := make([]int, 0)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count < result[j].Count
	})
	for x := range result {
		name = append(name, result[x].Name)
		num = append(num, result[x].Count)
	}

	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "吃啥呢"})
	bar.AddXAxis(name).
		AddYAxis("总票数", num)
	for _, ip := range ipList {
		temp := make([]int, 0)
		for _, one := range result {
			c := 0
			if one.IpCount != nil {
				c = one.IpCount[ip]
			}
			temp = append(temp, c)
		}
		bar.AddYAxis(ip, temp)
	}
	err := bar.XYReversal().Render(c.Writer)
	if err != nil {
		log.Warn().Msg(err.Error())
		utils.ErrorResp(c, http.StatusBadRequest, 400, err.Error())
	}
}

var mc *memcache.Client

func init() {
	// 需要自己开启
	//mc = memcache.New("127.0.0.1:11211")
}

func TestMemCacheSet(c *gin.Context) {
	var key = "foo"
	single_cache.Set(key, "123", 10)
}

func TestMemCacheExpire(c *gin.Context) {
	var key = "foo"
	it, ok := single_cache.Get(key)
	if ok {
		println(ok)
	} else {
		println(it)
	}
}

func TestMemCacheSet1(c *gin.Context) {

	var key = "foo"
	mc.Set(&memcache.Item{Key: key, Value: []byte("my value")})
	it, err := mc.Get(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	if string(it.Key) == key {
		fmt.Println("value is ", string(it.Value))
	} else {
		fmt.Println("Get failed")
	}
	mc.Set(&memcache.Item{Key: key, Value: []byte("my value 2")})
	it2, _ := mc.Get(key)
	if string(it2.Key) == key {
		fmt.Println("value is ", string(it2.Value))
	} else {
		fmt.Println("Get failed")
	}
	mc.Touch(key, 100)
}

func TestMemCacheExpire1(c *gin.Context) {
	var key = "foo"
	it, _ := mc.Get(key)
	if string(it.Key) == "foo" {
		fmt.Println("value is ", string(it.Value))
	} else {
		fmt.Println("Get failed")
	}
}
