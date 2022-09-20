package main

import (
	"strings"
)

//处理单个url形式，找到真实的url
func ProcessUrl(url string) string {
	var apiUrl string
	var uuid string
	listUrl := strings.Split(url, "/")
	uuid = listUrl[4]
	apiUrl = "http://perfeye.console.testplus.cn/api/show/task/" + uuid
	return apiUrl
}

//判断是否输入了多行url
func IsMoreUrls(url string) bool {
	if strings.Contains(url, ";") {
		return true
	} else {
		return false
	}
}

//处理多个url形式，找到真实的url
func ProcessUrls(url string) ([]string, []string) {
	var everyUrl []string
	var apiurls []string
	var originUrls []string
	var apiUrl string
	var uuid string
	var oriUrl string
	everyUrl = strings.Split(url, ";")
	for _, item := range everyUrl {
		if len(item) < 10 {
			break
		}
		listUrl := strings.Split(item, "/")
		uuid = listUrl[4]
		apiUrl = "http://perfeye.console.testplus.cn/api/show/task/" + uuid
		apiurls = append(apiurls, apiUrl)
		if strings.Contains(item, "\n") {
			oriUrl = strings.Trim(item, "\n")
			originUrls = append(originUrls, oriUrl)
		} else {
			originUrls = append(originUrls, item)
		}
	}
	return apiurls, originUrls
}
