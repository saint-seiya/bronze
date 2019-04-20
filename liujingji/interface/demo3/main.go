package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	dog := Dog{"little pig"}

	var pet Pet
	fmt.Printf("动态类型%T, 动态值%v\n", pet, pet)
	pet = &dog
	fmt.Printf("动态类型%T, 动态值%v\n", pet, pet)
	// 赋值会报错
	// var pet1 Pet = dog
	// fmt.Printf("动态类型%T, 动态值%v\n", pet1, pet1)
}
