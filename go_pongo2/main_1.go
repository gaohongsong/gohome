package main

import (
    "fmt"

    "github.com/flosch/pongo2"
)

// Compile the template at application start (for performance reasons)
var tpl = pongo2.Must(pongo2.FromFile("example.html"))

func main() {
    // Execute the template
    rendered_template, err := tpl.Execute(pongo2.Context{
        "name": "miya",
    })
    if err != nil {
        // Handle any execution error (e. g. return a HTTP 500)
        panic(err)
    }
    fmt.Println(rendered_template)
}