package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

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
		"案例名", "AvgFPS", "MaxFPS", "MinFPS", "FPS TP90", "Jank(/10min)", "BigJank(/10min)", "RatioFluctuate>10%", "RatioFluctuate>30%", "AvgApp(%)", "MaxApp(%)", "InitMemory(MB)", "AvgMemory(MB)", "PeakMemory(MB)",
		"Avg(GPULoad)[%]", "Max(GPULoad)[%]", "Avg(GPUMemoryUsed)[MB]", "Peak(GPUMemoryUsed)[MB]", "Avg(Drawcall)",
		"Peak(Drawcall)", "Avg(Vertex)", "Peak(Vertex)", "Avg(Primitive)", "Peak(Primitive)", "AvgSend(KB/s)", "MaxSend(KB/s)", "AvgRecv(KB/s)",
		"MaxRecv(KB/s)", "AvgReadBytes(KB/s)", "MaxReadBytes(KB/s)", "AvgWrittenBytes(KB/s)", "MaxWrittenBytes(KB/s)", "基础数据"}
	w.Write(data)
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
		res.FPSTP90.(string), res.Jank.(string), res.BigJank.(string), res.RatioFluctuate10.(string), res.RatioFluctuate30.(string), res.AvgApp.(string), res.MaxApp.(string), res.InitMemory.(string), res.AvgMemory.(string),
		res.PeakMemory.(string), res.AvgGPULoad.(string), res.MaxGPULoad.(string), res.AvgGPUMemory.(string), res.MaxGPUMemory.(string),
		res.AvgDrawcall.(string), res.PeakDrawcall.(string), res.AvgVertex.(string), res.PeakVertex.(string), res.AvgPrimitive.(string),
		res.PeakPrimitive.(string), res.AvgSend.(string), res.MaxSend.(string), res.AvgRecv.(string), res.MaxRecv.(string), res.AvgRead.(string),
		res.MaxRead.(string), res.AvgWrite.(string), res.MaxWrite.(string), res.BasicData}
	w.Write(resData)
	w.Flush()
}
