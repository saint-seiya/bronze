package main

import (
	"fmt"
	_ "imooc/bronze/liujingji/unsafe/demo1"
	"unsafe"
)

func main() {
	var x struct {
		a bool
		b int16
		c []int
	}
	fmt.Printf("Sizeof(x)=%d\n", unsafe.Sizeof(x))
	fmt.Printf("Sizeof(x.a)=%d\n", unsafe.Sizeof(x.a))
	fmt.Printf("Sizeof(x.b)=%d\n", unsafe.Sizeof(x.b))
	fmt.Printf("Sizeof(x.c)=%d\n", unsafe.Sizeof(x.c))
	fmt.Println()
	fmt.Printf("Alignof(x)=%d\n", unsafe.Alignof(x))
	fmt.Printf("Alignof(x.a)=%d\n", unsafe.Alignof(x.a))
	fmt.Printf("Alignof(x.b)=%d\n", unsafe.Alignof(x.b))
	fmt.Printf("Alignof(x.c)=%d\n", unsafe.Alignof(x.c))
	fmt.Println()
	fmt.Printf("Offsetof(x.a)=%d\n", unsafe.Offsetof(x.a))
	fmt.Printf("Offsetof(x.b)=%d\n", unsafe.Offsetof(x.b))
	fmt.Printf("Offsetof(x.c)=%d\n", unsafe.Offsetof(x.c))

	type W struct {
		a byte
		b int32
		c int64
	}

	var w *W
	fmt.Printf("unsafe.Sizeof(*w)=%d\n", unsafe.Sizeof(*w))
	fmt.Printf("unsafe.Sizeof(w.a)=%d\n", unsafe.Sizeof(w.a))
	fmt.Printf("unsafe.Sizeof(w.b)=%d\n", unsafe.Sizeof(w.b))
	fmt.Printf("unsafe.Sizeof(w.c)=%d\n", unsafe.Sizeof(w.c))
	fmt.Println()
	fmt.Printf("Alignof(w.a)=%d\n", unsafe.Alignof(w.a))
	fmt.Printf("Alignof(w.b)=%d\n", unsafe.Alignof(w.b))
	fmt.Printf("Alignof(w.c)=%d\n", unsafe.Alignof(w.c))

	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))

	*pb = 42
	fmt.Println(x.b)

}
