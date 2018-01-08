package main

import (
	"fmt"
	"math"
)

type (
	byte int8
	rune int32
	文本 string
	Byte int64
)

var (
	x, y, z = 1, 2, 3
)

const xxx = 1231

func main() {
	var a int8
	var b byte = 1
	var c 文本
	var d [10]int
	var e = 123.0
	f := "hello"
	var g float32 = 123.1
	g1 := int(g)

	//var x, y, z int = 1, 2, 3
	//var x, y, z = 1, 2, 3
	x, _, z := 11, 22, 33

	c = "中文类型名"
	//xxx = 12312

	fmt.Println(a, b, c, d, e, f, x, y, z, xxx, g1, string(65))

	fmt.Println(math.MaxInt32)
	fmt.Println("hello world")
}
