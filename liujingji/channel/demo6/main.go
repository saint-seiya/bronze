package main

import "fmt"

var channels = [3]chan int{
	nil,
	make(chan int),
	nil,
}

/**
select语句的分支选择规则
1、对于每个case表达式，都至少包含一个代表发送操作的发送表达式或一个代表接收操作的接收表达式，同时也可能会包含其他的表达式。比如，如果case表达式是包含了接收表达式的短变量声明时，那么在赋值符合左边的就可以是一个或两个表达式，不过此处的表达式的结果必须是可以被赋值的。当这样的case表达式被求值时，它包含的多个表达式总会以从左到右的顺序被求值
2、select语句包含的候选分支中的case表达式都会在该语句执行开始时先被求值，并且求值的顺序是从上到下的
3、对于每个case表达式，如果其中的发送表达式或接收表达式在被求值时，相应的操作正处于阻塞状态，那么该case表达式的求值就是不成功的，即不满足当前的case分支
4、只有当select语句中的所有case表达式都被求值完毕后，它才会开始选择候选分支。如果有对应的候选分支满足则执行候选分支，如果没有候选分支满足条件，则执行默认分支。如果没有默认分支，则select语句进入阻塞状态，直到至少有一个候选分支满足选择条件为止。
5、如果select语句发现同时满足多个case分支，那么它会用一种伪随机算法在分支中选择一个执行。
6、一条select中只能有一个默认分支
7、select语句的执行，包括case表达式求值和分支选择，都是独立的。不过，至于它的执行是否是并发安全的，就要看其中的case表达式已经分钟，是否包含并发不安全的代码了
*/
var numbers = []int{1, 2, 3}

func main() {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("第一个case被执行")
	case getChan(1) <- getNumber(1):
		fmt.Println("第二个case被执行")
	case getChan(2) <- getNumber(2):
		fmt.Println("第三个case被执行")
	default:
		fmt.Println("默认分支被执行")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}
