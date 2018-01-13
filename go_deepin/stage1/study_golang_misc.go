package main

import (
	"fmt"
	"math"
	"reflect"
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

func getRes(a int) (int, error) {
	i := 0

Here:
	println(i)
	//i ++
	if i++; i < 10 {
		goto Here
	}

	fmt.Printf("final: %d\n", i)

	return 2, nil
}

func main() {
	var a int8
	var b byte = 1
	var c 文本
	var d [10]int
	var e = 123.0
	f := "hello"
	var g float32 = 123.1
	g1 := int(g)

	h := `adfasdfa
	asdfasdfasdfad
	asdfasdfasdfaf`

	//var x, y, z int = 1, 2, 3
	//var x, y, z = 1, 2, 3
	x, _, z := 11, 22, 33

	c = "中文类型名"
	//xxx = 12312

	fmt.Println(a, b, c, d, e, f, x, y, z, xxx, g1, string(65), h)

	fmt.Println(math.MaxInt32)
	fmt.Println("hello world")

	if res, err := getRes(2); err == nil {
		fmt.Printf("res is %v", res)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

	i := 0
	for i < 10 {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Printf("i = %d\n", i)
	}

	for {
		fmt.Printf("i = %d\n", i)
		i++
		if i > 10 {
			fmt.Println("return from while true loop")
			break
		}
	}

	t := [...]int{1, 2, 3, 4, 5, 6}
	t1 := [6]int{1, 2, 3, 4, 5, 6}
	t2 := [3][2]int{{2, 3}, {4, 5}, {5, 6}}
	var t3 [3]int
	t3[0] = 1
	fmt.Println(t, t1, t2, t3)

	for j, k := 0, len(t)-1; j < k; j, k = j+1, k-1 {
		t[j], t[k] = t[k], t[j]
	}

	//len(t1) = 6, cap(t1) = 6len(sl1) = 3, cap(tl1) = 6
	sl1 := t1[0:3]
	//sl1 := t1[:3]
	//sl1 := t1[1:]
	sl2 := sl1[:]
	fmt.Printf("len(t1) = %d, cap(t1) = %d", len(t1), cap(t1))
	fmt.Printf("len(sl1) = %d, cap(tl1) = %d", len(sl1), cap(sl1))
	fmt.Printf("len(sl1) = %d, cap(tl1) = %d\n", len(sl1), cap(sl1))
	//addr(sl1) = 0xc04200a240, addr(sl2) = 0xc04200a240
	fmt.Printf("addr(sl1) = %p, addr(sl2) = %p\n", sl1, sl2)

	//t1 changed too
	sl1[2] = 123
	//panic: runtime error: index out of range
	//sl1[3] = 123
	sl1 = append(sl1, 1, 2, 3)
	fmt.Printf("enough: after append addr(sl1) = %p, addr(sl2) = %p\n", sl1, sl2)
	//sl1: [1 2 123 1 2 3], t1: [1 2 123 1 2 3]
	fmt.Printf("sl1: %v, t1: %v", sl1, t1)

	//sl1 point to another address
	sl1 = append(sl1, 88, 99, 100, 101)
	fmt.Printf("exceed: after append addr(sl1) = %p, addr(sl2) = %p\n", sl1, sl2)
	//sl1: [1 2 123 1 2 3 88 99 100 101], t1: [1 2 123 1 2 3]
	fmt.Printf("sl1: %v, t1: %v", sl1, t1)

	s0 := []int{0, 0, 2}
	println(reflect.TypeOf(s0))
	//sl5 := append([]byte("hello "), "world")

OUTER:
	for i := 0; i < 10; i++ {
		fmt.Printf("loop i: %d", i)
		for j := 0; j < 10; j++ {
			if j > 5 {
				fmt.Printf("break j: %d", j)
				break OUTER
			}
		}
	}
	fmt.Println(t)

	for pos, cc := range "hello" {
		fmt.Printf("%d: %c\n", pos, cc)
	}
	m := map[byte]int{'a': int('a'), 'b': int('b'), 'c': int('c')}
	delete(m, 'c')
	for k, v := range m {
		//for k, v := range map[byte]int{'a': int('a'), 'b': int('b')} {
		fmt.Printf("%c->%d\n", k, v)
	}

	//if-elif-...-else
	var entry int = 2
	//switch entry {
	switch {
	case entry >= 90:
		println("very good")
	case entry > 80 && entry < 90:
		println("good")
	case entry > 60 && entry < 80:
		println("just soso")
	default:
		println("bad")
	}

	//entry = 1
	entry = 2
	//entry = 5
	//entry = 10

	switch entry {
	case 1:
		fmt.Printf("got you entry: %d", entry)
	case 2:
		fmt.Printf("2 - got you entry: %d\n", entry)
		fallthrough
	case 3, 4, 5, 6:
		fmt.Printf("3,4,5,6 - got you entry: %d", entry)
	default:
		fmt.Printf("default - got you entry: %d", entry)
	}

	s1 := make([]int, 10)
	i1 := []int{1,2,3,4,5,6,7,8,9,10,11}
	copy(s1[5:], i1)
	fmt.Printf("i1: %v, s1: %v", i1, s1)

	str1 := "bilibili abasdfas asfasdfs"
	str1slice := []rune(str1)
	copy(str1slice[9:], []rune("asdfasfasdfsd"))
	fmt.Printf("strlince: %v", len(str1slice))
}
