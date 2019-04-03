package main

import (
	"fmt"
	"time"
)

func main(){
	datach := make(chan int, 100)

	go func() {
		fmt.Println("2 秒后接受到信号开始发送")
		for i := 0; i < 100000; i++ {
			datach <- i
		}
		fmt.Println("发送端关闭数据通道")
		close(datach)

	}()

	time.Sleep(2 * time.Second)
	fmt.Println("开始通知发送信息")
	time.Sleep(1 * time.Second)
	fmt.Println("3秒后接受到数据通道数据 此时datach 在接收端已经关闭")
	for i := 0; i < 5; i++ {
		go func() {
			for {
				if i, ok := <-datach; ok {
					fmt.Println(i)

				} else {
					break
				}
			}
		}()

	}
	time.Sleep(5 * time.Second)
}
