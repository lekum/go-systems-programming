package main

import (
	"fmt"
	"net/http"
	"os"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func main() {
	PORT := "8081"
	if len(os.Args) == 1 {
		fmt.Println("Using default port number", PORT)
	} else {
		PORT = os.Args[1]
	}

	ADDRESS := ":" + PORT

	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(ADDRESS, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
