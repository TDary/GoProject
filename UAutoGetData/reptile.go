package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var url_input string
	var isSuccess string
	var realUrl string
	file_name := "./" + "result" + ".csv"

	_, err := os.Stat(file_name)
	if err == nil {
		err = os.Remove(file_name)
		if err != nil {
			fmt.Print("result.csv文件正在打开中，请将其关闭后再执行本程序...")
			time.Sleep(time.Second * 10)
			return
		} else {
			//fmt.Print("清除旧文件完毕...")
			WriteHead(file_name)
		}
	} else {
		WriteHead(file_name)
	}
	//主逻辑
	for true {
		fmt.Println("Please Input an url：")
		fmt.Scan(&url_input)
		if strings.Contains(url_input, "http") {
			realUrl = ProcessUrl(url_input)
			isSuccess = getData(realUrl, url_input)
			if isSuccess == "Success" {
				fmt.Println("Get datas success!")
			}
		} else {
			fmt.Println("Error Input----")
		}
	}

}
