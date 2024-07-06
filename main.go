package main

//go get -d github  // -d - скачать, но не использовать. Скачивает в $GOPATH/src
//go install github  // -d - скачать, но не использовать. Устанавливает в $GOPATH/pkg и $GOPATH/bin
//go get -u - Обновит пакет
//go build - компилирует пакет и пересобирает пакет
//go mod init github.com/my git/.... - компилирует пакет и пересобирает пакет

import (
	"fmt"
)

var fullName = "Tets Test"

const number = 5
const (
	min = 1
	max = 20
)

// const (
// 	a = iota
// 	max = iota
// )

func main() {
	// функции с заглавной буквы это PUBLIC
	// функции с маленькой буквы это PRIVATE
	//fmt.getField()

	/*ourBoolean := true
	test := "fsdfsdfsdfsdf"
	q := int8(2)

	testNumber := 12345678910
	testNumber2 := 12345678910
	fmt.Println(testNumber + testNumber2)

	var testNumber int64 = 12345678910
	var testNumber2 int = 12345678910
	fmt.Println(testNumber + int64(testNumber2))

	//var name = "test"
	var ddd = "test"
	_ = ddd
	name := "fnsdfsdfksdfsdf"
	//name = "fnsdfsdfksdfsdf";

	fmt.Printf("Type is %T   Value: %v\n", name, name)
	fmt.Printf("Type is %T   Value: %v\n", ourBoolean, ourBoolean)
	fmt.Printf("Type is %T   Value: %v\n", test, test)
	fmt.Printf("Type is %T   Value: %v\n", fullName, fullName)
	fmt.Printf("Type is %T   Value: %v\n", testNumber, testNumber)

	//Greet("Vasya")*/
	//second, first := 5, 2
	/*sum := sum(second, first)
	fmt.Println(sum)*/

	/*sum, multiply := someOperation(second, first)
	fmt.Println(sum, multiply)*/

	/*sum, multiply := nameSomeOperation(second, first)
	fmt.Println(sum, multiply)*/

	/*multiply := func(x, y int) int {
		return x * y
	}
	fmt.Println(multiply(second, first))
	fmt.Println(testSumFunction(second, first, multiply))*/

	/*divider := createDivider(10)
	fmt.Println(divider(2))*/

	/*dollar := 30
	getDollarValue := func() int {
		return dollar
	}

	fmt.Println(getDollarValue())*/

	/*sum := sum(second, first)
	if sum == 10 {
		fmt.Println(sum)
	} else {
		fmt.Println("Not 10")
	}*/

	/*if sum := sum(second, first); sum == 7 {
		fmt.Println(sum)
	} else {
		fmt.Println("Not 10")
	}*/

	/*if sum := sum(second, first); sum == 78 {
		fmt.Println(sum)
	} else {
		fmt.Println("It's 7")
	}*/

	/*sum := sum(second, first)
	if sum >= 7 && sum <= 9 {
		fmt.Println(sum)
	}*/

	/*sum := sum(second, first)
	if sum >= 4 && sum <= 6 {
		fmt.Println(sum)
	} else {
		fmt.Println("It's", sum)
	}*/

	/*sum := sum(second, first)
	if (sum == 4 || sum == 10) || sum == 7 {
		fmt.Println(sum)
	}*/

	/*second++
	first++
	fmt.Println(second, first)*/

	/*second += first
	fmt.Println(second)*/

	/*for i := 0; i < 10; i++ {
		fmt.Println(second, first)
	}*/

	//for i := 0; i < 10; i++ {
	//	fmt.Println(second, first)
	//}

	/*sum := 0
	for sum < 20 {     // FOR AS A WHILE
		fmt.Println(sum)
		sum += 2
	}*/

	/*for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}*/

	/*value := rand.IntN(max-min) + min
	fmt.Println(value)

	switch value {
	case 1:
		fmt.Println("Number ", 1)
	case 2, 3:
		fmt.Println("Number ", 2, "or ", 3)
	case getFour():
		fmt.Println("Number ", 4)
		//fallthrough  // продолжать выполнение след case
	default:
		fmt.Println("Default case is shown")
	}

	switch {
	case value < 3:
		fmt.Println("Number less then ", 3)
	case value >= 3:
		fmt.Println("Number more then 3 ")
	default:
		fmt.Println("Default case is shown")
	}*/

	/*a := [...]int{81, 54, 43, 86}
	for i, v := range a {
	    fmt.Println(i, v)
	}*/
}

func Greet(name string) {
	fmt.Printf("TTTTTTTType is %T   Value: %v\n", number, number)
	fmt.Printf("TTTTTTTType is %T   Value: %v\n", name, name)
}

func sum(a int, b int) int {
	return a + b
}

func someOperation(a int, b int) (int, int) {
	return a + b, a * b
}

func nameSomeOperation(a, b int) (sum int64, multiply int64) {
	sum = int64(a + b)
	multiply = int64(a * b)

	//return sum, multiply
	return
}

func testSumFunction(a, b int, action func(int, int) int) int {
	return action(a, b)
}

func createDivider(value int) func(x int) int {
	return func(x int) int {
		return value / x
	}
}

func getFour() int {
	return 4
}
