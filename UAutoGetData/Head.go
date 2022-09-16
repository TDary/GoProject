package main

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
	CaseName         interface{}
	AvgFPS           interface{}
	MaxFPS           interface{}
	MinFPS           interface{}
	FPSTP90          interface{}
	Jank             interface{}
	BigJank          interface{}
	RatioFluctuate10 interface{}
	RatioFluctuate30 interface{}
	AvgApp           interface{}
	MaxApp           interface{}
	InitMemory       interface{}
	AvgMemory        interface{}
	PeakMemory       interface{}
	AvgGPULoad       interface{}
	MaxGPULoad       interface{}
	AvgGPUMemory     interface{}
	MaxGPUMemory     interface{}
	AvgDrawcall      interface{}
	PeakDrawcall     interface{}
	AvgVertex        interface{}
	PeakVertex       interface{}
	AvgPrimitive     interface{}
	PeakPrimitive    interface{}
	AvgSend          interface{}
	MaxSend          interface{}
	AvgRecv          interface{}
	MaxRecv          interface{}
	AvgRead          interface{}
	MaxRead          interface{}
	AvgWrite         interface{}
	MaxWrite         interface{}
	BasicData        string
}
