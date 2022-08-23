package main

import (
	"fmt"
)

func main() {
	var typed string
	fmt.Print("欢迎使用自动下载最新版本机甲脚本：")
	fmt.Print("请输入要下载的包类型(release/debug)--其他非法输入可触发定时下载功能")
	typed = "tes"
	fmt.Print(typed)
}
