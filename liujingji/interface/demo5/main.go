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
	var dog1 *Dog
	dog2 := dog1
	fmt.Printf("nil的值是 %v:\n", nil)
	fmt.Println()
	if dog1 == nil {
		fmt.Printf("dog1的值是 nil%v, dog1的类型是%T\n", dog1, dog1)
	} else {
		fmt.Printf("dog1的值不是 nil%v, dog1的类型是%T\n", dog1, dog1)
	}
	fmt.Println()
	if dog2 == nil {
		fmt.Printf("dog2的值是 nil%v:\n", dog2)
	} else {
		fmt.Printf("dog2的值不是 nil%v:\n", dog2)
	}
	fmt.Println()

	var pet Pet
	if pet == nil {
		fmt.Printf("pet的值是 nil. pet的类型是%T", pet)
	} else {
		fmt.Printf("pet的值不是 nil. pet的类型是%T", pet)
	}
	fmt.Println()
	pet = dog2
	if pet == nil {
		fmt.Printf("pet=dog2后, pet的值是 nil. pet的类型是%T", pet)
	} else {
		fmt.Printf("pet=dog2后, pet的值不是 nil. pet的类型是%T", pet)
	}
	fmt.Println()

}
