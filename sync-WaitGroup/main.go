package main

import (
	"fmt"
	"sync"
)

func main() {
	/*WaitGroup 内部实现了一个计数器，用来记录未完成的操作个数，它提供了三个方法：
	  Add() 用来添加计数
	  Done() 用来在操作结束时调用，使计数减一
	  Wait() 用来等待所有的操作结束，即计数变为 0，该函数会在计数不为 0 时等待，在计数为 0 时立即返回*/
	var wg sync.WaitGroup
	wg.Add(2) //因为有两个任务执行，所以增加两个计数。
	go func() {
		fmt.Println("groutine1")
		wg.Done() //操作完成，减少一个计数。
	}()

	go func() {
		fmt.Println("groutine2")
		wg.Done() //操作完成，减少一个计数。
	}()
	wg.Wait() //等待，直到计数为0.
}
