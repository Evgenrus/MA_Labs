package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/service1/getString", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Service1")
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}
