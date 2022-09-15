package main

import (
	"strings"
)

//处理url形式，找到真实的url
func ProcessUrl(url string) string {
	var apiUrl string
	var uuid string
	listUrl := strings.Split(url, "/")
	uuid = listUrl[4]
	apiUrl = "http://perfeye.console.testplus.cn/api/show/task/" + uuid
	return apiUrl
}
