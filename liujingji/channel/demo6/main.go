package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	example1()
	example2()
}

// 示例1。
func example1() {
	// 准备好几个通道。
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	rand.Seed(time.Now().UnixNano())
	// 随机选择一个通道，并向它发送元素值。
	index := rand.Intn(3)
	fmt.Printf("index值是: %d\n", index)
	intChannels[index] <- index
	// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	select {
	case <-intChannels[0]:
		fmt.Println("第一个case分支被执行")
	case <-intChannels[1]:
		fmt.Println("第二个case分支被执行")
	case elem := <-intChannels[2]:
		fmt.Printf("第三个case分支被执行，元素是 %d.\n", elem)
	default:
		fmt.Println("没有case分支被执行")
	}
}

// 示例2。
func example2() {
	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	// time.AfterFunc(time.Second*5, func() {
	close(intChan)
	// })
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("case分支被关闭.")
			break
		}
		fmt.Println("case分支被执行.")
	default:
		fmt.Println("default分支执行")
	}
}
