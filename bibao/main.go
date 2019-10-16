package main

import "fmt"

//闭包就是一个函数与其相关的引用环境组合的一个整体。
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int { //返回一个匿名函数，但是这个匿名函数引用到函数外的变量n，
		//所以这个匿名函数就和n形成了一个整体，构成闭包。
		n = n + x
		return n
	}
}

func main() {
	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
}
