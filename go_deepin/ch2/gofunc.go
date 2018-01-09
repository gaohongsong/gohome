package main

import (
	"fmt"
	"reflect"
)

var total float64 = 123

// variadic functions
//	total := sum(arr...)
//	total := sum(1,2,3,4,5)
func sum(args ...float64) float64 {
	total := float64(0)
	for _, arg := range args {
		total += arg
	}
	return total
}

func average(arr []float64) float64 {
	//total := float64(0)
	//for _, v := range arr {
	//	total += v
	//}
	total := sum(arr...)
	//total := sum(10, 10, 10, 10, 10)
	//total in average:  92
	fmt.Println("total in average: ", total)
	return total / float64(len(arr))
}

func max(a, b int) (maxValue int) {
	if a > b {
		maxValue = a
	} else {
		maxValue = b
	}
	// return maxValue
	return
}

// return a function
func makeCounter() func() int {
	cnt := 0
	return func() (ret int) {
		ret = cnt
		cnt ++
		println("ret: ", ret)
		return
	}
}

func clean() {
	fmt.Println("Do some clean job.")
}

//
func one(xPtr *int) {
	*xPtr = 1
}

// recursion
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	// do third
	defer func() {
		fmt.Println("First defer run at last")
	}()

	//do second
	defer clean()
	// recover from panic
	defer func() {
		str := recover()
		fmt.Printf("Recover from panic: %s\n", str)
	}()

	//closure
	increment := makeCounter()

	increment()
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))

	//x = 1, y = 1 new pointer * &
	x := 10
	one(&x)
	y := new(int)
	one(y)
	fmt.Printf("x = %v, y = %v\n", x, *y)

	increment()
	arr := []float64{
		10,
		10,
		10,
		10,
		10,
	}
	var arr1 = [...]float64{
		10,
		10,
		10,
		10,
		10,
	}
	//arr is a []float64(slice), arr1 is a [5]float64(array)
	fmt.Printf("arr is a %v, arr1 is a %v\n", reflect.TypeOf(arr), reflect.TypeOf(arr1))
	//average(arr)= 15.333333333333334
	//total in main:  123
	fmt.Println("average(arr)=", average(arr))
	fmt.Println("average(arr1)=", average(arr1[:]))
	fmt.Println("total in main: ", total)
	fmt.Println(max(100, 200))
	increment()

	func() {
		fmt.Println("Looks like javascript?")
	}()

	//panic: Panic test
	panic("Panic test")

	fmt.Println("Never run to here.")
}
