package main

import "fmt"

type Tester interface {
	getName() string
}
type Tester2 interface {
	printName()
}

type Person struct {
	name string
}

func (p Person) getName() string {
	return p.name
}
func (p Person) printName() {
	fmt.Println(p.name)
}

type Person1 struct {
	name string
}

func (p Person1) getName() string {
	return p.name
}
func (p Person1) printName() {
	fmt.Println(p.name)
}

func main() {
	var t Tester
	var m Tester
	t = Person{"xiaohua"}
	m = Person1{"360"}
	fmt.Printf("t.(Person)前, t的动态类型%T", t)
	fmt.Println()
	fmt.Printf("t.(Person)前, t的动态值%v", t)
	fmt.Println()
	if f, ok1 := t.(Person); ok1 {
		fmt.Printf("t.(Person)后, t的动态类型%T", t)
		fmt.Println()
		fmt.Printf("t.(Person)后, t的动态值%v", t)
		fmt.Println()
		fmt.Printf("t.(Person)后, f的动态类型%T", f)
		fmt.Println()
		fmt.Printf("t.(Person)后, f的动态值%v", f)
		fmt.Println()
	}
	fmt.Printf("----------------------------------------------\n")
	fmt.Printf("t.(Tester2)前, m的动态类型%T", m)
	fmt.Println()
	fmt.Printf("t.(Tester2)前, m的动态值%v", m)
	fmt.Println()
	if m, ok2 := t.(Tester2); ok2 {
		fmt.Printf("t.(Tester2)后, t的动态类型%T", t)
		fmt.Println()
		fmt.Printf("t.(Tester2)后, t的动态值%v", t)
		fmt.Println()
		fmt.Printf("t.(Tester2)后, m的动态类型%T", m)
		fmt.Println()
		fmt.Printf("t.(Tester2)后, m的动态值%v", m)
		fmt.Println()
		m.printName()
		fmt.Println()
	}
}
