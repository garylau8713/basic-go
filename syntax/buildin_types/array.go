package main

import "fmt"

func Array() {

	a1 := [3]int{7, 8, 9}
	fmt.Printf("a1: %v, len: %d, cap: %d \n", a1, len(a1), cap(a1))

	a2 := [3]int{7, 8}
	fmt.Printf("a2: %v, len: %d, cap: %d \n", a2, len(a2), cap(a2))

	//Just Init, but not assign value
	var a3 [3]int
	fmt.Printf("a3: %v, len: %d, cap: %d \n", a3, len(a3), cap(a3))

	//a2 := append(a2, 12) Wrong

}
