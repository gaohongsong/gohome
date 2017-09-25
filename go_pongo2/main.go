package main

import (
	"fmt"
	"os"
	"github.com/flosch/pongo2"
)

func init() {
	fmt.Printf("Init in: %s\n", os.Args[0])
}

func main() {
        // Compile the template first (i. e. creating the AST)
    tpl, err := pongo2.FromString("Hello {{ name|capfirst }}!")
    if err != nil {
        panic(err)
    }
    // Now you can render the template with the given
    // pongo2.Context how often you want to.
    out, err := tpl.Execute(pongo2.Context{"name": "florian"})
    if err != nil {
        panic(err)
    }
    fmt.Println(out) // Output: Hello Florian!
}