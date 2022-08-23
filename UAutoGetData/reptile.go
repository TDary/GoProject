package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type result2 struct {
	CaseName        interface{}
	AvgFPS          interface{}
	MinFPS          interface{}
	FPSTP90         interface{}
	SmoothPlay      interface{}
	PeakMemory      interface{}
	MaxGPUMemory    interface{}
	AvgApp          interface{}
	MaxApp          interface{}
	AvgGPULoad      interface{}
	MaxGPULoad      interface{}
	AvgDrawcall     interface{}
	PeakDrawcall    interface{}
	AvgTriAngles    interface{}
	MaxTriAngles    interface{}
	AverageUpload   interface{}
	MaxUpload       interface{}
	AverageDownLoad interface{}
	MaxDownLoad     interface{}
	BasicData       string
}

type result struct {
	CaseName      interface{}
	AvgFPS        interface{}
	MaxFPS        interface{}
	MinFPS        interface{}
	FPSTP90       interface{}
	AvgApp        interface{}
	MaxApp        interface{}
	InitMemory    interface{}
	AvgMemory     interface{}
	PeakMemory    interface{}
	AvgGPULoad    interface{}
	MaxGPULoad    interface{}
	AvgGPUMemory  interface{}
	MaxGPUMemory  interface{}
	AvgDrawcall   interface{}
	PeakDrawcall  interface{}
	AvgVertex     interface{}
	PeakVertex    interface{}
	AvgPrimitive  interface{}
	PeakPrimitive interface{}
	AvgSend       interface{}
	MaxSend       interface{}
	AvgRecv       interface{}
	MaxRecv       interface{}
	AvgRead       interface{}
	MaxRead       interface{}
	AvgWrite      interface{}
	MaxWrite      interface{}
	BasicData     string
}

var client = http.Client{
	Timeout: 10 * time.Second,
}

//获取数据  版本修改入口
func getData(url string, report_url string) string {
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err.Error()
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5060.66 Safari/537.36 Edg/103.0.1264.44")
	request.Header.Add("Cookie", "_ga=GA1.2.116075688.1657612386; app_key=e40280a0; Hm_lvt_4bfddcb32e5c5626aa3d10997c3dacd8=1657884787; Hm_lvt_eefc5ff12060e96822df38857e4cd9ed=1658981006,1659336247,1659600018,1660013626; project_key=jx3sim; mysession=MTY2MDAxNTcyN3xOd3dBTkV0RlJVMDBSRUpNTmt0Vk5VTllWMFZCUXpWSFdWcEJTbFUxUmtsWlRWSlRORGRFTlVsSk4xQkxXVGRYVmxOUlVrZzNOVkU9fCijNjwb45E-TbcLp3vaRhyrPkiBSO289Sy7R5DTS5Hf; email=chenderui1%40thewesthill.net; Hm_lpvt_eefc5ff12060e96822df38857e4cd9ed=1660015737")
	request.Header.Add("Referer", "http://perfeye.console.testplus.cn/case/list?appKey=mecha")
	response, err := client.Do(request)
	if err != nil {
		return err.Error()
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error()
	}
	ProcessData(string(body), report_url)
	return "Success"
}

//周性能版本
func ProcessData2(resultData string, report_Url string) {
	var DetailMap = make(map[string]interface{})
	err := json.Unmarshal([]byte(resultData), &DetailMap)
	if err != nil {
		panic(err)
	}
	casename := DetailMap["data"].(map[string]interface{})["BaseInfo"].(map[string]interface{})["CaseName"]
	DetailFPS := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelFPS"]
	DetailCpu := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelCPU"]
	DetailGpu := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelGPU"]
	DetailMemory := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelMemory"]
	DetailRender := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelRenderer"]
	DetailNetWork := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelNetwork"]
	avgfps := DetailFPS.(map[string]interface{})["AvgFPS"]
	minfps := DetailFPS.(map[string]interface{})["MinFPS"]
	tp90fps := DetailFPS.(map[string]interface{})["TP90"]
	smooth := DetailFPS.(map[string]interface{})["Smoothness(%)"]
	avgapp := DetailCpu.(map[string]interface{})["AvgApp(%)"]
	maxapp := DetailCpu.(map[string]interface{})["MaxApp(%)"]
	PeakMemory := DetailMemory.(map[string]interface{})["PeakMemory(MB)"]
	avgGpuLoad := DetailGpu.(map[string]interface{})["Avg(GPULoad)[%]"]
	maxGpuLoad := DetailGpu.(map[string]interface{})["Max(GPULoad)[%]"]
	maxGpuMemry := DetailGpu.(map[string]interface{})["Peak(GPUMemoryUsed)[MB]"]
	avgDrawcall := DetailRender.(map[string]interface{})["Avg(Drawcall)"]
	maxDrawcall := DetailRender.(map[string]interface{})["Peak(Drawcall)"]
	avgPrimitive := DetailRender.(map[string]interface{})["Avg(Primitive)"]
	maxPrimitive := DetailRender.(map[string]interface{})["Peak(Primitive)"]
	avgUpload := DetailNetWork.(map[string]interface{})["AvgSend(KB/s)"]
	maxUpload := DetailNetWork.(map[string]interface{})["MaxSend(KB/s)"]
	avgDownload := DetailNetWork.(map[string]interface{})["AvgRecv(KB/s)"]
	maxDownload := DetailNetWork.(map[string]interface{})["MaxRecv(KB/s)"]
	//获取到的数据，以json形式输出为csv
	resData := result2{casename, avgfps, minfps, tp90fps, smooth, PeakMemory, maxGpuMemry, avgapp, maxapp,
		avgGpuLoad, maxGpuLoad, avgDrawcall, maxDrawcall, avgPrimitive, maxPrimitive, avgUpload, maxUpload, avgDownload,
		maxDownload, report_Url}
	WriteData2(resData)
}

//处理数据初始版本
func ProcessData(resultData string, report_Url string) {
	var DetailMap = make(map[string]interface{})
	err := json.Unmarshal([]byte(resultData), &DetailMap)
	if err != nil {
		panic(err)
	}
	casename := DetailMap["data"].(map[string]interface{})["BaseInfo"].(map[string]interface{})["CaseName"]
	DetailFPS := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelFPS"]
	DetailCpu := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelCPU"]
	DetailGpu := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelGPU"]
	DetailMemory := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelMemory"]
	DetailRender := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelRenderer"]
	DetailNetWork := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelNetwork"]
	DetailIOBytyes := DetailMap["data"].(map[string]interface{})["LabelInfo"].(map[string]interface{})["All"].(map[string]interface{})["LabelIOBytes"]
	avgfps := DetailFPS.(map[string]interface{})["AvgFPS"]
	maxfps := DetailFPS.(map[string]interface{})["MaxFPS"]
	minfps := DetailFPS.(map[string]interface{})["MinFPS"]
	tp90fps := DetailFPS.(map[string]interface{})["TP90"]
	avgapp := DetailCpu.(map[string]interface{})["AvgApp(%)"]
	maxapp := DetailCpu.(map[string]interface{})["MaxApp(%)"]
	InitMemory := DetailMemory.(map[string]interface{})["InitMemory(MB)"]
	AvgMemory := DetailMemory.(map[string]interface{})["AvgMemory(MB)"]
	PeakMemory := DetailMemory.(map[string]interface{})["PeakMemory(MB)"]
	avgGpuLoad := DetailGpu.(map[string]interface{})["Avg(GPULoad)[%]"]
	maxGpuLoad := DetailGpu.(map[string]interface{})["Max(GPULoad)[%]"]
	avgGpuMemry := DetailGpu.(map[string]interface{})["Avg(GPUMemoryUsed)[MB]"]
	maxGpuMemry := DetailGpu.(map[string]interface{})["Peak(GPUMemoryUsed)[MB]"]
	avgDrawcall := DetailRender.(map[string]interface{})["Avg(Drawcall)"]
	maxDrawcall := DetailRender.(map[string]interface{})["Peak(Drawcall)"]
	avgVertex := DetailRender.(map[string]interface{})["Avg(Vertex)"]
	maxVertex := DetailRender.(map[string]interface{})["Peak(Vertex)"]
	avgPrimitive := DetailRender.(map[string]interface{})["Avg(Primitive)"]
	maxPrimitive := DetailRender.(map[string]interface{})["Peak(Primitive)"]
	avgSend := DetailNetWork.(map[string]interface{})["AvgSend(KB/s)"]
	maxSend := DetailNetWork.(map[string]interface{})["MaxSend(KB/s)"]
	avgRecv := DetailNetWork.(map[string]interface{})["AvgRecv(KB/s)"]
	maxRecv := DetailNetWork.(map[string]interface{})["MaxRecv(KB/s)"]
	avgRead := DetailIOBytyes.(map[string]interface{})["AvgReadBytes(KB/s)"]
	maxRead := DetailIOBytyes.(map[string]interface{})["MaxReadBytes(KB/s)"]
	avgWrite := DetailIOBytyes.(map[string]interface{})["AvgWrittenBytes(KB/s)"]
	maxWrite := DetailIOBytyes.(map[string]interface{})["MaxWrittenBytes(KB/s)"]
	//获取到的数据，以json形式输出为csv
	resData := result{casename, avgfps, maxfps, minfps, tp90fps, avgapp, maxapp, InitMemory, AvgMemory, PeakMemory,
		avgGpuLoad, maxGpuLoad, avgGpuMemry, maxGpuMemry, avgDrawcall, maxDrawcall, avgVertex, maxVertex, avgPrimitive, maxPrimitive,
		avgSend, maxSend, avgRecv, maxRecv, avgRead, maxRead, avgWrite, maxWrite, report_Url}
	WriteData(resData)
}

//写入数据
func WriteData2(res result2) {
	file_name := "./" + "result" + ".csv"
	f, err := os.OpenFile(file_name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") //写入UTF-8
	w := csv.NewWriter(f)         //创建一个新的写入文件流
	AvgTris := res.AvgTriAngles.(string)
	MaxTris := res.MaxTriAngles.(string)

	AVGtrian, err := strconv.ParseFloat(AvgTris, 32)
	if err != nil {
		panic(err)
	}
	MAXtrian, err := strconv.ParseFloat(MaxTris, 32)
	if err != nil {
		panic(err)
	}
	AVGtrian = AVGtrian / 1000
	MAXtrian = MAXtrian / 1000

	resData := []string{res.CaseName.(string), res.AvgFPS.(string), res.MinFPS.(string), res.FPSTP90.(string),
		res.SmoothPlay.(string), res.PeakMemory.(string), res.MaxGPUMemory.(string), res.AvgApp.(string), res.MaxApp.(string), res.AvgGPULoad.(string),
		res.MaxGPULoad.(string), res.AvgDrawcall.(string), res.PeakDrawcall.(string), strconv.FormatFloat(AVGtrian, 'E', -1, 32), strconv.FormatFloat(MAXtrian, 'E', -1, 32), res.AverageUpload.(string),
		res.MaxUpload.(string), res.AverageDownLoad.(string), res.MaxDownLoad.(string), res.BasicData}
	w.Write(resData)
	w.Flush()
}

//写入数据
func WriteData(res result) {
	file_name := "./" + "result" + ".csv"
	f, err := os.OpenFile(file_name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") //写入UTF-8
	w := csv.NewWriter(f)         //创建一个新的写入文件流

	resData := []string{res.CaseName.(string), res.AvgFPS.(string), res.MaxFPS.(string), res.MinFPS.(string),
		res.FPSTP90.(string), res.AvgApp.(string), res.MaxApp.(string), res.InitMemory.(string), res.AvgMemory.(string),
		res.PeakMemory.(string), res.AvgGPULoad.(string), res.MaxGPULoad.(string), res.AvgGPUMemory.(string), res.MaxGPUMemory.(string),
		res.AvgDrawcall.(string), res.PeakDrawcall.(string), res.AvgVertex.(string), res.PeakVertex.(string), res.AvgPrimitive.(string),
		res.PeakPrimitive.(string), res.AvgSend.(string), res.MaxSend.(string), res.AvgRecv.(string), res.MaxRecv.(string), res.AvgRead.(string),
		res.MaxRead.(string), res.AvgWrite.(string), res.MaxWrite.(string), res.BasicData}
	w.Write(resData)
	w.Flush()
}

//周性能版本
func WriteHead2(file_name string) {
	f, err := os.Create(file_name) //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") //写入UTF-8
	w := csv.NewWriter(f)         //创建一个新的写入文件流
	data := []string{
		"案例名", "平均FPS", "最低FPS", "90%FPS大于(>=40FPS)", "流畅度", "私有提交内存峰值<=8GB", "显存峰值", "平均CPU占用", "最大CPU占用", "平均GPU占用",
		"最大GPU占用", "平均Draw Call", "最大Draw Call", "平均三角面数(千)", "最大三角面数(千)",
		"平均上传(KB/s)", "上传峰值(KB/s)", "平均下载(KB/s)", "下载峰值(KB/s)", "性能报告地址"}
	w.Write(data)
	w.Flush()
}

//初始版本
func WriteHead(file_name string) {
	f, err := os.Create(file_name) //创建文件
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") //写入UTF-8
	w := csv.NewWriter(f)         //创建一个新的写入文件流
	data := []string{
		"案例名", "AvgFPS", "MaxFPS", "MinFPS", "FPS TP90", "AvgApp(%)", "MaxApp(%)", "InitMemory(MB)", "AvgMemory(MB)", "PeakMemory(MB)",
		"Avg(GPULoad)[%]", "Max(GPULoad)[%]", "Avg(GPUMemoryUsed)[MB]", "Peak(GPUMemoryUsed)[MB]", "Avg(Drawcall)",
		"Peak(Drawcall)", "Avg(Vertex)", "Peak(Vertex)", "Avg(Primitive)", "Peak(Primitive)", "AvgSend(KB/s)", "MaxSend(KB/s)", "AvgRecv(KB/s)",
		"MaxRecv(KB/s)", "AvgReadBytes(KB/s)", "MaxReadBytes(KB/s)", "AvgWrittenBytes(KB/s)", "MaxWrittenBytes(KB/s)", "基础数据"}
	w.Write(data)
	w.Flush()
}

//处理url形式，找到真实的url
func ProcessUrl(url string) string {
	var apiUrl string
	var uuid string
	listUrl := strings.Split(url, "/")
	uuid = listUrl[4]
	apiUrl = "http://perfeye.console.testplus.cn/api/show/task/" + uuid
	return apiUrl
}

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
