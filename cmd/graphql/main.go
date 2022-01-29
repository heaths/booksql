package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/graphql", func (w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello, world!")
	})

	log.Fatal(http.ListenAndServe(port(), nil))
}

func port() string {
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		return ":" + val
	}

	return ":8080"
}
