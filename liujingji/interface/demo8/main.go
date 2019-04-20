package main

import "fmt"

type myInt int

func main() {
	var a myInt
	a = 2

	ifElse(nil)
	ifElse(1)
	ifElse("1")
	ifElse(true)
	ifElse(a)

	// typeSwtich(nil)
	// typeSwtich(1)
	// typeSwtich("1")
	// typeSwtich(true)
	// typeSwtich(a)
}
func ifElse(x interface{}) {
	if x == nil {
		fmt.Println("ifElse nil")
	} else if _, ok := x.(int); ok {
		fmt.Println("ifElse int")
	} else if _, ok := x.(uint); ok {
		fmt.Println("ifElse int")
	} else if _, ok := x.(bool); ok {
		fmt.Println("ifElse bool")
	} else if _, ok := x.(string); ok {
		fmt.Println("ifElse string")
	} else {
		fmt.Println("ifElse default")
	}
}
func typeSwtich(x interface{}) {
	switch x.(type) {
	case nil:
		fmt.Println("typeSwtich nil")
	case int, uint:
		fmt.Println("typeSwtich int, uint")
	case bool:
		fmt.Println("typeSwtich bool")
	case string:
		fmt.Println("typeSwtich string")
	default:
		fmt.Println("typeSwtich default")
	}
}
