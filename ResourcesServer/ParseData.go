package main

//接受到开始解析的消息
func ParseBegin(data string) {
	loger.Print("收到开始传输信号----", data)
	PutData("BeginData", data)
}

//真正的处理数据逻辑
func ParseReal(data string) {
	loger.Print("收到真正数据----")
	PutData("RealData", data)
}

//停止解析的消息处理
func ParseStop(data string) {
	loger.Print("收到结束传输信号-----", data)
	PutData("StopData", data)
}

//结束解析的消息处理
func ParseEnd(data string) {
	loger.Print("收到结束传输信号----", data)
	PutData("EndData", data)
}

//放弃解析的消息处理
func ParseGiveUp(data string) {
	loger.Print("收到放弃数据信号----", data)
	PutData("GiveUpedData", data)
}
