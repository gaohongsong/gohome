package main

import "fmt"

func main() {
	var total float64 = 0
	// Array
	arr1 := []float64{
		//98,
		76,
		45,
		32,
		12.34,
	}
	for _, value := range arr1 {
		total += value
	}
	avg1 := total / float64(len(arr1))
	fmt.Printf("Avg(%v) = %f\n", arr1, avg1)

	// Slice
	var sl1 []float64
	//panic: runtime error: index out of range
	//sl1[0] = 1
	sl2 := make([]float64, 5)
	sl3 := arr1[:3]
	//sl1=[], sl2=[0 0 0 0 0], sl3=[76 45 32]
	fmt.Printf("sl1=%v, sl2=%v, sl3=%v\n", sl1, sl2, sl3)

	sl3 = append(sl3, 6, 6, 6)
	//sl1=[], sl2=[0 0 0 0 0], sl3=[76 45 32 6 6 6]
	fmt.Printf("sl1=%v, sl2=%v, sl3=%v\n", sl1, sl2, sl3)

	//sl1=[], sl2=[6 6 6 0 0], sl3=[76 45 32 6 6 6]
	//copy(sl2, sl3[3:])
	//sl1=[], sl2=[76 45 32 6 6], sl3=[76 45 32 6 6 6]
	copy(sl2, sl3)
	fmt.Printf("sl1=%v, sl2=%v, sl3=%v\n", sl1, sl2, sl3)

	//var m1 map[string]int
	//panic: assignment to entry in nil map
	//m1["key"] = 123

	m1 := make(map[string]int)
	m1["key"] = 123

	m2 := map[string]string{
		"a": "ali",
		"t": "tencent",
		"b": "baidu",
	}
	delete(m2, "b")
	value, ok := m2["b"]
	//m1=map[key:123], m2=map[a:ali t:tencent], value=, ok=false
	fmt.Printf("m1=%v, m2=%v, value=%s, ok=%v\n", m1, m2, value, ok)

	//We got b[t] = tencent
	if value, ok := m2["t"]; ok {
		fmt.Println("We got b[t] =", value)
	}else {
		fmt.Println("No such key!")
	}

	// unordered
	for k, v := range map[int]int{
		1: 1,
		2: 4,
		3: 9,
	} {
		fmt.Println(k, v)
	}


}
