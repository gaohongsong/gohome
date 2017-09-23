package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"Body": "Hello World!"})
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))

	osCurDir, err := os.Getwd()
	fileCurDir , err := filepath.Abs(filepath.Dir(os.Args[0]))
	staticDir := filepath.Join(osCurDir, "src/gohome/go_rest_json/static")

	log.Printf("os: %s, file: %s, static: %s\n", osCurDir, fileCurDir, staticDir)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("."))))
	//http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(staticDir))))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
