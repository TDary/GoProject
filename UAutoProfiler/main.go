package main

import (
	GetTimes "UAutoProfiler/TimeTools"
	"fmt"
)

func main() {
	fmt.Print("Welcome to use UAutoServer")
	time := GetTimes.GetLogicTime()
	loger.Print(time)
}
