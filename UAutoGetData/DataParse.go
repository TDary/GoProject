package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

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
	request.Header.Add("Cookie", "_ga=GA1.2.116075688.1657612386; Hm_lvt_4bfddcb32e5c5626aa3d10997c3dacd8=1657884787; app_key=5c316f30; Hm_lvt_eefc5ff12060e96822df38857e4cd9ed=1664279860,1664332706,1664449964,1665215179; project_key=mecha; mysession=MTY2NTMxMTk5M3xOd3dBTkRaV04waFlNbFZKVWpKRU4wVkpNa1UyVWs0eVZVbzNTalpVVlVKT04wUlBSVlZXUmtsS00wRTNRemMxTkVJeVNrSlJWa0U9fEdVdHRF6-3U_EWYs63zdverH8aaO92cXzchzbedpDnC; email=chenderui1%40thewesthill.net; Hm_lpvt_eefc5ff12060e96822df38857e4cd9ed=1665367975")
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
	result := ProcessData(string(body), report_url)
	if result != "" {
		return "false"
	}
	return "Success"
}

//周性能版本
func ProcessData2(resultData string, report_Url string) string {
	var DetailMap = make(map[string]interface{})
	err := json.Unmarshal([]byte(resultData), &DetailMap)
	if err != nil {
		return err.Error()
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
	return ""
}

//处理数据初始版本
func ProcessData(resultData string, report_Url string) string {
	var DetailMap = make(map[string]interface{})
	err := json.Unmarshal([]byte(resultData), &DetailMap)
	if err != nil {
		return err.Error()
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
	jank := DetailFPS.(map[string]interface{})["Jank(/10min)"]
	bigjank := DetailFPS.(map[string]interface{})["BigJank(/10min)"]
	ratio10 := DetailFPS.(map[string]interface{})["RatioFluctuate\u003e10(%)"]
	ratio30 := DetailFPS.(map[string]interface{})["RatioFluctuate\u003e30(%)"]
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
	resData := result{casename, avgfps, maxfps, minfps, tp90fps, jank, bigjank, ratio10, ratio30, avgapp, maxapp, InitMemory, AvgMemory, PeakMemory,
		avgGpuLoad, maxGpuLoad, avgGpuMemry, maxGpuMemry, avgDrawcall, maxDrawcall, avgVertex, maxVertex, avgPrimitive, maxPrimitive,
		avgSend, maxSend, avgRecv, maxRecv, avgRead, maxRead, avgWrite, maxWrite, report_Url}
	WriteData(resData)
	return ""
}
