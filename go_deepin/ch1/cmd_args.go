package main

import (
	"flag"
	"fmt"
	"math/rand"
)

/*
$ go run cmd_args.go  -v
flag provided but not defined: -v
Usage of cmd_args.exe:
  -max int
        the max value (default 6)
exit status 2

$ go run cmd_args.go  -max 129
77

*/
func main() {
	// define flags
	maxp := flag.Int("max", 6, "the max value")
	enable := flag.Bool("enable", false, "enable it?")

	// parse flag
	flag.Parse()
	fmt.Println(*maxp, *enable)

	// 0~maxp
	fmt.Println(rand.Intn(*maxp))
}
