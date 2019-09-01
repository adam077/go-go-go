package weibo

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GetUrlPicId(urlStr string) string {
	base64Str, _ := getBase64(urlStr)
	return updloadPic(base64Str)
}

func updloadPic(base64 string) string {
	urlStr := "https://picupload.weibo.com/interface/pic_upload.php"
	params := map[string]string{
		"cb":          "https://weibo.com/aj/static/upimgback.html?_wv=5&callback=STK_ijax_" + strconv.Itoa(int(time.Now().UnixNano()/1000000)),
		"mime":        "image/jpeg",
		"data":        "base64",
		"url":         "weibo.com/u/6491407817",
		"markpos":     "1",
		"logo":        "1",
		"nick":        "@NobodyHu",
		"marks":       "0",
		"app":         "miniblog",
		"s":           "rdxt",
		"pri":         "null",
		"file_source": " 1",
	}
	header := map[string]string{
		"Cookie":                    cookie,
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
		"Accept-Encoding":           "gzip, deflate, br",
		"Accept-Language":           "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7",
		"Cache-Control":             "max-age=0",
		"Connection":                "keep-alive",
		"Content-Length":            "18177",
		"Content-Type":              "application/x-www-form-urlencoded",
		"Host":                      "picupload.weibo.com",
		"Origin":                    "https://weibo.com",
		"Referer":                   "https://weibo.com/u/6491407817/home?wvr=5",
		"Sec-Fetch-Mode":            "nested-navigate",
		"Sec-Fetch-Site":            "same-site",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36",
	}
	formData := url.Values{}
	formData.Add("b64_data", base64)
	a, err := QueryPostWithFormData(urlStr,
		params, header, formData)
	if err != nil {
	}
	l := strings.Split(a, "pid=")
	if len(l) > 0 {
		return l[len(l)-1]
	}
	return ""
}

func getBase64(imgUrl string) (string, error) {
	resp, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println("A error occurred!")
		return "", err
	}
	defer resp.Body.Close()
	a, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(a), nil

}

func QueryPostWithFormData(url string, params map[string]string, headers map[string]string, form url.Values) (string, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()
	for x := range headers {
		req.Header.Set(x, headers[x])
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := (&http.Client{
		Timeout: 15 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// https://studygolang.com/articles/2974
			if len(via) >= 0 {
				return errors.New("stopped after 10 redirects")
			}
			return nil
		},
	}).Do(req)
	return resp.Header.Get("Location"), err
}
