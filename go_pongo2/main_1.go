package main

import (
	"github.com/flosch/pongo2"
	"net/http"
)

// Pre-compiling the templates at application startup using the
// little Must()-helper function (Must() will panic if FromFile()
// or FromString() will return with an error - that's it).
// It's faster to pre-compile it anywhere at startup and only
// execute the template later.
var tplExample = pongo2.Must(pongo2.FromFile("example.html"))

func examplePage(w http.ResponseWriter, r *http.Request) {
	// Execute the template per HTTP request
	err := tplExample.ExecuteWriter(pongo2.Context{"query": "miya"}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", examplePage)
	http.ListenAndServe(":8080", nil)
}