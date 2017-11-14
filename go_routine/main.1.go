package main

import (
	"time"
	"fmt"
)

func fib(x int) int {
	if x < 2 {
		return x
	}
	return  fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration){
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main()  {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFib(%d) = %d\n", n, fibN)
}
