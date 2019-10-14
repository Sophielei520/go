package main

import (
	"fmt"
	"runtime"
)

func main() {

	//获取当前系统cpu的数量，NumCPU返回本地机器的逻辑CPU个数。
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum) //4
	//我这里设置num-1的cpu运行go程序
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
