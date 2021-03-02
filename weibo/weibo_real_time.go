package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func main() {
	pattern := regexp.MustCompile(``)
}

func getPageContent() string {
	var html string
	url := "https://weibo.com/a/hot/realtime"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return html
	}
	req.Header.Add("authority", "weibo.com")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"88\", \"Google Chrome\";v=\"88\", \";Not A Brand\";v=\"99\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")
	req.Header.Add("accept", "*/*")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://weibo.com/a/hot/realtime")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Add("cookie", "SUB=abc")
	req.Header.Add("if-modified-since", fmt.Sprintf("%s", time.Now().Format(time.RFC1123)))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return html
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return html
	}
	html = string(body)
	return html
}
