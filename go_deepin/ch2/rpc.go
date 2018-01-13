package main

import (
	"net"
	"log"
	"fmt"
	"net/rpc"
)

type Server struct {

}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}


func rpcServer() {
	// register
	rpc.Register(new(Server))

	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	for {
		if conn, err := ln.Accept(); err == nil {
			// handle
			go rpc.ServeConn(conn)

		}
	}
}


func rpcClient() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}

	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server.Negate(999) =", result)
	c.Close()
}

/*
Server.Negate(999) = -999
HKJHKHKHJKHJK
You just input:  HKJHKHKHJKHJK
*/
func main() {
	go rpcServer()
	go rpcClient()

	var input string
	fmt.Scanln(&input)
	fmt.Println("You just input: ", input)
}
