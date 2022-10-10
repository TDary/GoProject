//go:generate goversioninfo -icon=resource/app.ico -manifest=resource/goversioninfo.exe.manifest
package main

import (
	"fmt"
	"os"
	"time"

	_ "embed"

	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

//go:embed resource/Reptitle.zip
var zip []byte

//go:embed resource/xcgui.dll
var dll []byte

//go:embed resource/Title.ico
var icon []byte

//窗口图标句柄设置
var hIcon = 0

func main() {
	var url_input string
	var isSuccess string
	var realUrl string
	var isMoreUrl bool
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

	err = xc.WriteDll(dll)
	if err != nil {
		panic(err)
	}
	// 炫彩_初始化, 参数填true是启用D2D硬件加速, 效果更好. 但xp系统不支持d2d, 这时候你就得填false来关闭d2d了
	ap := app.New(true)

	w := window.NewByLayoutZipMem(zip, "main.xml", "", 0, 0)

	// 从内存加载图片自适应大小
	hIcon = xc.XImage_LoadMemoryAdaptive(icon, 0, 0, 0, 0)
	// 因为下面写信息框还要用, 所以这里禁止图片自动销毁, 这样就可以复用了, 否则用过之后它会自动释放掉的
	xc.XImage_EnableAutoDestroy(hIcon, false)
	// 设置窗口图标
	w.SetIcon(hIcon)

	//标题样式设计
	title := widget.NewShapeTextByName("Title")
	title.SetSize(20, 20)

	//获取按钮以及输入框
	btn := widget.NewButtonByName("Btn")
	input := widget.NewEditByName("Info")
	input.SetDefaultText("请输入url...")
	//注册获取数据按钮被单击事件
	btn.Event_BnClick(func(pbHandled *bool) int {
		url_input = input.GetText_Temp()
		if len(url_input) < 10 {
			ap.Alert("提示", "输入url有误,请重新输入")
			return 0
		}
		isMoreUrl = IsMoreUrls(url_input)
		if isMoreUrl {
			realUrls, originUrls := ProcessUrls(url_input)
			for index, item := range realUrls {
				isSuccess = getData(item, originUrls[index])
				if isSuccess == "Success" {
					continue
				} else if isSuccess == "false" {
					ap.Alert("警告", "参数已过期，请联系开发人员更新@陈德睿")
				} else {
					ap.Alert("提示", "有输入错误的url,请重新输入")
				}
			}
			// 创建信息框, 本质是一个模态窗口
			hWindow := ap.Msg_Create("提示", "获取成功!", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
			// 设置窗口边框大小
			xc.XWnd_SetBorderSize(hWindow, 1, 34, 1, 1)
			// 设置窗口图标
			xc.XWnd_SetIcon(hWindow, hIcon)
			// 显示模态窗口
			xc.XModalWnd_DoModal(hWindow)
		} else {
			realUrl = ProcessUrl(url_input)
			isSuccess = getData(realUrl, url_input)
			if isSuccess == "Success" {
				// 创建信息框, 本质是一个模态窗口
				hWindow := ap.Msg_Create("提示", "获取成功!", xcc.MessageBox_Flag_Ok|xcc.MessageBox_Flag_Icon_Info, w.GetHWND(), xcc.Window_Style_Modal)
				// 设置窗口边框大小
				xc.XWnd_SetBorderSize(hWindow, 1, 34, 1, 1)
				// 设置窗口图标
				xc.XWnd_SetIcon(hWindow, hIcon)
				// 显示模态窗口
				xc.XModalWnd_DoModal(hWindow)
			} else if isSuccess == "false" {
				ap.Alert("警告", "参数已过期，请联系开发人员更新@陈德睿")
			} else {
				ap.Alert("提示", "输入url有误,请重新输入")
			}
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
