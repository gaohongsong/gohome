package main

import (
    "os"
    "strconv"
    "time"

    "github.com/kataras/iris"
    "runtime"
)

func main() {
    //runtime.GOMAXPROCS(1)
    api := iris.New()

    api.Get("/rest/hello", func(c iris.Context) {
        sleepTime, _ := strconv.Atoi(os.Args[1])
        if sleepTime > 0 {
            time.Sleep(time.Duration(sleepTime) * time.Millisecond)
        }
        c.Text("Hello world")
    })
    api.Run(iris.Addr(":8080"))
}