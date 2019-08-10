package weibo

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//全局变量
var CurCookies []*http.Cookie
var CurCookieJar *cookiejar.Jar //管理cookie

//初始化
func init() {
	CurCookies = nil
	//var err error;
	CurCookieJar, _ = cookiejar.New(nil)
}

//get url response html
func getUrlRespHtml(strUrl string, postDict map[string]string) string {
	fmt.Printf("in getUrlRespHtml, strUrl=%s\n", strUrl)
	fmt.Printf("postDict=%s\n", postDict)

	var respHtml string = ""

	httpClient := &http.Client{
		//Transport:nil,
		//CheckRedirect: nil,
		Jar: CurCookieJar,
	}

	var httpReq *http.Request
	//var newReqErr error
	if nil == postDict {
		fmt.Printf("is GET\n")
		//httpReq, newReqErr = http.NewRequest("GET", strUrl, nil)
		httpReq, _ = http.NewRequest("GET", strUrl, nil)
		// ...
		//httpReq.Header.Add("If-None-Match", `W/"wyzzy"`)
	} else {
		//【记录】go语言中实现http的POST且传递对应的post data
		//http://www.crifan.com/go_language_http_do_post_pass_post_data
		fmt.Printf("is POST\n")
		postValues := url.Values{}
		for postKey, PostValue := range postDict {
			postValues.Set(postKey, PostValue)
		}
		fmt.Printf("postValues=%s\n", postValues)
		postDataStr := postValues.Encode()
		fmt.Printf("postDataStr=%s\n", postDataStr)
		postDataBytes := []byte(postDataStr)
		fmt.Printf("postDataBytes=%s\n", postDataBytes)
		postBytesReader := bytes.NewReader(postDataBytes)
		//httpReq, newReqErr = http.NewRequest("POST", strUrl, postBytesReader)
		httpReq, _ = http.NewRequest("POST", strUrl, postBytesReader)
		//httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	httpResp, err := httpClient.Do(httpReq)
	// ...

	//httpResp, err := http.Get(strUrl)
	//gLogger.Info("http.Get done")
	if err != nil {
		fmt.Printf("http get strUrl=%s response error=%s\n", strUrl, err.Error())
	}
	//fmt.Printf("httpResp.Header=%s\n", httpResp.Header)
	//fmt.Printf("httpResp.Status=%s\n", httpResp.Status)

	defer httpResp.Body.Close()
	// gLogger.Info("defer httpResp.Body.Close done")

	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	//gLogger.Info("ioutil.ReadAll done")
	if errReadAll != nil {
		fmt.Printf("get response for strUrl=%s got error=%s\n", strUrl, errReadAll.Error())
	}
	//fmt.Printf("body=%s\n", body)
	//gCurCookies = httpResp.Cookies()
	//gCurCookieJar = httpClient.Jar;
	CurCookies = CurCookieJar.Cookies(httpReq.URL)
	//gLogger.Info("httpResp.Cookies done")
	//respHtml = "just for test log ok or not"
	respHtml = string(body)
	//gLogger.Info("httpResp body []byte to string done")
	return respHtml
}

//打印cookie
func printCurCookies() {
	var cookieNum int = len(CurCookies)
	fmt.Printf("cookieNum=%d\r\n", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = CurCookies[i]
		fmt.Printf("curCk.Raw=%s\r\n", curCk.Raw)
	}
}

//打印详细cookie
func printCurCookies2() {
	var cookieNum = len(CurCookies)
	fmt.Printf("cookieNum=%d\n", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = CurCookies[i]
		//fmt.Printf("curCk.Raw=%s", curCk.Raw)
		fmt.Printf("------ Cookie [%d]------\n", i)
		fmt.Printf("Name\t=%s\n", curCk.Name)
		fmt.Printf("Value\t=%s\n", curCk.Value)
		fmt.Printf("Path\t=%s\n", curCk.Path)
		fmt.Printf("Domain\t=%s\n", curCk.Domain)
		fmt.Printf("Expires\t=%s\n", curCk.Expires)
		fmt.Printf("RawExpires=%s\n", curCk.RawExpires)
		fmt.Printf("MaxAge\t=%d\n", curCk.MaxAge)
		fmt.Printf("Secure\t=%t\n", curCk.Secure)
		fmt.Printf("HttpOnly=%t\n", curCk.HttpOnly)
		fmt.Printf("Raw\t=%s\n", curCk.Raw)
		fmt.Printf("Unparsed=%s\n", curCk.Unparsed)
	}
}

//获取unix时间
func getMillisecond() int64 {
	MS := time.Now().UnixNano() / 1000
	return MS
}

//用户名base64加密
func encryptUname(uname string) string { // 获取username base64加密后的结果
	//println(base64.RawURLEncoding.EncodeToString([]byte(uname)))
	return base64.URLEncoding.EncodeToString([]byte(uname))
}

//密码加密
//把字符串转换bigint
func string2big(s string) *big.Int {
	ret := new(big.Int)
	ret.SetString(s, 16) // 将字符串转换成16进制
	return ret
}

func encryptPassword(pubkey string, servertime string, nonce string, password string) string {
	pub := rsa.PublicKey{
		N: string2big(pubkey),
		E: 65537, // 10001是十六进制数，65537是它的十进制表示
	}

	// servertime、nonce之间加\t，然后在\n ,和password拼接
	encryString := servertime + "\t" + nonce + "\n" + password

	// 拼接字符串加密
	encryResult, _ := rsa.EncryptPKCS1v15(rand.Reader, &pub, []byte(encryString))
	return hex.EncodeToString(encryResult)
}

func asdasd() {

	name := "13003638736" //新浪用户名
	password := "hu5845"  //登录密码

	//获取时间
	millisecond := getMillisecond()

	//username加密

	uname := encryptUname(name)

	//请求servertime、nonce、pubkey、rsakv

	getKeyUrl := "https://login.sina.com.cn/sso/prelogin.php?entry=account&callback=sinaSSOController.preloginCallBack&su=" + uname + "&rsakt=mod&client=ssologin.js(v1.4.15)&_=" + strconv.FormatInt(millisecond, 10)

	fmt.Println(getKeyUrl)

	body1 := getUrlRespHtml(getKeyUrl, nil)
	fmt.Printf("第一次返回结果：%s\n", body1)

	//拿参数：servertime、nonce、pubkey、rsakv
	//正则匹配拿json格式数据
	//这个测试一个字符串是否符合一个表达式。
	compile1, _ := regexp.Compile("{.*}")
	match1 := compile1.FindString(body1)
	fmt.Println(match1)
	//json str 转map
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(match1), &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)
	}
	servertime := dat["servertime"]
	//查看servertime数据类型，发现时float64
	//fmt.Print(reflect.TypeOf(servertime))
	servertime = strconv.FormatFloat(servertime.(float64), 'f', -1, 64)
	nonce := dat["nonce"]
	pubkey := dat["pubkey"]
	rsakv := dat["rsakv"]
	fmt.Print(nonce, pubkey, rsakv)

	//加密密码

	ep := encryptPassword(pubkey.(string), servertime.(string), nonce.(string), password)
	fmt.Printf("加密后的密码为：\n%s\n", ep)

	postDict := map[string]string{}
	postDict["entry"] = "account"
	postDict["gateway"] = "1"
	postDict["from"] = ""
	postDict["savestate"] = "30"
	postDict["qrcode_flag"] = "true"
	postDict["useticket"] = "0"
	postDict["pagerefer"] = ""
	postDict["vsnf"] = "1"
	postDict["su"] = uname
	postDict["service"] = "account"
	postDict["servertime"] = servertime.(string)
	postDict["nonce"] = nonce.(string)
	postDict["pwencode"] = "rsa2"
	postDict["rsakv"] = rsakv.(string)
	postDict["sp"] = ep
	postDict["sr"] = "1395*822"
	postDict["cdult"] = "3"
	postDict["domain"] = "sina.com.cn"
	postDict["prelt"] = "170"
	postDict["returntype"] = "TEXT"

	url2 := "https://login.sina.com.cn/sso/login.php?client=ssologin.js(v1.4.15)&_=" + strconv.FormatInt(getMillisecond(), 10)

	body2 := getUrlRespHtml(url2, postDict)
	fmt.Printf("第2次返回结果：%s\n", body2)

	printCurCookies2()

	url3 := "https://login.sina.com.cn/crossdomain2.php?action=login&r=https://login.sina.com.cn/"
	body3 := getUrlRespHtml(url3, nil)
	fmt.Printf("第3次返回结果：%s\n", body3)

	//  enc := mahonia.NewDecoder("gbk")
	//  output := enc.ConvertString(body3)
	//  fmt.Printf("%s\n",output[17:len(output)])
	//  fmt.Printf("%s\n",output)

	printCurCookies2()

	//正则匹配拿url
	//这个测试一个字符串是否符合一个表达式。
	compile2, _ := regexp.Compile("https:\\\\/\\\\/[^,]*")
	match2 := compile2.FindString(body3)
	fmt.Printf("匹配到的字符串为：%s\n", match2)
	replace := strings.Replace(match2, "\\", "", -1)
	println(replace)
	replace = strings.Replace(replace, "\"", "", -1)
	println(replace)

	body4 := getUrlRespHtml(replace, nil)
	fmt.Printf("第4次返回结果：%s\n", body4)
	printCurCookies2()

	//至此登录已经成功，cookie已拿到，下面为测试

	url4 := "http://weibo.com"

	body5 := getUrlRespHtml(url4, nil)
	fmt.Printf("第5次返回结果：%s\n", body5)
	printCurCookies2()

	url6 := "http://s.weibo.com/user/%25E9%2587%2591%25E8%259E%258D&Refer=weibo_user"
	body6 := getUrlRespHtml(url6, nil)
	fmt.Printf("第6次返回结果：%s\n", body6)
	printCurCookies2()

}

func GetCookie(name, password string) string {
	millisecond := getMillisecond()
	uname := encryptUname(name)
	getKeyUrl := "https://login.sina.com.cn/sso/prelogin.php?entry=account&callback=sinaSSOController.preloginCallBack&su=" + uname + "&rsakt=mod&client=ssologin.js(v1.4.15)&_=" + strconv.FormatInt(millisecond, 10)
	body1 := getUrlRespHtml(getKeyUrl, nil)
	compile1, _ := regexp.Compile("{.*}")
	match1 := compile1.FindString(body1)
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(match1), &dat); err == nil {
		fmt.Println("==============json str 转map=======================")
		fmt.Println(dat)
	}
	servertime := dat["servertime"]
	servertime = strconv.FormatFloat(servertime.(float64), 'f', -1, 64)
	nonce := dat["nonce"]
	pubkey := dat["pubkey"]
	rsakv := dat["rsakv"]
	ep := encryptPassword(pubkey.(string), servertime.(string), nonce.(string), password)
	postDict := map[string]string{}
	postDict["entry"] = "account"
	postDict["gateway"] = "1"
	postDict["from"] = ""
	postDict["savestate"] = "30"
	postDict["qrcode_flag"] = "true"
	postDict["useticket"] = "0"
	postDict["pagerefer"] = ""
	postDict["vsnf"] = "1"
	postDict["su"] = uname
	postDict["service"] = "account"
	postDict["servertime"] = servertime.(string)
	postDict["nonce"] = nonce.(string)
	postDict["pwencode"] = "rsa2"
	postDict["rsakv"] = rsakv.(string)
	postDict["sp"] = ep
	postDict["sr"] = "1395*822"
	postDict["cdult"] = "3"
	postDict["domain"] = "sina.com.cn"
	postDict["prelt"] = "170"
	postDict["returntype"] = "TEXT"
	url2 := "https://login.sina.com.cn/sso/login.php?client=ssologin.js(v1.4.15)&_=" + strconv.FormatInt(getMillisecond(), 10)
	getUrlRespHtml(url2, postDict)
	return printCurCookies3()

}

//打印详细cookie
func printCurCookies3() string {
	cookieList := make([]string, 0)
	for _, curCk := range CurCookies {
		cookieList = append(cookieList, fmt.Sprintf("%s=%s", curCk.Name, curCk.Value))
	}
	result := strings.Join(cookieList, "; ")
	return result
}
