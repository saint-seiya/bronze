package main

import (
	"fmt"
)

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string
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
	// 示例1。
	fmt.Println("----------------示例一开始----------------")
	dog := Dog{"little pig"}
	fmt.Printf("dog的name是 %q.\n", dog.Name())
	var pet Pet = dog
	dog.SetName("monster")
	fmt.Printf("dog的name是 %q.\n", dog.Name())
	fmt.Printf("pet的名字是 %q.\n", pet.Name())
	fmt.Println("----------------示例一结束----------------")
	fmt.Println()

	// 示例2。
	fmt.Println("----------------示例二开始----------------")
	dog1 := Dog{"little pig"}
	fmt.Printf("dog1的名字是 %q.\n", dog1.Name())
	dog2 := dog1
	fmt.Printf("dog2的名字是 %q.\n", dog2.Name())
	dog1.name = "monster"
	fmt.Printf("dog1的名字是 %q.\n", dog1.Name())
	fmt.Printf("dog2的名字是 %q.\n", dog2.Name())
	fmt.Println("----------------示例二结束----------------")
	fmt.Println()

	// 示例3。
	fmt.Println("----------------示例三开始----------------")
	dog = Dog{"little pig"}
	fmt.Printf("dog的名字是 %q.\n", dog.Name())
	pet = &dog
	dog.SetName("monster")
	fmt.Printf("dog的名字是 %q.\n", dog.Name())
	fmt.Printf("pet 是 %s, 它的名字是 %q.\n", pet.Category(), pet.Name())
	fmt.Println("----------------示例三结束----------------")
}
