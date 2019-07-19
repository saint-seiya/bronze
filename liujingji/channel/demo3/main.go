package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)
	// 发送方。
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("发送方，正在发送元素 %v...\n", i)
			ch1 <- i
			fmt.Printf("发送方，发送完毕元素 %v...\n", i)
			fmt.Printf("发送方，发送完毕ch1长度 %v...\n", len(ch1))
		}
		fmt.Println("发送放，关闭通道...")
		close(ch1)
		// 通道已经关闭，再继续操作会触发panic
		// ch1 <- 100

		// 关闭已经关闭的通道，会触发panic
		// close(ch1)
	}()

	// 接收方。
	for {
		fmt.Printf("接收方，ch1的长度: %v\n", len(ch1))
		elem, ok := <-ch1
		if !ok {
			fmt.Println("接收方，关闭通道")
			break
		}
		fmt.Printf("接收方，接收元素: %v\n", elem)
	}

	fmt.Println("结束")
}

/**
Q: 发送操作和接收操作在什么时候会引发panic?
A:
1、对于一个已经初始化，但并为关闭的通道来说，收发操作一定不会引发panic。但是通道一旦关闭，再对它进行发送操作，就会引发panic
2、如果试图关闭一个已经关闭了的通道，也会引发panic。注：接收操作是可以感知通道的关闭的，并能安全退出
当我们把接收表达式的结果同时赋值给两个变量时，第二个变量的类型一定是bool类型。它的值如果是false说明通道关闭，并且再没有元素值可取了。
注：如果通道关闭时，里面还有元素值未被取出，那么接收表达式的第一个结果，仍会是通道中的某一个元素值，而第二个结果值一定会是true。因此千万不要让接收方关闭通道，而应该让发送方关闭通道

3、通道需要手动关闭
*/
