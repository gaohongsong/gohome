package main

import (
	"net/http"
	"io"
)

func hello(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set(
		"Content-Type",
		"text/html",
	)

	io.WriteString(
		resp,
		`
		<html><head>title</head><body>Apologize</body></html>
		`,
	)
}

func main() {
	http.HandleFunc("/", hello)
	//http.Handle(
	//	"/static",
	//	http.StripPrefix("/static", http.FileServer(http.Dir("static"))),
	//)
	http.ListenAndServe(":9000", nil)
}
