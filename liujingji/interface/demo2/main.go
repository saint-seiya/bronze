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

type Cat struct {
	name string
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) Category() string {
	return "cat"
}

func main() {
	dog := Dog{"little pig"}
	cat := Cat{"black cat"}
	var pet Pet
	fmt.Printf("动态类型%T, 动态值%v\n", pet, pet)
	pet = &dog
	fmt.Printf("动态类型%T, 动态值%v\n", pet, pet)
	pet = &cat
	fmt.Printf("动态类型%T, 动态值%v\n", pet, pet)
}
