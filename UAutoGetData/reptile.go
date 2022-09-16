package main

import (
	"fmt"
	"os"
	"time"

	_ "embed"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
)

//go:embed Reptitle.zip
var zip []byte

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

	// 炫彩_初始化, 参数填true是启用D2D硬件加速, 效果更好. 但xp系统不支持d2d, 这时候你就得填false来关闭d2d了
	ap := app.New(true)

	w := window.NewByLayoutZipMem(zip, "main.xml", "", 0, 0)

	//获取窗口布局文件中的按钮以及文本框
	btn := widget.NewButtonByName("Btn")
	input := widget.NewEditByName("Info")
	input.SetDefaultText("请输入一个url...")
	//注册按钮被单击事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		url_input = input.GetText_Temp()
		if len(url_input) < 10 {
			ap.Alert("提示", "输入url有误,请重新输入")
			return 0
		}
		realUrl = ProcessUrl(url_input)
		isSuccess = getData(realUrl, url_input)
		if isSuccess == "Success" {
			ap.Alert("提示", "获取成功(可继续获取)")
		} else {
			ap.Alert("提示", "输入url有误,请重新输入")
		}
		return 0
	})

	//调整布局
	w.AdjustLayout()
	// 显示窗口
	w.Show(true)
	// 运行消息循环, 程序会被阻塞在这里不退出, 当炫彩窗口数量为0时退出
	ap.Run()
	// 退出界面库释放资源
	ap.Exit()
}
