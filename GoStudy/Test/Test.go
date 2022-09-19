//go:generate goversioninfo -icon=app.ico -manifest=goversioninfo.exe.manifest
package main

import "fmt"

func main() {
	var s []byte
	s = append(s, 'a')
	s = append(s, 't')
	fmt.Print(s[0])

}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //把sum发送到通道c
}
