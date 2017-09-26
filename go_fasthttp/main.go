package main

import (
	"log"

	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"encoding/json"
)

func indexHandle(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	b, _ := json.Marshal(map[string]interface{}{
		"name": "miya",
		"age": 1,
	})
	ctx.SetBody(b)
}


func main() {
	router := fasthttprouter.New()
	router.GET("/", indexHandle)
	if err := fasthttp.ListenAndServe(":8080", router.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
