package main

import (
    "github.com/emicklei/go-restful"
    "log"
    "net/http"
	"gohome/go_restful/userservice"
)

func main() {
    restful.Add(userservice.New())
    log.Fatal(http.ListenAndServe(":8888", nil))
}