package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := ":18080"
	err := http.ListenAndServe(port,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
		}),
	)
	if err != nil {
		fmt.Printf("failed to terminate server: %s", err)
		os.Exit(1)
	}
}
