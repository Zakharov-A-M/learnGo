package main

import "fmt"

func main() {
	/*var intPointer *int
	fmt.Printf("%T %#v \n", intPointer, intPointer)

	var a int64 = 7
	fmt.Printf("%T %d \n", a, a)

	var a int64 = 7
	fmt.Printf("%T %#v \n", a, a)

	var pointerA *int64 = &a
	fmt.Printf("%T %#v %#v \n", pointerA, pointerA, *pointerA)

	var newPointer = new(float32)
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
	*newPointer = 3
	fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)*/

	num := 3
	squerePointer(&num)
	fmt.Println(num)

	var wallet1 *int
	fmt.Println(hasWaller(wallet1))

	wallet2 := 0
	fmt.Println(hasWaller(&wallet2))

	wallet3 := 0
	fmt.Println(hasWaller(&wallet3))
}

func squerePointer(num *int) {
	*num *= *num
}

func hasWaller(money *int) bool {
	return money != nil
}
