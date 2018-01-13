package main

import (
	"net"
	"log"
	"encoding/gob"
	"fmt"
)

func server() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	for {
		if conn, err := ln.Accept(); err == nil {
			go handleConnection(conn)
		}
	}
}

func handleConnection(c net.Conn) {
	var msg string
	if err := gob.NewDecoder(c).Decode(&msg); err != nil {
		log.Fatal(err)
	} else {
		log.Print("Recv: ", msg)
	}
	c.Close()
}

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}

	msg := "hello world"
	fmt.Println("sending: ", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		log.Fatal(err)
	}
	c.Close()
}

/*
sending:  hello world
2018/01/13 10:59:14 Recv: hello world
asdfasdfasdf
You just input:  asdfasdfasdf
*/
func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
	fmt.Println("You just input: ", input)
}
