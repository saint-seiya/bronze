package main

import (
	"fmt"
	"math/rand"
)

/**
单向通道：
1、<-用在通道类型的字面量中，就表示通道的方向，称为单向通道
2、分为发送通道和接收通道两种
3、单向通道是无用的
*/
func main() {
	// 示例1。
	// 只能发不能收的通道叫发送通道。
	var uselessChan = make(chan<- int, 1)
	// 只能收不能发的通道叫接收通道。
	var anotherUselessChan = make(<-chan int, 1)
	// 这里打印的是可以分别代表两个通道的指针的16进制表示。
	fmt.Printf("无用通道: %v, %v\n",
		uselessChan, anotherUselessChan)

	// 示例2。
	intChan1 := make(chan int, 3)
	SendInt(intChan1)

	// 示例4。
	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}

	// 示例5。
	_ = GetIntChan(getIntChan)
}

// 示例2。
func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

// 示例3。
type Notifier interface {
	SendInt(ch chan<- int)
}

// 示例4。
func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

// 示例5。
type GetIntChan func() <-chan int
