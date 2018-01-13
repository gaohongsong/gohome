package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		du := time.Duration(rand.Intn(250)) * time.Millisecond
		fmt.Println(n, "(", du, ")", ":", i)
		time.Sleep(du)
	}
}

func ping(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
		//invalid operation: <-c (receive from send-only type chan<- string)
		//send-only: send to channel only
		//msg := <- c
		//fmt.Println(msg)
	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	/*
		6 ( 97ms ) : 0
		1 ( 137ms ) : 0
		2 ( 59ms ) : 0
		3 ( 81ms ) : 0
		7 ( 68ms ) : 0
		8 ( 175ms ) : 0
		0 ( 40ms ) : 0
		5 ( 206ms ) : 0
		4 ( 81ms ) : 0
		9 ( 50ms ) : 0
		0 ( 194ms ) : 1
		9 ( 11ms ) : 1
	*/
	for i := 0; i < 10; i++ {
		go f(i)
	}

	// ping >|
	//		  ]---> chan -> printer
	// pong >|
	//strChan := make(chan string)
	// buffered chan
	strChan := make(chan string, 2)
	go ping(strChan)
	go pong(strChan)
	go printer(strChan)

	// select

	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from c1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from c2"
			time.Sleep(time.Second * 3)
		}
	}()

	// "select" between c1 and c2
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("msg1: ", msg1)
			case msg2 := <-c2:
				fmt.Println("msg2: ", msg2)
			case <-time.After(time.Second * 2):
				fmt.Println("timeout")
				//return
			//default:
			//	fmt.Println("nothing ready")
			}
		}
	}()

	// goroutines died quickly if not block here
	var input string
	fmt.Scanln(&input)
}
